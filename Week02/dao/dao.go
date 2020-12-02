package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)


func QueryUser(userId int64)  (name string, err error) {
	return "", sql.ErrNoRows
}

func GetUserInfo(user int64) (userInfo string, err error) {
	info, err := QueryUser(user)
	if err != nil {
		return info, errors.Wrap(err, "getUserInfo err")
	}
	return info, nil
}