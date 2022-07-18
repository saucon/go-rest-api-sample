package main

import (
	"github.com/Saucon/go-rest-api-sample/configs/db"
	"github.com/Saucon/go-rest-api-sample/configs/env"
	"github.com/Saucon/go-rest-api-sample/configs/gin"
	"github.com/Saucon/go-rest-api-sample/configs/log"
	"github.com/Saucon/go-rest-api-sample/configs/models"
	"github.com/Saucon/go-rest-api-sample/internal/customer/controller"
	"github.com/Saucon/go-rest-api-sample/internal/customer/repos/repodb"
	"github.com/Saucon/go-rest-api-sample/internal/customer/usecase"
	"os"
)

func main() {
	// get env
	env.NewEnv(os.ExpandEnv("$GOPATH/src/github.com/Saucon/go-rest-api-sample/local/.env.local"))
	cfg := env.Config

	// init db
	dBase := db.NewDB(cfg, false)

	// init db log
	dBaseLog := db.NewDB(cfg, true)
	//dBaseLog.DropTable(models.Logs{})
	dBaseLog.AutoMigrate(models.Logs{})

	// set log DB to log custom
	logger := log.NewLogCustom(cfg)
	loggerDb := log.NewLogDbCustom(dBaseLog.DB)
	logger.LogDb = loggerDb

	// repositories
	custRepoDb := repodb.NewCustomerRepoDB(dBase)

	//usecase
	custUsecase := usecase.NewCustomerUsecase(custRepoDb, logger)

	//controller
	custHandler := controller.NewCustomerHandler(custUsecase, logger)

	router := gin.NewRouter(custHandler)
	if err := router.Run(cfg.Host + ":" + cfg.Port); err != nil {
		logger.Fatal(err, "main : run router", nil)
	}
}
