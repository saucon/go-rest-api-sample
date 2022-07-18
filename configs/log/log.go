package log

import (
	"fmt"
	"github.com/Saucon/go-rest-api-sample/configs/models"
	"sync"

	"github.com/olivere/elastic/v7" //nolint:typecheck
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7" //nolint:typecheck
)

type LogCustom struct {
	Logrus *logrus.Logger
	LogDb  *LogDbCustom
	WhoAmI iAm
}

type iAm struct {
	Name string
	Host string
	Port string
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configServer models.ServerConfig) *LogCustom {
	var log *logrus.Logger

	configElstc := configServer.ElasticConfig

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL( //nolint:typecheck
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.User, configElstc.Password)) //nolint:typecheck
	if err != nil {
		selfLogError(err, "config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook( //nolint:typecheck
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "config/log: elastic client", log)
		}
		log.Hooks.Add(hook)
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configServer.Name,
				Host: configServer.Host,
				Port: configServer.Port,
			},
		}
	})
	return instance
}

func (l *LogCustom) Success(req, resp, reqBe, respBe interface{}, description, respTime string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"request":      req,
		"response":     resp,
		"request_be":   reqBe,
		"response_be":  respBe,
	}).Info(description)

	l.LogDb.SuccessLogDb(req, resp, reqBe, respBe, description, respTime, traceHeader)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Info(description string, traceHeader map[string]string, data ...interface{}) {
	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"message":      data,
	}).Info(description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Error(err error, description string, respTime string, traceHeader map[string]string, req interface{},
	resp interface{}, reqBE interface{}, respBE interface{}) {

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
	}).Error(description)

	l.LogDb.ErrorLogDb(err, description, respTime, stFormat, traceHeader, req, resp, reqBE, respBE)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Fatal(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogFatal(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
