package controllers

import "order/funcs"

type CeshiController struct {
	BaseController
}

func(this *CeshiController)Ceshi(){
	pwd := funcs.MakeMd5("123456")
	this.ReturnJson(pwd,200)
}
