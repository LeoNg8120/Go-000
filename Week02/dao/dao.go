package dao

import (
	"Go-000/Week02/model"
	"Go-000/Week02/sentinel"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"
	"time"
)

func mysqlQuery(uid int) (*model.User,error){
	rand.Seed(time.Now().Unix())
	switch rand.Intn(4) {
		case 1:
			return &model.User{},sql.ErrNoRows
		case 2:
			return &model.User{},sql.ErrConnDone
	}
	return &model.User{UserId: 1,Username: "zhangshan",Age: 17,Sex: "man",Email: "ok@qq.com"},nil
}


func GetUserDetailInfo(uid int)(*model.User,error)  {
	sqllog:=fmt.Sprintf("SELECT * FROM user where uid=%d",uid)
	userinfo,err:=mysqlQuery(uid)
	if errors.Is(err,sql.ErrNoRows) {
		return userinfo,sentinel.DaoErrNoRows
	}
	if err != nil {
		return userinfo,errors.Wrapf(sentinel.NotFound(),"sql: %s error: %v", sqllog, err)
	}
	return userinfo,nil
}