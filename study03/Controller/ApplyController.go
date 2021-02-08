package Controller

import "github.com/gogf/gf/net/ghttp"

type Apply struct {}

func (this *Apply) Show(r *ghttp.Request)  {

	r.Response.Writefln("我是对象注册")

}

func (this *Apply) Name(r *ghttp.Request)  {

	r.Response.Writefln("我的名字是黄建斌")

}