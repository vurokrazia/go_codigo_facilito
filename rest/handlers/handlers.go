package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetUsers())
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi((vars["id"]))

	if user, err := models.GetUser(userId); err != nil {
		models.SendNotFound(w)
	} else {
		models.SendData(w, user)
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		models.SendData(w, models.SaveUser(user))
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "D")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "E")
}
