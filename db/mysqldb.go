package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconn *gorm.DB
var err error

func Init() {

	var (
		UserName     string = os.Getenv("MYSQL_USER")
		Password     string = os.Getenv("MYSQL_PASSWD")
		Addr         string = os.Getenv("MYSQL_HOST")
		Port         string = os.Getenv("MYSQL_PORT")
		Database     string = os.Getenv("MYSQL_DB")
		MaxLifetime  int    = 10
		MaxOpenConns int    = 10
		MaxIdleConns int    = 10
	)

	fmt.Println("DB Init start")
	port, _ := strconv.Atoi(Port)
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", UserName, Password, Addr, port, Database)
	fmt.Println("addr : " + addr)
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	sqldb, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	sqldb.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	sqldb.SetMaxIdleConns(MaxIdleConns)
	sqldb.SetMaxOpenConns(MaxOpenConns)
	fmt.Println("DB Init end")

	DBconn = conn
}

func GetDb() *gorm.DB {
	return DBconn
}
