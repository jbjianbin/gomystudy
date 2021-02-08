package main

import (
	"github.com/gogf/gf/frame/g"
	"goland_01/study03/Controller"
	"goland_01/study03/Method"
)




func main()  {

	gserver := g.Server()
	apply :=new(Controller.Apply)
	groupget:=gserver.Group("/api")
	//grouppost:=gserver.Group("post:/api")
	groupget.GET("/loanappy/list", Method.ApplyList) //注册函数
	gserver.BindObject("/get/api/apply",apply)
	gserver.SetPort(8030)
	gserver.Run()


}
