package repodb

import (
	"github.com/Saucon/go-rest-api-sample/configs/db"
	"github.com/Saucon/go-rest-api-sample/configs/log"
	"github.com/Saucon/go-rest-api-sample/internal/customer/models"
	"github.com/Saucon/go-rest-api-sample/internal/customer/repos"
)

type customerRepoDB struct {
	db  *db.Database
	log *log.LogCustom
}

func NewCustomerRepoDB(dbase *db.Database) repos.ICustomerRepoDb {
	//dbase.DropTable(models.Customer{})
	dbase.AutoMigrate(models.Customer{})

	return customerRepoDB{
		db: dbase,
	}
}

func (c customerRepoDB) AddCustomer(customer models.Customer) error {
	dbGorm := c.db.DB

	trx := dbGorm.Begin()
	err := trx.Debug().Create(&customer).Error
	if err != nil {
		c.log.Error(err, "repo AddCustomer", "", nil, nil, nil, nil, nil)
		trx.Rollback()
		return err
	}
	trx.Commit()

	return err
}

func (c customerRepoDB) GetAllCustomer() ([]models.Customer, error) {
	var customers []models.Customer

	dbGorm := c.db.DB
	trx := dbGorm.Begin()
	err := trx.Find(&customers).Error
	if err != nil {
		c.log.Error(err, "repo GetAllCustomer", "", nil, nil, nil, nil, nil)
		return nil, err
	}

	return customers, nil
}

func (c customerRepoDB) DeleteCustomer() error {
	//TODO implement me
	panic("implement me")
}

func (c customerRepoDB) EditCustomer() models.Customer {
	//TODO implement me
	panic("implement me")
}
