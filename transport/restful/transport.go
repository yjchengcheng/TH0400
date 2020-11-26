package restful

import (
	"github.com/gin-gonic/gin"
)

// NewRestServer ...
func NewRestServer() *gin.Engine {

	e := gin.New()

	// -----define routres here------
	v1 := e.Group("/v1")
	v1.Use(holdOn)
	{

	}

	return e
}

func holdOn(c *gin.Context) {

}
