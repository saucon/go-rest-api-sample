package gin

import (
	"github.com/Saucon/go-rest-api-sample/internal/customer/controller"
	"github.com/Saucon/go-rest-api-sample/internal/myapp/middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

// setup gin's router
func NewRouter(handler controller.CustomerHandler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"responseCode": "40404", "responseMessage": "Invalid Path"})
	})

	// this will set default timeout to all endpoints
	router.Use(middlewares.TimeoutMiddleware(20*time.Second, func(c *gin.Context) {
		c.Next()
	}))

	api := router.Group("/v1.0/simple-app")
	api.POST("/customer", handler.AddCustomer)

	return router
}
