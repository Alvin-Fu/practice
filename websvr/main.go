package main

import (
	"flag"
	"github.com/astaxie/beego"
	beegoConfig "github.com/astaxie/beego/config"
	"github.com/judwhite/go-svc/svc"
	"github.com/mreiferson/go-options"
	"os"
	"path/filepath"
	_ "practice/websvr/routers"
	"practice/websvr/uitl"
	"rpcx/log"
	"practice/websvr/option"
	"practice/websvr/routers"
)

// 定义一个保存配置的类型
type config map[string]interface{}
// 服务启动结构体
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
// 使用svc的方式启动服务
func (p *program) Init(env svc.Environment) error {
	if  env.IsWindowsService() {
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}
func (p *program) Start() error {
	opts := option.NewOption()
	flagSet := webFlagSet(opts)
	flagSet.Parse(os.Args[1:])
	os.Args = os.Args[:1]
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		os.Exit(0)
	}
	var cfg config
	// 获取命令行参数中的配置文件
	configFile := flagSet.Lookup("configFile").Value.String()
	if configFile != ""{
		parseConfigFile(configFile, (*map[string]interface{})(&cfg))
	}
	// 在使用resolve的时候需要在opts中写标签，可以查看resolve的函数说明
	options.Resolve(opts, flagSet, cfg)
	routers.RegRouter()
	// 启动框架的http服务
	p.wg.Wrap(func() {
		beego.Run(opts.HTTPHost + ":" +opts.HTTPPort)
	})
	return nil
}
func (p *program) Stop() error {
	return nil
}

/* parseConfigFile  获取配置中的参数
 *
 */
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

/* webFlagSet  设置命令行参数
 *
 */
func webFlagSet(opt *option.Option) *flag.FlagSet {
	flagSet := flag.NewFlagSet("web", flag.ExitOnError)
	flagSet.Bool("version", false, "http sever version!")
	flagSet.String("configFile", opt.ConfigFile, "http sever config file!")
	flagSet.String("HTTPHost", opt.HTTPHost, "http sever host!")
	flagSet.String("HTTPPort", opt.HTTPPort, "http sever port!")
	return flagSet
}
