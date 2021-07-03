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


func (d *Dao) CheckUserPwd (account string) (*model.UserAccountInfo, error) {
	querySQL, err := d.DB.Prepare("SELECT user_id, password FROM ts_account WHERE account = ?")
	if err != nil {
		logrus.Errorf("Prepare Query SQL Error: %s", err.Error())
	}
	ua := new(model.UserAccountInfo)
	err = querySQL.QueryRow(account).Scan(&ua.UserID, &ua.Password)

	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query PurcaseDetail Error: %s", err.Error())
		return nil, err
	}
	logrus.Infof("Query account result: %+v",*ua)
	return ua, nil
}

func (d *Dao) QueryUserInfo (userID string) (*model.User, error) {
	querySQL, err := d.DB.Prepare("SELECT * FROM ts_user WHERE user_id = ?")
	if err != nil {
		logrus.Errorf("Prepare query sql error: %s", err.Error())
	}
	u := new(model.User)
	err = querySQL.QueryRow(userID).Scan(&u.UserID, &u.UserName, &u.Telephone, &u.Email, &u.Age, &u.CreateTime, &u.UpdateTime)

	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query user info rror: %s", err.Error())
		return nil, err
	}
	logrus.Infof("Query account result: %+v",*u)
	return u, nil
}