package models

import "gorm/db"

type User struct { //Si quiero datos tipo XML:`xml:"id"`      //Si quiero datos tipo yaml:`yaml:"id"`
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type Users []User

func MigrarUser() {
	db.Database.AutoMigrate(User{})
}
