package usecase

import (
	"github.com/Saucon/go-rest-api-sample/configs/log"
	"github.com/Saucon/go-rest-api-sample/internal/customer/models"
	"github.com/Saucon/go-rest-api-sample/internal/customer/repos"
	"github.com/google/uuid"
)

type customerUsecase struct {
	custRepoDB repos.ICustomerRepoDb
	log        *log.LogCustom
}

func (c customerUsecase) AddCustomer(customer models.RequestAddCustomer) error {
	err := c.custRepoDB.AddCustomer(models.Customer{
		Id:      uuid.NewString(),
		Name:    customer.Name,
		Phone:   customer.Phone,
		Address: customer.Address,
	})
	if err != nil {
		return err
	}

	return nil
}

func NewCustomerUsecase(a repos.ICustomerRepoDb, b *log.LogCustom) ICustomer {
	return &customerUsecase{
		custRepoDB: a,
		log:        b,
	}
}
