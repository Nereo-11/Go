package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:admin@tcp(localhost:3306)/goweb_db"

// Guarda la conexión
var db *sql.DB

// Realiza la conexión
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion Exitosa")
	db = conection
}

// Cierra la conexión
func Close() {
	db.Close()
}

//Verificar conexión

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Verifica si existe tabla

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()

}

//Crea una tabla

func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}

}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	resust, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return resust, err
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}

// Borrar registros en una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

//Registrar registro

//func registrar()
