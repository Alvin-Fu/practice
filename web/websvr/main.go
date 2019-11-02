package main

import (
	"flag"
	"github.com/astaxie/beego"
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/judwhite/go-svc/svc"
	"github.com/mreiferson/go-options"
	"os"
	"path/filepath"
	"practice/web"
	_ "practice/web/routers"
	"practice/web/uitl"
	"rpcx/log"
)

type config map[string]interface{}
type program struct {
	wg   uitl.WaitGroupWrapper
	quit chan struct{}
}

func main() {
	prg := new(program)
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if  env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}
func (p *program) Start() error {
	opts := web.NewOption()
	flagSet := webFlagSet(opts)
	flagSet.Parse(os.Args[1:])
	os.Args = os.Args[:1]
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		os.Exit(0)
	}
	var cfg config
	configFile := flagSet.Lookup("configFile").Value.String()
	if configFile != ""{
		parseConfigFile(configFile, (*map[string]interface{})(&cfg))
	}
	options.Resolve(opts, flagSet, cfg)
	p.wg.Wrap(func() {
		beego.Run(opts.HTTPHost + ":" +opts.HTTPPort)
	})
	return nil
}
func (p *program) Stop() error {
	return nil
}

func parseConfigFile(configFile string, result *map[string]interface{}) {
	config, err := beegoConfig.NewConfig("ini", configFile)
	if err != nil{
		log.Fatal(err)
		return
	}
	cfg := make(map[string]interface{})
	*result = cfg
	host := config.String("HTTPHost")
	if host != ""{
		cfg["HTTPHost"] = host
	}
	port := config.String("HTTPPort")
	if port != ""{
		cfg["HTTPPort"] = port
	}
	name := config.String("APPName")
	if name != ""{
		cfg["APPName"] = name
	}

}

func webFlagSet(opt *web.Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("web", flag.ExitOnError)
	flag.Bool("version", false, "http sever version!")
	flag.String("configFile", opt.ConfigFile, "http sever config file!")
	flag.String("HTTPHost", opt.HTTPHost, "http sever host!")
	flag.String("HTTPPort", opt.HTTPPort, "http sever port!")
	return flagSet
}
