package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	//匿名导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

//数据库基本信息
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go-test"
)

//DB 定义数据库指针
var DB *sql.DB

//User user表基本信息
type User struct {
	id           int
	email        string
	role         string
	creatat      int64
	lockat       int64
	password     string
	passwordsalt string
	sessionsalt  string
}

func init() {
	InitDB()
	CreateUser()
}

//InitDB 连接数据库
func InitDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(20)

	if err := DB.Ping(); err != nil {
		log.Println("open database fail")
		return
	}
	log.Println("open database success")
}

//CreateUser 创建用户表
func CreateUser() {
	sqlStr := `CREATE TABLE if not exists User (
					id             INTEGER,
					email          VARCHAR(40),
					role           VARCHAR(40),
					creatat       INT(64),
					lockat        INT(64),
					password       VARCHAR(40),
					passwordsalt  VARCHAR(40),
					sessionsalt   VARCHAR(40),
					PRIMARY KEY (email)
					)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := DB.Exec(sqlStr)
	if err != nil {
		fmt.Println("create table user failed ")
		return
	}
}
