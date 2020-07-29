package th0400

import (
	"TH0400/transport/restful"

	"github.com/gin-gonic/gin"
)

// RunRestServer create a new gin instance with routers
func RunRestServer() {

	// -----end routers define-------

	// get the server instance
	s := restful.NewRestServer()

	// use recover middleware
	s.Use(gin.Recovery())

	// // use log format
	// s.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf(
	// 		"%s - [%s] \"%s %s %s %d %s \" \" %s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format(time.RFC1123),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.ErrorMessage,
	// 	)
	// }))

	// use default log middleware
	// todo use global logger rather than gin default logger
	s.Use(gin.Logger())

	// run and linsten
	s.Run(":8080")

}

// RunGrpcServer ...
func RunGrpcServer() {
	// todo
}
