package models

import "gomysql/db"

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}
type Users []User

const UserSchema = `CREATE TABLE users (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(100),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// Construir Usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

//Crear Usuario e insertar
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Insertar registro a base de datos
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

//Listar registros
func ListUsers() Users {
	sql := ("SELECT id, username, password, email FROM users")
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

//Obtener registro

func GetUser(id int) *User {
	user := NewUser("", "", "")

	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user
}


//Actualizar registro
func (user *User) Update(){
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

//Guardar o editar registro
func (user *User) Save(){
	if user.Id == 0 {
		user.insert()
	} else{
		user.Update()
	}
}

//Eliminar registro
func (user *User) Delet(){
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
}