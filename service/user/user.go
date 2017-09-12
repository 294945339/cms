package user

import (
	. "cms/structs"
	"cms/database/mysql"
	log "github.com/sirupsen/logrus"
)

func GetUserByName(username string) (user User) {
	user, err := mysql.FindUserByName(username)
	if err != nil {
		log.Error(err)
	}
	return
}
