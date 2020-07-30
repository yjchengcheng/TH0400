package restful

import (
	"TH0400/errors"
	"TH0400/logger"
	"TH0400/service"

	"github.com/gin-gonic/gin"
)

func gettopics(c *gin.Context) {
	var topicget service.GetTopicsCondition

	if err := c.ShouldBind(&topicget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := topicget.GetTopics()

	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkAndData(c, res)
	return
}

func gettopic(c *gin.Context) {
	var topicget service.GetTopicCondition

	if err := c.ShouldBind(&topicget); err != nil {
		ErrAndInfo(c, errors.BadRequestErr)
		logger.Error(err)
		return
	}

	res, err := topicget.GetTopic()

	if err != nil {
		ErrAndInfo(c, errors.InternalErr)
		logger.Error(err)
	}

	OkAndData(c, res)
	return
}
