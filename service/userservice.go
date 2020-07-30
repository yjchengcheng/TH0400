package service

import (
	"TH0400/logger"
	"TH0400/repo"
	"TH0400/utils"
)

// UserRegister 。。。
type UserRegister struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	School   string `json:"school"`
}

//UserResponseRegister ...UserResponseRegister
type UserResponseRegister struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

// UserGet ...
type UserGet struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

// UserUpdate ...
type UserUpdate struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	School   string `json:"school"`
}

// UserDelete ...
type UserDelete struct {
	ID int `json:"id"`
}

// UserGetByID ...
type UserGetByID struct {
	ID int `json:"id"`
}

// UserGetPublicInfoRes ...
type UserGetPublicInfoRes struct {
	UserName   string `json:"user_name"`
	Level      int    `json:"level"`
	School     string `json:"school"`
	LikesCount int    `json:"likes_count"`
	// QuestionsCount int    `json:"questions_count"`
	// AnswersCount   int    `json:"answers_count"`
}

// UserGetPrivateInfoRes ...
type UserGetPrivateInfoRes struct {
	UserName   string `json:"user_name"`
	Level      int    `json:"level"`
	School     string `json:"school"`
	LikesCount int    `json:"likes_count"`
	ViewsCount int    `json:"views_count"`
	// QuestionsCount int    `json:"questions_count"`
	// AnswersCount   int    `json:"answers_count"`
}

// Register do register
func (u *UserRegister) Register() (regRes UserResponseRegister, err error) {

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

	regRes.UserID = uid
	regRes.Token = token

	return

}

// Get ...
func (u *UserGet) Get() error {

	euser := repo.UserRepo{
		ID:       u.ID,
		UserName: u.UserName,
	}

	_, err := euser.GetUser()

	return err
}

// Update ...
func (u *UserUpdate) Update() error {

	u.PassWord = utils.EncodePassword(u.PassWord)

	euser := repo.UserRepo{
		ID:       u.ID,
		UserName: u.UserName,
		Password: u.PassWord,
		School:   u.School,
	}

	return euser.UpdateUser()
}

// Delete ...
func (u *UserDelete) Delete() error {

	euser := repo.UserRepo{
		ID: u.ID,
	}

	return euser.DeleteUser()
}

// GetUserPublicInfo ...
func (u *UserGetByID) GetUserPublicInfo() (pubinfores *UserGetPublicInfoRes, err error) {

	euser := repo.UserRepo{
		ID: u.ID,
	}

	ur, err := euser.GetUserByID()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	pubinfores.UserName = ur.UserName
	pubinfores.Level = ur.Level
	pubinfores.School = ur.School
	pubinfores.LikesCount = ur.Likes

	return
}

// GetUserPrivateInfo ...
func (u *UserGetByID) GetUserPrivateInfo() (priinfores *UserGetPrivateInfoRes, err error) {

	euser := repo.UserRepo{
		ID: u.ID,
	}

	ur, err := euser.GetUserByID()
	if err != nil {
		logger.Errorf("not found user: %s", err.Error())
		return
	}

	priinfores.UserName = ur.UserName
	priinfores.Level = ur.Level
	priinfores.School = ur.School
	priinfores.LikesCount = ur.Likes
	priinfores.ViewsCount = ur.Views

	return
}
