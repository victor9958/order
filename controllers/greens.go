package controllers

import (
	"github.com/astaxie/beego/orm"
	"order/model"
)

type GreensController struct {
	BaseController
}

func (this *GreensController) Index() {
	//店家id
	shopId,err := this.GetStringChangeInt("shop_id")
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}

	var greens []model.Greens
	o := orm.NewOrm().QueryTable("greens").Filter("shop_id",shopId)

	o =o.Filter("deleted_time",0).OrderBy("-id")
	_,err2 :=o.All(&greens)
	if err2 != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	var greensClassifys []model.GreensClassify
	//var greensClassifysMaps []orm.Params
	_,err3 := orm.NewOrm().QueryTable("greens_classify").Filter("deleted_time",0).All(&greensClassifys)
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	/*
		想要修改map中值 需要传地址 才能修改它内部的值
	 */
	var newMap = map[int]*model.GreensClassifyRes{}


	for _,v := range greensClassifys{
		newMap[v.Id] = &model.GreensClassifyRes{&v,make([]model.Greens,0)}
	}
	for _,v := range greens {
		newMap[v.BossClassifyId].Greens = append(newMap[v.BossClassifyId].Greens,v)
	}
	var newSlice = make([]interface{},0)
	for _,v := range newMap{
		newSlice = append(newSlice,v)
	}
	this.ReturnJson(map[string]interface{}{"code":0,"data":newSlice},200)
}


