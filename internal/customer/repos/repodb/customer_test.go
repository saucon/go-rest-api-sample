package repodb

import (
	"fmt"
	"github.com/Saucon/go-rest-api-sample/configs/db"
	"github.com/Saucon/go-rest-api-sample/configs/env"
	"github.com/Saucon/go-rest-api-sample/internal/customer/models"
	"github.com/google/uuid"
	"os"
	"strconv"
	"testing"
)

func TestAddCustomerDB(t *testing.T) {

	env.NewEnv(os.ExpandEnv("$GOPATH/src/github.com/Saucon/go-rest-api-sample/.env"))
	cfg := env.Config

	dBase := db.NewDB(cfg, false)

	custRepoDB := NewCustomerRepoDB(dBase)

	err := custRepoDB.AddCustomer(models.Customer{
		Id:      uuid.NewString(),
		Name:    "PT Wow",
		Phone:   "080988888",
		Address: "Jl Buntu Daamai No.77",
	})
	if err != nil {
		t.Fatal(err)
	}
	dBase.DB.Exec("DELETE FROM customers")
}

func TestGetAllCustomerDB(t *testing.T) {
	var customers []models.Customer
	env.NewEnv(os.ExpandEnv("$GOPATH/src/github.com/Saucon/go-rest-api-sample/local/.env"))
	cfg := env.Config

	dBase := db.NewDB(cfg, false)

	custRepoDB := NewCustomerRepoDB(dBase)

	for i := 1; i <= 1000; i++ {
		err := custRepoDB.AddCustomer(models.Customer{
			Id:      uuid.NewString(),
			Name:    "PT Wow " + strconv.Itoa(i),
			Phone:   "080988888",
			Address: "Jl Buntu Daamai No." + strconv.Itoa(i),
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	customers, err := custRepoDB.GetAllCustomer()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(customers))

	dBase.DB.Exec("DELETE FROM customers")
}

func TestDeleteAllCustomerDB(t *testing.T) {

	env.NewEnv(os.ExpandEnv("$GOPATH/src/github.com/Saucon/go-rest-api-sample/local/.env"))
	cfg := env.Config

	dBase := db.NewDB(cfg, false)
	dBase.DB.Exec("DELETE FROM customers")
}
