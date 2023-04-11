package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dbCon, err := sql.Open("mysql", "root:00000000@tcp(localhost:3306)/nest_photo")
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}

	db = dbCon

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(3)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(3)

}

func Db() *sql.DB {
	return db
}
