package repo

import (
	"TH0400/logger"
	"TH0400/repo/ent"
	"TH0400/repo/ent/user"
	"TH0400/utils"
	"fmt"
)

// UserRepo ...
type UserRepo ent.User

// CreateUser ...
func (u *UserRepo) CreateUser() (uid int, err error) {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	ur, err := edb.User.
		Create().
		SetUserName(u.UserName).
		SetPassword(u.Password).
		SetSchool(u.School).
		Save(ctx)
	if err != nil {
		logger.Errorf("failed create user: %s", err.Error())
		return
	}

	return ur.ID, nil
}

// GetUser ...
func (u *UserRepo) GetUser() (*UserRepo, error) {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if u.ID != 0 {
		ur, err := edb.User.
			Query().
			Where(user.IDEQ(u.ID)).
			Only(ctx)
		return (*UserRepo)(ur), err
	}
	if u.UserName != "" {
		ur, err := edb.User.
			Query().
			Where(user.UserNameEQ(u.UserName)).
			Only(ctx)
		return (*UserRepo)(ur), err
	}

	err := fmt.Errorf("show me ID or UserName")

	return nil, err
}

// UpdateUser ...
func (u *UserRepo) UpdateUser() error {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if u.ID != 0 {
		_, err := edb.User.
			UpdateOneID(u.ID).
			SetUserName(u.UserName).
			SetPassword(u.Password).
			SetSchool(u.School).
			Save(ctx)
		if err != nil {
			logger.Errorf("update user failed: %s", err.Error)
			return err
		}
	}
	err := fmt.Errorf("ID not found")
	return err
}

// DeleteUser ...
func (u *UserRepo) DeleteUser() error {
	ctx, f := utils.GetMTimeoutCtx()
	defer f()

	if u.ID != 0 {
		err := edb.User.
			DeleteOneID(u.ID).
			Exec(ctx)
		if err != nil {
			logger.Errorf("delete user failed: %s", err.Error)
			return err
		}
	}

	err := fmt.Errorf("no User ID detected")
	return err
}
