package main

import (
	"encoding/json"
	"fmt"
	"goapi/app"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("School", "esgi")
	res, _ := json.Marshal("en vie")
	fmt.Fprintf(w, "%s", string(res))
}


func main (){
	fmt.Println("bonjour.")

	http.HandleFunc("GET /{$}", healthCheck)
	http.HandleFunc("GET /users/{$}", app.GetAllUsers)
	http.HandleFunc("POST /users/{$}", app.CreateUser)

	fmt.Println("Listen at http://localhost:4242")
	http.ListenAndServe(":4242", nil)
}