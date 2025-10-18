package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("School", "esgi")
	res, _ := json.Marshal("en vie")
	fmt.Fprintf(w, "%s", string(res))
}
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("esgi", "user")
	res, _ := json.Marshal("All users")
	fmt.Fprintf(w, "%s", string(res))
}
func createUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "create user")
}

func main (){
	fmt.Println("bonjour.")

	http.HandleFunc("GET /{$}", healthCheck)
	http.HandleFunc("GET /users/{$}", GetAllUsers)
	http.HandleFunc("POST /users/{$}", createUser)

	fmt.Println("Listen at http://localhost:4242")
	http.ListenAndServe(":4242", nil)
}