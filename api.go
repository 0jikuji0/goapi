package main

import (
	"encoding/json"
	"fmt"
	"goapi/app"
	"net/http"
	"goapi/db"
)


func healthCheck(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("school", "esgi")
	res, _ := json.Marshal("en vie.")
	fmt.Fprintf(w, "%s", string(res))
}

func main() {

	db.Conn = db.NewDB()

	http.HandleFunc("GET /{$}", healthCheck)
	http.HandleFunc("GET /users/{$}", app.GetAllUsers)
	http.HandleFunc("POST /users/{$}", app.CreateUser)
	//
	//http.HandleFunc("GET /users/{userId}", app.GetUserById)
	//http.HandleFunc("DELETE /users/{userId}", app.DeleteUser)
	//http.HandleFunc("PUT /users/{userId}", app.UpdateUser)

	//http.HandleFunc("GET /books/{$}", app.GetAllBooks)
	//http.HandleFunc("POST /books/{$}", app.CreateBooks)

	fmt.Println("Listening at http://localhost:4242")
	http.ListenAndServe(":4242", nil)
}