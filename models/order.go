package models

type Order struct {
	Id int64 `xorm:"autoincr pk" json:"id"`
	UserId int64 `xorm:"user_id" json:"user_id"`
	GoodsId int64 `xorm:"goods_id" json:"goods_id"`
	Status int64 `xorm:"status" json:"status"`
}
