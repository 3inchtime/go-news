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
