package controllers

import (
	"github.com/astaxie/beego"
	"net/http/pprof"
			)

type ProfController struct{
	beego.Controller
}

func (p *ProfController) Get(){
	switch p.Ctx.Input.Param(":app") {
	default:
		pprof.Index(p.Ctx.ResponseWriter, p.Ctx.Request)
	case "":
		pprof.Index(p.Ctx.ResponseWriter, p.Ctx.Request)
	case "cmdline":
		pprof.Cmdline(p.Ctx.ResponseWriter, p.Ctx.Request)
	case "profile":
		pprof.Profile(p.Ctx.ResponseWriter, p.Ctx.Request)
	case "trace":
		pprof.Trace(p.Ctx.ResponseWriter, p.Ctx.Request)
	case "symbol":
		pprof.Symbol(p.Ctx.ResponseWriter, p.Ctx.Request)
	}
}

