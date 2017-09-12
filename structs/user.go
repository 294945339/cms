package structs

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)


type User struct {
	Id           int64       `xorm:"pk autoincr"`
	CreateTime   time.Time   `xorm:"created"`
	CreateUser   int64
	UpdateTime   time.Time	 `xorm:"updated"`
	UpdateUser   int64
	DeletedTime  time.Time   `xorm:"deleted"`
	Status     	 int64       `xorm:"default 1"`
	UserName     string	 	 `xorm:"varchar(25) notnull unique 'user_name'"`
	Password     string
	Salt         string
	Name 		 string
	Email        *string 	 `xorm:"contactEmail"`
	jwt.StandardClaims
}

func (User) TableName() string {
	return "user"
}
