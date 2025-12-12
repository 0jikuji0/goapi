package main

import (
	"encoding/json"
	"fmt"
	"tpcours/app"
	"net/http"
	"tpcours/db"
	"tpcours/utils"
)


func healthCheck(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	_, err := utils.VerifyJWT(tokenString)
    if err != nil {
        http.Error(w, "Invalid token", http.StatusUnauthorized)
        return
    }
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("school", "esgi")
	res, _ := json.Marshal("en vie.")
	fmt.Fprintf(w, "%s", string(res))
}

func main() {

	db.Conn = db.NewDB()

	http.HandleFunc("GET /health/{$}", healthCheck)
	http.HandleFunc("GET /users/", app.GetAllUsers)
	http.HandleFunc("POST /users/{$}", app.CreateUser)
	http.HandleFunc("POST /login/{$}", app.Login)
	
	// http.HandleFunc("GET /users/{userId}", app.GetUserById)
	// http.HandleFunc("DELETE /users/{userId}", app.DeleteUser)
	// http.HandleFunc("PUT /users/{userId}", app.UpdateUser)

	fmt.Println("Listening at http://localhost:4242")
	http.ListenAndServe(":4242", nil)
}


