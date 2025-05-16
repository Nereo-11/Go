package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realiza la conexión
var dsn = "root:admin@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexión", err)
		panic(err)
	} else {
		fmt.Println("Conexión exitosa")
		return db
	}
}()
