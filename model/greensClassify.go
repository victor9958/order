package model

type GreensClassify struct {
	Id int             		`json:"id"`
	ShopId int             	`json:"shop_id"`
	ClassifyName string 		`json:"classify_name"`
	CreatedTime int			`json:"created_time"`
	UpdatedTime int			`json:"updated_time"`
	DeletedTime int			`json:"deleted_time"`
	Greens []*Greens	`orm:"reverse(many)"`

}

type GreensClassifyRes struct {
	*GreensClassify
	Greens []Greens
}


