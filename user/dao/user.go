package dao

import (
	"database/sql"
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

func (d *Dao) GetUserInfoByAccount (account string) (*model.User, error) {
	querySQL, err := d.DB.Prepare("SELECT id, user_name, account, password FROM user WHERE account = ?")
	if err != nil {
		logrus.Errorf("Prepare Query SQL Error: %s", err.Error())
	}
	u := new(model.User)
	err = querySQL.QueryRow(account).Scan(&u.ID, &u.UserName, &u.Account, &u.HashPassword)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query PurcaseDetail Error: %s", err.Error())
		return nil, err
	}
	return u, nil
}