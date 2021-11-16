package models

type Goods struct {
	Id    int64  `xorm:"autoincr pk" json:"id"`
	Name  string `xorm:"varchar(60)" json:"name"`
	Num   int64  `xorm:"int" json:"num"`
	Image string `xorm:"varchar(255)" json:"image"`
	Url   string `xorm:"varchar(255)" json:"url"`
}
