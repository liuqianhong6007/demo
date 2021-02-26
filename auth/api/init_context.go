package api

import (
	"context"
	"database/sql"
	"io/ioutil"

	"github.com/liuqianhong6007/demo/auth/config"
	"github.com/liuqianhong6007/demo/auth/internal"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func read(filename string) string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(buf)
}

func createTableIfNotExist(ctx context.Context, db *sql.DB, libName, tableName string, createTableSql string) {
	_, err := db.ExecContext(ctx, createTableSql)
	checkErr(err)
}

func CheckContext(ctx context.Context) {
	// 账号表
	createTableIfNotExist(ctx, internal.Db(), config.DbLib(), "account", read("sql/account.sql"))
	// 邀请码表
	createTableIfNotExist(ctx, internal.Db(), config.DbLib(), "invite_code", read("sql/invite_code.sql"))
}
