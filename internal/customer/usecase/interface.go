package usecase

import "github.com/Saucon/go-rest-api-sample/internal/customer/models"

type ICustomer interface {
	AddCustomer(customer models.RequestAddCustomer) error
}
