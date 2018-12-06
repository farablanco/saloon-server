// author : Kenneth Phang
package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func ConnectGORM() *gorm.DB {
	DBMS := os.Getenv("DB_DIALET")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DB_DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	db.LogMode(true)
	if err != nil {
		panic(err.Error())
	}

	return db
}
