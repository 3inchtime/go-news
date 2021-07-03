package model

type User struct {
	UserID     string
	UserName   string
	Telephone  string
	Email      string
	Age        int
	CreateTime int64
	UpdateTime int64
}

type UserAccountInfo struct {
	UserID     string
	Account    string
	Password   string
	CreateTime int64
	UpdateTime int64
}
