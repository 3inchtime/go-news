package model

type User struct {
	ID           int
	UserName     string
	Account      string
	HashPassword string
	Telephone    string
	Email        string
	Age          int
	CreateTime   int
	UpdateTime   int
}

type UserBaseInfo struct {
	UserName     string `json:"user_name"`
	Account      string `json:"account"`
	HashPassword string `json:"hash_password"`
}
