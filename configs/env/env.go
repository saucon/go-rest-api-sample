package env

import (
	"fmt"
	"github.com/Saucon/go-rest-api-sample/configs/models"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Config models.ServerConfig

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewEnv(filenames ...string) {
	err := loadConfig(filenames...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error_cause":   PrintErrorStack(err),
			"error_message": err.Error(),
		}).Fatal("config/env: load config")
	}
}

func loadConfig(filenames ...string) (err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		logrus.Fatal(err, " config/env: load gotdotenv")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.DBConfig)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.ElasticConfig)
	if err != nil {
		return err
	}

	return err
}

func PrintErrorStack(err error) string {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	return stFormat
}
