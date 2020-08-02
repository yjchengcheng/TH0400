package restful

import (
	"TH0400/errors"
	"TH0400/logger"
	"TH0400/service"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var userRegister service.UserService

	if err := c.ShouldBindJSON(&userRegister); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := userRegister.Register()
	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
		return
	}

	OkAndData(c, res)
	return
}

func getuserpublicinfo(c *gin.Context) {
	var userget service.UserService

	if err := c.ShouldBind(&userget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := userget.GetUserPublicInfo()

	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkAndData(c, res)
	return
}

func getuserprivateinfo(c *gin.Context) {
	var userget service.UserService

	if err := c.ShouldBind(&userget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := userget.GetUserPrivateInfo()

	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkAndData(c, res)
	return
}
