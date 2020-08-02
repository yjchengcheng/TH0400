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
	userregister := UserService{
		UserName: "yjc",
		PassWord: "123456",
	}

	resUserRegister, err := userregister.Login()
	if err != nil {
		t.Error(err.Error())
	}

	t.Log(resUserRegister)
}
