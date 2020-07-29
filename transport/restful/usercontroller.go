package restful

import (
	"TH0400/errors"
	"TH0400/logger"
	"TH0400/service"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var userRegister service.UserRegister

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

func get(c *gin.Context) {
	var userget service.UserGet

	if err := c.ShouldBind(&userget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	if err := userget.Get(); err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkButNoData(c)
	return
}

func update(c *gin.Context) {
	var userupdate service.UserUpdate

	if err := c.ShouldBind(&userupdate); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	if err := userupdate.Update(); err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkButNoData(c)
	return
}

func delete(c *gin.Context) {
	var userDelete service.UserDelete

	if err := c.ShouldBind(&userDelete); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	if err := userDelete.Delete(); err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkButNoData(c)
	return
}
