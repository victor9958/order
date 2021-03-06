package model

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
type MyDb struct {
	*xorm.Engine
}
var Engine *xorm.Engine


//func init(){
//	user := beego.AppConfig.String("mysqluser")
//	passwd := beego.AppConfig.String("mysqlpass")
//	host := beego.AppConfig.String("mysqlurls")
//	port,err := beego.AppConfig.Int("mysqlport")
//	dbname := beego.AppConfig.String("mysqldb")
//
//	if err!=nil {
//		port =3306
//	}
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//// set default database
	//orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	user, passwd, host, port, dbname))
	//
	//orm.RegisterModel(new(Greens),new(GreensClassify))
	//
	//if beego.AppConfig.String("runmode") == "dev" {
	//	orm.Debug = true
	//}
	//
	//orm.RunSyncdb("default", false, false)


	//orm.RunSyncdb("default", false, false)
//}


func init(){

	user := beego.AppConfig.String("mysqluser")
	passwd := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlurls")
	port,err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")

	if err!=nil {
		port =3306
	}
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//// set default database
	//orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
	//	user, passwd, host, port, dbname))
	//
	//orm.RegisterModel(new(Order))
	//
	//if beego.AppConfig.String("runmode") == "dev" {
	//	orm.Debug = true
	//}
	//
	//orm.RunSyncdb("default", false, false)
	Engine,_ = xorm.NewEngine("mysql",fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbname))
	//if err != nil {
	//	beego.Error(err)
	//	return
	//}
	//连接测试
	if err := Engine.Ping();err != nil{
		return
	}

	//日志打印sql
	Engine.ShowSQL(true)

	//设置连接池的大小
	Engine.SetMaxIdleConns(5)

	//设置最大打开的连接数
	Engine.SetMaxOpenConns(5)


	//名称映射规则主要负责结构名称到表名和结构体field到表字段的名称
	Engine.SetTableMapper(core.SnakeMapper{})

}


func Pluck(db *xorm.Session,columnName string)(map[int]interface{}, error){
	//两个数组
	res := make(map[int]interface{})
	var ids []string
	var columns []string
	db2 := *db
	err := db.Cols("id").Find(&ids)
	if err != nil {
		return res,err
	}
	err2 := db2.Cols(columnName).Find(&columns)
	if err != nil {
		return res,err2
	}
	//循环对应产生一个map
	for k,v := range ids{
		id,_:= strconv.Atoi(v)
		res[id] = columns[k]
		//res = append(res,)
	}
	return res,nil
}
