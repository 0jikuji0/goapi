package app

import (
	"encoding/json"
	"fmt"
	"goapi/db"
	"goapi/models"
	"strings"

	// "goapi/models"
	"net/http"
)


func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("esgi", "user")

	users, err := db.GetAllUsers()
	if err != nil{
		fmt.Println(err.Error())
		http.Error(w, "erreur de recuperation des users", http.StatusInternalServerError)
	}
	
	res, _ := json.Marshal(users)
	fmt.Fprintf(w, "%s", string(res))
}
func CreateUser(w http.ResponseWriter, r *http.Request){

	var userDto models.User

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Error dans le format du user", http.StatusBadRequest)
		return
	}

	var errMsg []string

	if len(userDto.Username) < 5 {
		errMsg = append(errMsg, "username length must be at least 5", http.StatusBadRequest)
		return
	}
	if len(userDto.Password) < 6 {
		errMsg = append(errMsg, "Password length must be at least 5", http.StatusBadRequest)
		return
	}

	if userDto.Credit < 0 {
		errMsg = append(errMsg, "username length must be at least 5", http.StatusBadRequest)
		return
	}

	if !strings.ContainsAny(userDto.Password, "!-£+/@"){
		errMsg = append(errMsg, "Password must contain at least 1 special char (!-£+/@)", http.StatusBadRequest)
		return
	}
	if !strings.ContainsAny(userDto.Password, "ANTOINE"){
		http.Error(w, "Password must contain at least 1 special char (ANTOINE)", http.StatusBadRequest)
		return
	}

	//si username est unique
	http.Error(w, "username not unique", http.StatusConflict)
	// err := db.CreateUser(userDto)
	//si pb insertion 
	//http.Error(w)

}

