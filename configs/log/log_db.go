package log

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
	"sync"
)

type LogDbCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
	Db     *gorm.DB
}

var instanceDb *LogDbCustom
var onceDb sync.Once

func NewLogDbCustom(db *gorm.DB) *LogDbCustom {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	sqlDb, err := db.DB()

	if err != nil {
		//logger.Error(err, "db/NewLogDbCustom", nil, nil, nil)
		fmt.Println("adsgf")
	}

	hook := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		level := entry.Level.String()
		if level == "info" {
			level = "success"
		}

		err = db.Debug().Exec("INSERT INTO logs(level, message, path_error, trace_header, request, response, request_be, response_be, created_at, response_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			level, entry.Message, entry.Data["error_cause"], entry.Data["trace_header"], entry.Data["request"], entry.Data["response"], entry.Data["request_be"], entry.Data["response_be"], entry.Time, entry.Data["response_time"]).Error
		if err != nil {
			err := sqlDb.Rollback()
			if err != nil {
				return err
			}
		}
		return err
	}

	log.AddHook(hook)

	onceDb.Do(func() {
		instanceDb = &LogDbCustom{
			Logrus: log,
		}
		instance.LogDb = instanceDb
	})

	return instanceDb
}

func (l *LogDbCustom) ErrorLogDb(err error, description, respTime, strFormat string, traceHeader map[string]string, req, resp, reqBE, respBE interface{}) {
	err = errors.WithStack(err)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   strFormat,
		"error_message": err.Error(),
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Error(description)
}

func (l *LogDbCustom) SuccessLogDb(req, resp, reqBE, respBE interface{}, description, respTime string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Info(description)
}
