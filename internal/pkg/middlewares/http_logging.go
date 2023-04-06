package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HttpLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		request := c.Request
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientIp := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		log.Infof("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			clientIp,
			timeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			method,
			path,
			request.Proto,
			statusCode,
			latency,
			request.UserAgent(),
			errorMessage,
		)
	}
}
