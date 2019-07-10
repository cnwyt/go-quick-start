package main

import (
	"fmt"
	"time"
	"log"
	"github.com/uniplaces/carbon"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Hello, World!");

    fmt.Printf("Unix timestamp:  %d \n", time.Now().Unix())
    fmt.Printf("当前时间是:  %s \n", carbon.Now().DateTimeString())

    // 连接数据库
	db, err := sql.Open("mysql", "homestead:secret@tcp(192.168.10.10:3306)/test")
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(3 * time.Second)
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatal("----> 数据库连接失败.")
	    panic(err.Error()) // proper error handling instead of panic in your app
	} else {
		fmt.Println("----> 数据库连接成功.");
	}
	defer db.Close() 
}