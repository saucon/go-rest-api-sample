package db

import (
	"github.com/Saucon/go-rest-api-sample/configs/log"
	"github.com/Saucon/go-rest-api-sample/configs/models"
	//nolint:typecheck

	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Database struct {
	DB *gorm.DB
	l  *log.LogCustom
}

func NewDB(conf models.ServerConfig, isDbLog bool) *Database {
	var DB *gorm.DB
	var err error

	var host, user, password, name, port string

	l := log.NewLogCustom(conf)

	defer func() {
		if r := recover(); r != nil {
			l.Error(errors.New("recover"), "config/db: recover from error db init", "",
				nil, nil, nil, nil, nil)
		}
	}()

	// check DB version
	if isDbLog {
		host = conf.DBConfig.HostLogDb
		port = conf.DBConfig.PortLogDb
		user = conf.DBConfig.UserLogDb
		password = conf.DBConfig.PasswordLogDb
		name = conf.DBConfig.NameLogDb
	} else {
		host = conf.DBConfig.Host
		port = conf.DBConfig.Port
		user = conf.DBConfig.User
		password = conf.DBConfig.Password
		name = conf.DBConfig.Name
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, name, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		l.Fatal(err, "config/db: gorm open connect", nil)
	}

	dbSQL, err := DB.DB()
	if err != nil {
		l.Fatal(err, "config/DB: gorm open connect", nil)
	}

	//Database Connection Pool
	dbSQL.SetMaxIdleConns(10)
	dbSQL.SetMaxOpenConns(100)
	dbSQL.SetConnMaxLifetime(time.Hour)

	err = dbSQL.Ping()
	if err != nil {
		l.Fatal(err, "config/DB: can't ping the DB, WTF", nil)
	} else {
		go doEvery(10*time.Minute, pingDb, DB, l)
		return &Database{
			DB: DB,
			l:  l,
		}
	}

	return &Database{
		DB: DB,
		l:  l,
	}
}

func doEvery(d time.Duration, f func(*gorm.DB, *log.LogCustom), x *gorm.DB, y *log.LogCustom) {
	for range time.Tick(d) {
		f(x, y)
	}
}

func pingDb(db *gorm.DB, l *log.LogCustom) {
	dbSQL, err := db.DB()
	if err != nil {
		l.Error(err, "config/db: can't ping the db, WTF", "", nil, nil, nil, nil, nil)
	}

	err = dbSQL.Ping()
	if err != nil {
		l.Error(err, "config/db: can't ping the db, WTF", "", nil, nil, nil, nil, nil)
	}
}

func (d *Database) AutoMigrate(schemas ...interface{}) {
	for _, schema := range schemas {
		if err := d.DB.AutoMigrate(schema); err != nil {
			d.l.Error(err, "", "", nil, nil, nil, nil, nil)
		}
	}
}

func (db *Database) DropTable(schemas ...interface{}) error {
	for _, schema := range schemas {

		if err := db.DB.Migrator().DropTable(schema); err != nil {
			db.l.Error(err, "", "", nil, nil, nil, nil, nil)
			return err
		}
	}
	return nil
}
