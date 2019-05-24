package model

type Greens struct {
	Id int             		`json:"id"`
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
	GreensClassify     *GreensClassify        `orm:"rel(fk)" json:"greens_classify"`
	//Posts        []*Post   `orm:"reverse(many)" json:"-"`

}


