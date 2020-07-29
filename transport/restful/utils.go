package restful

import (
	"TH0400/errors"

	"github.com/gin-gonic/gin"
)

// OkAndData ...
//
// 返回httpcode 200
//
// 返回错误码200
//
// 返回json数据
func OkAndData(c *gin.Context, data interface{}) {

	c.JSON(200, gin.H{
		"code":    errors.OK,
		"message": "OK",
		"data":    data,
	})
	return

}

// ErrAndInfo ...
//
// 返回httpcode 200
//
// 返回错误码
//
// 没有数据
func ErrAndInfo(c *gin.Context, errCode errors.ERRCODE) {

	c.JSON(200, gin.H{
		"code":    errCode,
		"message": errors.ErrMap[errCode],
		"data":    nil,
	})
	return

}

// OkButNoData ...
//
// 返回httpcode 200
//
// 返回错误码 200
//
// 没有数据
func OkButNoData(c *gin.Context) {

	c.JSON(200, gin.H{
		"code":    errors.OK,
		"message": errors.ErrMap[errors.OK],
		"data":    nil,
	})
	return

}
