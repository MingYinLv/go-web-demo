package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	Conn()
}

func Conn() {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(192.168.1.104)/test")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connect Successify!")
}

func Close() {
	DB.Close()
}
