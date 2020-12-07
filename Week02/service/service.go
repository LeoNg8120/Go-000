package service

import (
	"Go-000/Week02/dao"
	"Go-000/Week02/model"
)

type Service interface {
	GetUserInfo()(user *model.User,err error)
}

type service struct {
	uid int
}

func NewServcie(uid int)(Service,error)  {
	return &service{uid: uid},nil
}

func (s *service)GetUserInfo()(user *model.User,err error) {
	return dao.GetUserDetailInfo(s.uid)
}
