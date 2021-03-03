package internal

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	Driver      string
	Host        string
	Port        int
	User        string
	Password    string
	Lib         string
	MaxIdleConn int
	MaxOpenConn int
}

func NewDb(conf DatabaseConf) (*sql.DB, error) {
	var (
		err error
		dsn string
	)
	switch conf.Driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Lib)
	default:
		return nil, errors.New("unsupported database driver: " + conf.Driver)
	}
	db, err := sql.Open(conf.Driver, dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if conf.MaxIdleConn > 0 {
		db.SetMaxIdleConns(conf.MaxIdleConn)
	}
	if conf.MaxOpenConn > 0 {
		db.SetMaxOpenConns(conf.MaxOpenConn)
	}
	return db, nil
}

var (
	gDb *sql.DB
)

func InitDatabase(conf DatabaseConf) {
	var err error
	gDb, err = NewDb(conf)
	if err != nil {
		panic(err)
	}
}

func Db() *sql.DB {
	return gDb
}
