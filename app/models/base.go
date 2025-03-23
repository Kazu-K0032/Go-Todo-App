package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"todo_app/config"

	"github.com/google/uuid"
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

// 一意のUUIDを生成する関数
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// 文字をSHA-1ハッシュし、16進数の文字列として返す関数
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
