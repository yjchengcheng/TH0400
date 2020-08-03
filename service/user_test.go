package service

import (
	"testing"
)

func TestRegister(t *testing.T) {

	userregister := UserService{
		UserName: "yjc",
		PassWord: "123456",
		School:   "qinhua",
	}

	resUserRegister, err := userregister.Register()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(resUserRegister)
}

func TestLogin(t *testing.T) {
	userLogin := UserService{
		UserName: "yjc",
		PassWord: "123456",
	}

	resLogin, err := userLogin.Login()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(resLogin)
}

func TestGetUserPublicInfo(t *testing.T) {
	userGetUserPublicInfo := UserService{
		ID: 1,
	}

	resGetUserPublicInfo, err := userGetUserPublicInfo.GetUserPublicInfo()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(resGetUserPublicInfo)
}

func TestGetUserPrivateInfo(t *testing.T) {
	userGetUserPrivateInfo := UserService{
		ID: 1,
	}

	resGetUserPrivateInfo, err := userGetUserPrivateInfo.GetUserPrivateInfo()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(resGetUserPrivateInfo)
}
