package middlewares

import (
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func timeoutResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{"responseCode": "408", "responseMessage": "Timeout"})
}

func TimeoutMiddleware(time time.Duration, handler func(c *gin.Context)) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time),
		timeout.WithHandler(handler),
		timeout.WithResponse(timeoutResponse),
	)
}
