package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConf struct {
	driver      string
	host        string
	port        int
	user        string
	password    string
	lib         string
	maxIdleConn int
	maxOpenConn int
}

func NewDb(conf DatabaseConf) (*sql.DB, error) {
	var (
		err error
		dsn string
	)
	switch conf.driver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.password, conf.host, conf.port, conf.lib)
	default:
		return nil, WrapError{ErrMsg: "unsupported database driver: " + conf.driver}
	}
	db, err := sql.Open(conf.driver, dsn)
	if err != nil {
		return nil, WrapError{Err: err}
	}
	if conf.maxIdleConn > 0 {
		db.SetMaxIdleConns(conf.maxIdleConn)
	}
	if conf.maxOpenConn > 0 {
		db.SetMaxOpenConns(conf.maxOpenConn)
	}
	err = db.Ping()
	if err != nil {
		return nil, WrapError{Err: err}
	}
	return db, nil
}

var (
	gDb *sql.DB
)

func InitDatabase(driver, host string, port int, user, password, lib string, maxIdleConn, maxOpenConn int) {
	var err error
	gDb, err = NewDb(DatabaseConf{
		driver:      driver,
		host:        host,
		port:        port,
		user:        user,
		password:    password,
		lib:         lib,
		maxIdleConn: maxIdleConn,
		maxOpenConn: maxOpenConn,
	})
	if err != nil {
		panic(err)
	}
}

func Db() *sql.DB {
	return gDb
}
