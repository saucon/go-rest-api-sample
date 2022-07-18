package repos

import "github.com/Saucon/go-rest-api-sample/internal/customer/models"

type ICustomerRepoDb interface {
	AddCustomer(customer models.Customer) error
	GetAllCustomer() ([]models.Customer, error)
	DeleteCustomer() error
	EditCustomer() models.Customer
}
