package dao

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"user/model"
)

func (d *Dao) CreateUser (u *model.User, ua *model.UserAccountInfo) error {
	tx, err := d.DB.Begin()
	if err != nil {
		logrus.Errorf("Begin Tx Error: %s\n", err.Error())
		return err
	}

	accountSQL := "INSERT INTO ts_account (user_id, account, password, create_time, update_time) VALUES (?,?,?,?,?)"
	_, err = tx.Exec(accountSQL, ua.UserID, ua.Account, ua.Password, ua.CreateTime, ua.UpdateTime)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("accountSQL error: %s", err.Error())
		return err
	}

	userSQL := "INSERT INTO ts_user (user_id, create_time, update_time) VALUES (?,?,?)"
	_, err = tx.Exec(userSQL, u.UserID, u.CreateTime, u.UpdateTime)
	if err != nil {
		tx.Rollback()
		logrus.Errorf("userSQL error: %s", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("Commit TX error: %s", err.Error())
		return err
	}
	return nil
}


func (d *Dao) CheckUserPwd (account, password string) (string, error) {
	querySQL, err := d.DB.Prepare("SELECT user_id FROM ts_account WHERE account = ? and password = ?")
	if err != nil {
		logrus.Errorf("Prepare Query SQL Error: %s", err.Error())
	}
	var userID string
	err = querySQL.QueryRow(account, password).Scan(&userID)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query PurcaseDetail Error: %s", err.Error())
		return "", err
	}
	return userID, nil
}
