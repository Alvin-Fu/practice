package main

import (
	"github.com/judwhite/go-svc/svc"
		"log"
	"practice/go-daily-lib/freecache/cache"
	"flag"
	beegoConfig "github.com/astaxie/beego/config"
	"os"
	"github.com/mreiferson/go-options"
	"path/filepath"
)

type config map[string]interface{}

type program struct {
	sever *cache.CacheSvr
	quitChan chan struct{}
}
func init(){
	log.SetFlags( log.Lshortfile |log.LstdFlags| log.Llongfile)
}

func main(){
	prg := new(program)
	prg.quitChan = make(chan struct{})
	if err := svc.Run(prg); err != nil {
		log.Fatalf("run err: %v", err)
	}
}

func (p *program) Init(env svc.Environment)error{
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start()error{
	opts := cache.NewOptions()
	flagSet := svrFlagSet(opts)
	flagSet.Parse(os.Args[1:])
	os.Args = os.Args[:1]

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool){
		os.Exit(0)
	}
	var cfg config
	configFile := flagSet.Lookup("config-file").Value.String()
	if configFile != ""{
		pareConfigFile(configFile, (*map[string]interface{})(&cfg))
	}
	options.Resolve(opts, flagSet, cfg)

	cacheSvr := cache.NewCacheSvr(opts)
	cache.Cache = cacheSvr
	cacheSvr.Main()
	p.sever = cacheSvr
	return nil
}

func (p *program) Stop()error{
	return nil
}

func pareConfigFile(configFile string, result *map[string]interface{}){
	configer, err := beegoConfig.NewConfig("ini", configFile)
	if err != nil {
		log.Fatalf("new config err: %v", err)
		return
	}
	cfg := make(map[string]interface{})
	*result = cfg
	tcpAddr := configer.String("tcp.addr")
	if tcpAddr != ""{
		cfg["tcp-addr"] = tcpAddr
	}

	appName := configer.String("AapName")
	if appName != ""{
		cfg["app-name"] = appName
	}

	firstReadTimeout, err := configer.Int64("first.readTimeout")
	if err == nil && firstReadTimeout > 0{
		cfg["firstReadTimeout"] = firstReadTimeout
	}
}

func svrFlagSet(opts *cache.Options)*flag.FlagSet{
	flg := flag.NewFlagSet("cache", flag.ExitOnError)
	flg.Bool("version", false, "cache sever version")
	flg.String("config-file", opts.ConfigFile, "config file")
	flg.String("tcp-addr", opts.TCPAddr, "cache sever tcp addr")
	flg.Int64("first-read-timeout", int64(opts.FirstReadTimeOut), "first read timeout")
	return flg
}