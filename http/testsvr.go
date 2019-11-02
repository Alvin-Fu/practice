package main

import (
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"net"
	"os"
	"path/filepath"
	"practice/http/testsvr"
	"practice/http/util"
	"rpcx/log"
)

type program struct{
	service  *testsvr.HTTPSever
	waitGroup util.WaitGroupWrapper
}

func main(){
	pro := new(program)
	if err := svc.Run(pro); err != nil{
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment)error{
	if env.IsWindowsService(){
		dir := filepath.Dir(os.Args[0])
		return os.Chdir(dir)
	}
	return nil
}

func (p *program) Start()error{
	exitChan := make( <-chan struct{})
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Tcp listen err: %v", err)
		os.Exit(0)
	}
	httpSvr := testsvr.NewHTTPSever(listen)
	p.service = httpSvr
	p.waitGroup.Wrap(func() {
		p.service.Start(exitChan)
	})
	fmt.Println("Service start success!")
	return nil
}


func (p *program) Stop()error{

	return nil
}