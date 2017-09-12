package structs

type Role struct {
	Id         int64  `xorm:"pk"`
	Rolename   string `xorm:"unique varchar(255)"`
	Desc       string `xorm:"varchar:(255)"`
	Data       string `xorm:"text"`
}