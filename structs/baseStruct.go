package structs

import "time"

type BaseStruct struct {
	Id           int64       `xorm:"pk autoincr"`
	CreateTime   time.Time   `xorm:"created"`
	CreateUser   int64
	UpdateTime   time.Time	 `xorm:"updated"`
	UpdateUser   int64
	DeletedTime  time.Time   `xorm:"deleted"`
}