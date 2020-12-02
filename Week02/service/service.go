package service

import (
	"../dao"
	"database/sql"
	"errors"
	"log"
)

func GetDaoUserInfo(userId int64) (info string, err error) {
	info, err = dao.GetUserInfo(userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error(err)
			return info, nil
		}
		return info, err
	}
	return info,nil
}