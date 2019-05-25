package controllers

import (
	"order/model"
)

type GreensController struct {
	BaseController
}

//func (this *GreensController) Index() {
//	//店家id
//	shopId,err := this.GetStringChangeInt("shop_id")
//	if err != nil {
//		this.ReturnJson(map[string]string{"message":err.Error()},400)
//	}
//
//	var greens []model.Greens
//	o := orm.NewOrm().QueryTable("greens").Filter("shop_id",shopId)
//
//	o =o.Filter("deleted_time",0).OrderBy("-id")
//	_,err2 :=o.All(&greens)
//	if err2 != nil {
//		this.ReturnJson(map[string]string{"message":"查询错误"},400)
//	}
//	var greensClassifys []model.GreensClassify
//	//var greensClassifysMaps []orm.Params
//	_,err3 := orm.NewOrm().QueryTable("greens_classify").Filter("deleted_time",0).All(&greensClassifys)
//	if err3 != nil {
//		this.ReturnJson(map[string]string{"message":"查询错误"},400)
//	}
//	/*
//		想要修改map中值 需要传地址 才能修改它内部的值
//	*/
//	var newMap = map[int]*model.GreensClassify{}
//
//
//	for _,v := range greensClassifys{
//		newMap[v.Id] = &v
//	}
//	for _,v := range greens {
//		newMap[v.BossClassifyId].Greens = append(newMap[v.BossClassifyId].Greens,v)
//	}
//	var newSlice = make([]interface{},0)
//	for _,v := range newMap{
//		newSlice = append(newSlice,v)
//	}
//	this.ReturnJson(map[string]interface{}{"code":0,"data":newSlice},200)
//}

//func (this *GreensController) Ceshi() {
//	shopId,err := this.GetStringChangeInt("shop_id")
//	if err != nil {
//		this.ReturnJson(map[string]string{"message":err.Error()},400)
//	}
//
//	var greens []model.Greens
//	o := orm.NewOrm().QueryTable("greens").Filter("shop_id",shopId).RelatedSel()
//
//	o =o.Filter("deleted_time",0).OrderBy("-id")
//	_,err2 :=o.All(&greens)
//	if err2 != nil {
//		this.ReturnJson(map[string]string{"message":"查询错误"},400)
//	}
//	var greensClassifys []model.GreensClassify
//	//var greensClassifysMaps []orm.Params
//	_,err3 := orm.NewOrm().QueryTable("greens_classify").RelatedSel().All(&greensClassifys,"classify_name")
//
//	/*****************************************************************************/
//	//test := &model.GreensClassify{Id:1}
//	//o := orm.NewOrm()
//	//err4 := o.Read(test)
//	//o.LoadRelated(test,"Greenss")
//	//
//	//if err4 != nil {
//	//	this.ReturnJson(map[string]string{"message":"查询错误"},400)
//	//}
//	/*****************************************************************************/
//	if err3 != nil {
//		this.ReturnJson(map[string]string{"message":"查询错误"},400)
//	}
//	this.ReturnJson(map[string]interface{}{"code":0,"data":greensClassifys,"greens":greens},200)
//}

func (this *GreensController) Index() {
	var  shopId int
	var err error
	shopId,err = this.GetStringChangeInt("shop_id")
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()+"1"},400)
	}
	var greens []model.GreenClassifyJoin
	err = model.Engine.Join("INNER", "greens_classify", "greens_classify.id = greens.greens_classify_id").
		Where("greens.shop_id = ?",shopId).Where("greens.deleted_time = 0").Find(&greens)
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()+"2"},400)
	}

	greensClassifys := make(map[int64]model.GreensClassify)
	err = model.Engine.Where("deleted_time = 0").Find(&greensClassifys)
	if err != nil {
		this.ReturnJson(map[string]string{"message":err.Error()+"3"},400)
	}

	greensClassifysRes := make(map[int]*model.GreensClassifyRes)
	for _,v := range greens {
		if _, ok := greensClassifysRes[v.GreensClassifyId];ok {
			greensClassifysRes[v.GreensClassifyId].Greens = append(greensClassifysRes[v.GreensClassifyId].Greens,model.Greens{Id:v.Id,ShopId:v.ShopId,Name:v.Name,Price: v.Price,ImgUrl:v.ImgUrl,Num:v.Num,Status:v.Status,CreatedTime:v.CreatedTime,UpdatedTime:v.UpdatedTime,DeletedTime:v.UpdatedTime})
		}else {
			greensClassifysRes[v.GreensClassifyId] = &model.GreensClassifyRes{ Greens: []model.Greens{model.Greens{Id:v.Id,
				ShopId:v.ShopId,Name:v.Name,Price: v.Price,ImgUrl:v.ImgUrl,Num:v.Num,Status:v.Status,CreatedTime:v.CreatedTime,
				UpdatedTime:v.UpdatedTime,DeletedTime:v.UpdatedTime}},Id:v.GreensClassifyId,ClassifyName:v.ClassifyName,ShopId:v.ShopId}
		}
		//greensClassifysRes[v.GreensClassifyId].ClassifyName = v.ClassifyName
	}
	this.ReturnJson(map[string]interface{}{"code":0,"data":greensClassifys,"greens":greens},200)
//}
}
//func (this *GreensController) Ceshi() {}



