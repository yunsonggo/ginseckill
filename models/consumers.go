package models

type ConsumerModel struct {
	Id int64	`xorm:"autoincr pk" json:"id"`
	NickName string `xorm:"nick_name" json:"nick_name"`
	Name string `xorm:"name" json:"name"`
	Password string `xorm:"password" json:"password"`
}
