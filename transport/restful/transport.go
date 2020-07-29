package restful

import (
	"TH0400/errors"
	"TH0400/logger"
	"TH0400/repo/rcache"

	"github.com/gin-gonic/gin"
)

// NewRestServer ...
func NewRestServer() *gin.Engine {

	e := gin.New()

	userGroup := e.Group("/user")
	{
		userGroup.POST("/register", register)
	}

	// -----define routres here------
	v1 := e.Group("/v1")
	v1.Use(holdOn)
	{

		// topicGroup := v1.Group("/topic")
		{
			// topicGroup.POST("/release", release)
		}

	}

	return e
}

func holdOn(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.Abort()
		ErrAndInfo(c, errors.AuthNeedErr)
		logger.Info("打日志")
		return
	}

	userID, ok := rcache.GetToken(token)
	if !ok {
		c.Abort()
		ErrAndInfo(c, errors.AuthNeedErr)
		logger.Info("打日志")
		return
	}

	// fixme example
	logger.Info(userID)

}
