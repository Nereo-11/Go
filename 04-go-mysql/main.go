package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")
	db.TruncateTable("users")
	//	db.Ping()
	db.Close()
}
