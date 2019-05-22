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
	//shopId,err := this.GetStringChangeInt("shop_id")
	//if err != nil {
	//	this.ReturnJson(map[string]string{"message":err.Error()},400)
	//}
	//
	//var greens []model.Greens
	//o := orm.NewOrm().QueryTable("greens").Filter("shop_id",shopId)
	//
	//o =o.Filter("deleted_time",0).OrderBy("-id")
	//_,err2 :=o.All(&greens)
	//if err2 != nil {
	//	this.ReturnJson(map[string]string{"message":"查询错误"},400)
	//}
	//var greensClassifys []model.GreensClassify
	////var greensClassifysMaps []orm.Params
	//_,err3 := orm.NewOrm().QueryTable("greens_classify").Filter("deleted_time",0).All(&greensClassifys)
	//if err3 != nil {
	//	this.ReturnJson(map[string]string{"message":"查询错误"},400)
	//}
	///*
	//	想要修改map中值 需要传地址 才能修改它内部的值
	// */
	//var newMap = map[int]*model.GreensClassifyRes{}
	//
	//
	//for _,v := range greensClassifys{
	//	newMap[v.Id] = &model.GreensClassifyRes{&v,make([]model.Greens,0)}
	//}
	//for _,v := range greens {
	//	newMap[v.BossClassifyId].Greens = append(newMap[v.BossClassifyId].Greens,v)
	//}
	//var newSlice = make([]interface{},0)
	//for _,v := range newMap{
	//	newSlice = append(newSlice,v)
	//}
	//this.ReturnJson(map[string]interface{}{"code":0,"data":newSlice},200)
}

func (this *GreensController) Ceshi() {
	shopId,err := this.GetStringChangeInt("shop_id")
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()},400)
	}

	var greens []model.Greens
	o := orm.NewOrm().QueryTable("greens").Filter("shop_id",shopId).RelatedSel()

	o =o.Filter("deleted_time",0).OrderBy("-id")
	_,err2 :=o.All(&greens)
	if err2 != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	var greensClassifys []model.GreensClassify
	//var greensClassifysMaps []orm.Params
	_,err3 := orm.NewOrm().QueryTable("greens_classify").Filter("deleted_time",0).RelatedSel().All(&greensClassifys)

	/*****************************************************************************/
	//test := &model.GreensClassify{Id:1}
	//o := orm.NewOrm()
	//err4 := o.Read(test)
	//o.LoadRelated(test,"Greenss")
	//
	//if err4 != nil {
	//	this.ReturnJson(map[string]string{"message":"查询错误"},400)
	//}
	/*****************************************************************************/
	if err3 != nil {
		this.ReturnJson(map[string]string{"message":"查询错误"},400)
	}
	this.ReturnJson(map[string]interface{}{"code":0,"data":greensClassifys,"greens":greens},200)
}

//--user=www --group=www --prefix=/www/server/nginx --add-module=/home/ann/nginx/src/ngx_devel_kit --add-module=/home/ann/nginx/src/lua_nginx_module --add-module=/home/ann/nginx/src/ngx_cache_purge --add-module=/home/ann/nginx/src/nginx-sticky-module --add-module=/home/ann/nginx/src/nginx-http-concat --with-http_stub_status_module --with-http_ssl_module --with-http_image_filter_module --with-http_spdy_module --with-http_gzip_static_module --with-http_gunzip_module --with-ipv6 --with-http_sub_module --with-http_flv_module --with-http_addition_module --with-http_realip_module --with-http_mp4_module --with-ld-opt=-Wl,-E --with-pcre=/home/ann/nginx/src/pcre-8.42 --add-module=/www/server/ngx_brotli


