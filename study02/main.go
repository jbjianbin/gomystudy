package main

import (



	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main()  {

	gserver:=g.Server()
	gserver.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writefln("hello wolcome go")
	})
	gserver.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writefln("hello hjb")
	})
	gserver.SetPort(8026)
	gserver.Run()

}

