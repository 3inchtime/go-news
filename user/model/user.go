package model

type User struct {
	ID        int
	UserName  string
	Telephone string
	Email     string
	Age       string
}

type UserLogin struct {
	ID       int
	PassWord string
}
