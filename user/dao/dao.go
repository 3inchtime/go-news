package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"user/utils"
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
	daoConfig := utils.Config
	logrus.Infof("init Mysql config: %+v", daoConfig)
	mysqlConfig := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4",
		daoConfig.Mysql.User,
		daoConfig.Mysql.Pwd,
		daoConfig.Mysql.Host,
		daoConfig.Mysql.Port,
		daoConfig.Mysql.Database,
	)
	logrus.Infof("init Mysql config: %s", mysqlConfig)

	Mysql, err := sqlx.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	Mysql.SetMaxOpenConns(10)
	return Mysql
}
