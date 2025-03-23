package models

import (
	"database/sql"
	"fmt"
	"log"

	"todo_app/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}
	if err := Db.Ping(); err != nil {
		log.Fatalln("DBに接続できませんでした。", err)
	}

	// テーブルが無ければ新規作成
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	_, err := Db.Exec(cmdU)
	if err != nil {
		log.Fatalln("デーブルの作成に失敗しました", err)
	}
}
