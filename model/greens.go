package model

import (
	_ "github.com/go-sql-driver/mysql"
)

type Greens struct {
	Id int             		`xorm:"not null pk autoincr INT(11)" json:"id"`
	ShopId int             	`json:"shop_id"`
	//BossClassifyId int 		`json:"boss_classify_id"`
	Name string     		`json:"name"`
	Price string     		`json:"price"`
	ImgUrl string 			`json:"img_url"`
	Num int   				`json:"num"`
	Status int				`json:"status"`
	CreatedTime int			`json:"created_time"`
	UpdatedTime int			`json:"updated_time"`
	DeletedTime int			`json:"deleted_time"`
	GreensClassifyId  int		`xorm:"not null INT(11)" json:"greens_classify_id"`
	//GreensClassify     *GreensClassify        `orm:"rel(fk)" json:"greens_classify"`
	//Posts        []*Post   `orm:"reverse(many)" json:"-"`

}





