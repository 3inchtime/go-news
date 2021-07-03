package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	DB *sqlx.DB
}

func NewDao() *Dao {
	return &Dao{
		DB: NewDB(),
	}
}

func NewDB() *sqlx.DB {
	Mysql, err := sqlx.Open("mysql", "root:123456@tcp(192.168.1.103:3306)/user?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxOpenConns(10)
	return Mysql
}
