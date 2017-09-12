package mysql

import (
	. "cms/structs"
)

func FindUserByName(username string) (user User, err error) {
	_, err = engine.Where("userName = ?", username).Get(&user)
	return
}
