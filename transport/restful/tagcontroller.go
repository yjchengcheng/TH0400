package restful

import (
	"TH0400/errors"
	"TH0400/logger"
	"TH0400/service"

	"github.com/gin-gonic/gin"
)

func gettags(c *gin.Context) {
	var tagget service.GetTagsCondition

	if err := c.ShouldBind(&tagget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := tagget.GetTags()

	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkAndData(c, res)
	return
}
