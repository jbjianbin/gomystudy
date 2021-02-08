package main

import (
	"fmt"

	"github.com/gogf/gf"
	gmd5 "github.com/gogf/gf/crypto/gmd5"



)


type  Dongzuo interface {
	walk() string
	Eate() string
}
type Human struct {
	Age int
	Sex string
}

func (this *Human) walk() string  {

	println("好吃")
	return  "好吃"
}

func (this *Human) Eate() string  {
	println("好走")
	return  "xx"
}

func main()  {
	var man Dongzuo
	man=&Human{Sex: "男",Age: 1}
	man.walk()
	man.Eate()
	fmt.Println("hello")
	fmt.Println(gf.VERSION)
	ss,_ := gmd5.EncryptString("123")
    fmt.Println(ss)
	i:=1

	k:=&i
	sum(k)
	fmt.Println(i)


}

func sum(p *int) {
    *p =3
}