package dao

import (
	"github.com/sirupsen/logrus"
	"time"
	"user/model"
)

func (d *Dao) CreateUser (u *model.User) (rows int, err error) {
	createTime := time.Now().Unix()
	insertSQL, err := d.DB.Prepare("INSERT INTO user (" +
		"user_name, account, password, telephone, email, age, create_time, update_time" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Prepare Insert SQL Error: %s", err.Error())
	}

	res, err := insertSQL.Exec(u.UserName, u.Account, u.HashPassword, u.Telephone, u.Email, u.Age, createTime, createTime)
	if err != nil {
		logrus.Errorf("Insert Video Info SQL Error: %s", err.Error())
		return 0, err
	}
	defer insertSQL.Close()
	id, _ := res.LastInsertId()
	return int(id), nil
}
