package service

import (
	"TH0400/logger"
	"TH0400/repo"
	"TH0400/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// UserService 。。。
type UserService struct {
	ID        int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	PassWord  string    `json:"pass_word"`
	School    string    `json:"school"`
	Level     int       `json:"level"`
	Likes     int       `json:"likes"`
	Views     int       `json:"views"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}

func (u *UserService) nullID() (yes bool) {
	if u.ID == 0 {
		yes = true
	}
	return
}

func (u *UserService) nullUserName() (yes bool) {
	if u.UserName == "" {
		yes = true
	}
	return
}

func (u *UserService) nullPassword() (yes bool) {
	if u.PassWord == "" {
		yes = true
	}
	return
}

func (u *UserService) nullSchool() (yes bool) {
	if u.School == "" {
		yes = true
	}
	return
}

func (u *UserService) isDeleted() (yes bool) {
	return u.IsDeleted
}

// Register do register
func (u *UserService) Register() (regRes gin.H, err error) {

	if u.nullUserName() || u.nullPassword() {
		err = fmt.Errorf("empty userName: [%s] or passWord: [%s]", u.UserName, u.PassWord)
		logger.Errorf("userservice; %s", err.Error())
		return
	}

	u.PassWord = utils.EncodePassword(u.PassWord)

	euser := repo.UserRepo{
		UserName: u.UserName,
		Password: u.PassWord,
		School:   u.School,
	}

	uid, err := euser.CreateUser()
	if err != nil {
		logger.Errorf("failed create token: %s", err.Error())
		return
	}

	token, err := utils.GenerateToken(uid, u.UserName, u.PassWord)
	if err != nil {
		logger.Errorf("failed create token: %s", err.Error())
		return
	}

	regRes = gin.H{
		"user_id": uid,
		"token":   token,
	}

	return

}

// Login ..
func (u *UserService) Login() (resLogin gin.H, err error) {
	u.PassWord = utils.EncodePassword(u.PassWord)

	if u.nullUserName() || u.nullPassword() {
		err = fmt.Errorf("empty userName: [%s] or passWord: [%s]", u.UserName, u.PassWord)
		logger.Errorf("userservice; %s", err.Error())
		return
	}

	enuser := repo.UserRepo{
		UserName: u.UserName,
	}

	ur, err := enuser.GetUser()
	if err != nil {
		logger.Error(err)
		return
	}
	if ur.Password != u.PassWord {
		err = fmt.Errorf("wrong password")
		return
	}

	token, err := utils.GenerateToken(ur.ID, u.UserName, u.PassWord)
	if err != nil {
		logger.Errorf("failed create token: %s", err.Error())
		return
	}

	resLogin = gin.H{
		"user_id": ur.ID,
		"token":   token,
	}

	return

}

// GetUserPublicInfo ...
func (u *UserService) GetUserPublicInfo() (pubinfores map[string]interface{}, err error) {
	pubinfores = make(map[string]interface{})

	if u.nullID() {
		err = fmt.Errorf("empty userID: [%d]", u.ID)
		logger.Errorf("userservice; %s", err.Error())
		return
	}

	euser := repo.UserRepo{
		ID: u.ID,
	}

	ur, err := euser.GetUser()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	pubinfores = gin.H{
		"user_name":   ur.UserName,
		"level":       ur.Level,
		"school":      ur.School,
		"likes_count": ur.Likes,
	}

	return
}

// GetUserPrivateInfo ...
func (u *UserService) GetUserPrivateInfo() (priinfores gin.H, err error) {

	euser := repo.UserRepo{
		ID: u.ID,
	}

	ur, err := euser.GetUser()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	priinfores = gin.H{
		"user_name":   ur.UserName,
		"level":       ur.Level,
		"school":      ur.School,
		"likes_count": ur.Likes,
		"view_count":  ur.Views,
	}

	return
}
