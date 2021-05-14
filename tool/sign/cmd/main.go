package main

// NOTE: change -lang option to add more languages support.
//go:generate gotext -srclang=en update -out=catalog.go -lang=en,zh-hans,zh-hant

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"practice/tool/sign/conf"
	"practice/tool/sign/server"
	"syscall"
	"time"

	"github.com/judwhite/go-svc/svc"
	"github.com/mreiferson/go-options"
)

type program struct {
	svr *server.StoreSvr
}

type config map[string]interface{}

// Validate settings in the config file, and fatal on errors
func (cfg config) Validate() {
	// special validation/translation
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
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
	opts := conf.NewOptions()
	flagSet := accdFlagSet(opts)
	flagSet.Parse(os.Args[1:])
	// NOTE: clear all args after we have parsed, avoid RPCServer parse these args again
	os.Args = os.Args[:1]

	rand.Seed(time.Now().UTC().UnixNano())

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		os.Exit(0)
	}
	var cfg config
	//configFile := flagSet.Lookup("config").Value.String()
	//if configFile != "" {
	//	err := parseConfigFile(configFile, (*map[string]interface{})(&cfg))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//cfg.Validate()

	options.Resolve(opts, flagSet, cfg)
	svr := server.NewStoreSvr(opts)
	err := svr.Main()
	if err != nil {
		log.Fatalf("svr start err: %v", err)
	}
	p.svr = svr
	return nil
}

func (p *program) Stop() error {
	if p.svr != nil {
		p.svr.Exit()
	}
	return nil
}

func accdFlagSet(opts *conf.Options) *flag.FlagSet {
	flagSet := flag.NewFlagSet("access", flag.ExitOnError)
	flagSet.Bool("version", false, "print version string")
	flagSet.Int("node-id", int(opts.ID), "match server node id([0-127])")
	flagSet.String("pid-file", opts.PIDFile, "path of pid file, default is ./matchsvr_<ID>.pid")

	flagSet.String("http-addr", opts.HTTPAddr, "<addr>:<port> to listen on for admin http server, no http server start if empty")
	flagSet.Bool("verbose", opts.Verbose, "")

	return flagSet
}

func parseConfigFile(configFile string, result *map[string]interface{}) error {
	//config, err := beegoConfig.NewConfig("ini", configFile)
	//if err != nil {
	//	return err
	//}
	//
	//cfg := make(map[string]interface{})
	//
	//
	//*result = cfg
	return nil
}
