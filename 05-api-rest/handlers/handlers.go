package handlers

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	//Obtener registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	//Obtener registro
	user := models.User{}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		db.Connect()
		user.Save()
		db.Close()
	}

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	//Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	db.Connect()
	user, _ := models.GetUser(userId)
	user.Delete()
	db.Close()

	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func getUserByRequest(r *http.Request) (models.User, error) {
	//Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}
