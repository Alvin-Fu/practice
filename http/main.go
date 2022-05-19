package main

import (
	"flag"
	"log"
	"net"
	"os"
	"path/filepath"
	options2 "practice/http/options"
	"practice/http/svrmonitoring"
	"practice/http/util"

	beegoConfig "github.com/astaxie/beego/config"
	"github.com/judwhite/go-svc"
	"github.com/mreiferson/go-options"
)

type config map[string]interface{}
type program struct {
	wg   util.WaitGroupWrapper
	quit chan struct{}
}

func main() {
	prg := new(program)
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	if env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}
func (p *program) Start() error {
	opts := options2.NewOption()
	flagSet := webFlagSet(opts)
	flagSet.Parse(os.Args[1:])
	os.Args = os.Args[:1]
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		os.Exit(0)
	}
	ExitChan := make(<-chan struct{})
	var cfg config
	//configFile := flagSet.Lookup("configFile").Value.String()
	//if configFile != "" {
	//	parseConfigFile(configFile, (*map[string]interface{})(&cfg))
	//}
	options.Resolve(opts, flagSet, cfg)
	tcpListener, err := net.Listen("tcp", opts.HTTPHost+":"+opts.HTTPPort)
	if err != nil {
		log.Fatalf("Tcp connect fail err: %v", err)
		return err
	}
	HTTPSever := svrmonitoring.NewHTTPSever(tcpListener)
	p.wg.Wrap(func() {
		HTTPSever.Start(ExitChan)
	})
	return nil
}
func (p *program) Stop() error {
	return nil
}

func parseConfigFile(configFile string, result *map[string]interface{}) {
	config, err := beegoConfig.NewConfig("ini", configFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	cfg := make(map[string]interface{})
	*result = cfg
	host := config.String("HTTPHost")
	if host != "" {
		cfg["HTTPHost"] = host
	}
	port := config.String("HTTPPort")
	if port != "" {
		cfg["HTTPPort"] = port
	}
	name := config.String("APPName")
	if name != "" {
		cfg["APPName"] = name
	}

}

func webFlagSet(opt *options2.Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("web", flag.ExitOnError)
	flagSet.Bool("version", false, "http sever version!")
	flagSet.String("configFile", opt.ConfigFile, "http sever config file!")
	flagSet.String("HTTPHost", opt.HTTPHost, "http sever host!")
	flagSet.String("HTTPPort", opt.HTTPPort, "http sever port!")
	return flagSet
}
