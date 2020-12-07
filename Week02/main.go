package main

import (
	"Go-000/Week02/sentinel"
	"Go-000/Week02/service"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	uid:=2
	svr, err := service.NewServcie(uid)
	if err != nil {
		panic("NewService err:"+err.Error())
	}
	userInfo,err :=svr.GetUserInfo()
	if errors.Is(err,sentinel.DaoErrNoRows) {
		fmt.Println("get user info is null,next job...")
		return
	}
	if err!=nil {
		fmt.Printf("original error: %T %v\n",errors.Cause(err),errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n",err)
		return
	}
	fmt.Println("use userInfo do other thing:",userInfo)
}
