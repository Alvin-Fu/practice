package main

import (
	"flag"
	"github.com/astaxie/beego"
	"github.com/judwhite/go-svc/svc"
	"practice/web"
	_ "practice/web/routers"
	"rpcx/log"
	"sync"
)

type config map[string]interface{}
type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func main() {
	prg := new(program)
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
	beego.Run()
}

func (p *program) Init(environment svc.Environment) error {
	return nil
}
func (p *program) Start() error {

	return nil
}
func (p *program) Stop() error {
	return nil
}

func parseConfigFile(configFile string, result *map[string]interface{}) {}

func webFlagSet(opt *web.Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("web", flag.ExitOnError)
	flag.Bool("version", false, "")
	flag.String("configFile", opt.ConfigFile, "")
	return flagSet
}
