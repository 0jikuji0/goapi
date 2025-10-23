package app

import (
	"encoding/json"
	"fmt"
	"goapi/db"

	// "goapi/models"
	"net/http"
)


func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("esgi", "user")

	users, err := db.GetAllUsers()
	if err != nil{
		println("pas content")
	}
	
	res, _ := json.Marshal(users)
	fmt.Fprintf(w, "%s", string(res))
}
func CreateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "create user")
}