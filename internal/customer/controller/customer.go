package controller

import (
	"github.com/Saucon/go-rest-api-sample/configs/log"
	"github.com/Saucon/go-rest-api-sample/internal/customer/models"
	"github.com/Saucon/go-rest-api-sample/internal/customer/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	custUsecase usecase.ICustomer
	log         *log.LogCustom
}

func NewCustomerHandler(a usecase.ICustomer, b *log.LogCustom) CustomerHandler {
	return CustomerHandler{
		custUsecase: a,
		log:         b,
	}
}

func (ch *CustomerHandler) AddCustomer(c *gin.Context) {
	var request models.RequestAddCustomer
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = ch.custUsecase.AddCustomer(request)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, "added")
}
