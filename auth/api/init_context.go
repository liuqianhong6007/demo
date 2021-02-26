package api

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/liuqianhong6007/demo/auth/config"
	"github.com/liuqianhong6007/demo/auth/internal"
	"io/ioutil"
)

func checkTableExistSql(libName, tableName string) string {
	return fmt.Sprintf(`select count(1) from information_schema.TABLES where TABLE_SCHEMA = '%s' and table_name = '%s'`, libName, tableName)
}

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
	rows, err := db.QueryContext(ctx, checkTableExistSql(libName, "account"))
	checkErr(err)
	defer rows.Close()

	rows.Next()

	var count int
	err = rows.Scan(&count)
	checkErr(err)

	if count == 0 {
		_, err = db.ExecContext(ctx, createTableSql)
		checkErr(err)
	}
}

func CheckContext() {
	// 账号表
	createTableIfNotExist(context.Background(), internal.Db(), config.DbLib(), "account", read("sql/account.sql"))
	// 邀请码表
	createTableIfNotExist(context.Background(), internal.Db(), config.DbLib(), "invite_code", read("sql/invite_code.sql"))
}
