package structs

type User struct {
	BaseStruct
	UserName     string	 `xorm:"varchar(25) notnull unique 'user_name'"`
	Password     string
	Description  *string
	Salt         string
	HashPass     string  `xorm:"hashpass"`
}

func (User) TableName() string {
	return "user"
}
