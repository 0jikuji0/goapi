
package app

import (
	"encoding/json"
	"fmt"
	"tpcours/db"
	"tpcours/models"
	"strings"
	"tpcours/utils"
	// "goapi/models"
	"net/http"

)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("school", "esgi")

	users, err := db.GetAllUsers()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "erreur de récupération des users", http.StatusInternalServerError)
	}

	res, _ := json.Marshal(users)
	fmt.Fprintf(w, "%s", string(res))
}

func Login(w http.ResponseWriter, r *http.Request) { 
	var creds models.Credentials 
	err := json.NewDecoder(r.Body).Decode(&creds) 
	if err != nil { 
		http.Error(w, "Invalid JSON", http.StatusBadRequest) 
		return 
	} 
	existing, err := db.GetUserByUsername(creds.Username) 
	if err != nil { 
		http.Error(w, "Error requesting user by username", http.StatusInternalServerError) 
		return 
	}
	if existing == nil || creds.Username != existing.Username || creds.Password != existing.Password { 
		http.Error(w, "Unauthorized", http.StatusUnauthorized) 
		return 
	} 
	token, err := utils.GenerateJWT(creds.Username) 
	if err != nil { 
		http.Error(w, "Error generating token", http.StatusInternalServerError) 
		return 
	} 
	encodedToken, _ := json.Marshal(token) 
	fmt.Fprintf(w, "%s", encodedToken)
}

func ValidateUser(userDto models.User) []string {
	// pour faire une regex
	//match, _ := regexp.Match("e([a-z]+)gi", []byte(userDto.Username))
	var errsMsg []string

	if len(userDto.Username) < 5 {
		errsMsg = append(errsMsg, "username length must be at least 5")
	}
	if len(userDto.Password) < 6 {
		errsMsg = append(errsMsg, "password length must be at least 6")
	}
	if userDto.Credit < 0 {
		errsMsg = append(errsMsg, "Credit cannot be negative at creation")
	}
	if !strings.ContainsAny(userDto.Password, "!-$+/") {
		errsMsg = append(errsMsg, "Password must contain at least 1 special char (!-$+/)")
	}
	if strings.Contains(userDto.Username, "TRANCHO") {
		errsMsg = append(errsMsg, "Username must not contain the forbidden word")
	}
	return errsMsg
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto models.User

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Incorrect body format", http.StatusBadRequest)
		return
	}

	errsMsg := ValidateUser(userDto)

	if len(errsMsg) > 0 {
		encoded, _ := json.Marshal(errsMsg)
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	}

	existing, err := db.GetUsersByUsername(userDto.Username)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error getting Users by username", http.StatusInternalServerError)
		return
	}

	if len(existing) > 0 {
		http.Error(w, "Username must be unique", http.StatusConflict)
		return
	}

	err = db.CreateUser(userDto)
	if err != nil {
		http.Error(w, "pb d'insertion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}







// package app

// import (
// 	"encoding/json"
// 	"fmt"
// 	"goapi/db"
// 	"goapi/models"
// 	"strings"

// 	// "goapi/models"
// 	"net/http"
// )


// func GetAllUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Add("school", "esgi")

// 	users, err := db.GetAllUsers()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		http.Error(w, "erreur de récupération des users", http.StatusInternalServerError)
// 	}

// 	res, _ := json.Marshal(users)
// 	fmt.Fprintf(w, "%s", string(res))
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	var userDto models.User

// 	err := json.NewDecoder(r.Body).Decode(&userDto)
// 	if err != nil {
// 		http.Error(w, "Incorrect body format", http.StatusBadRequest)
// 		return
// 	}

// 	// pour faire une regex
// 	//match, _ := regexp.Match("e([a-z]+)gi", []byte(userDto.Username))
// 	var errsMsg []string

// 	if len(userDto.Username) < 5 {
// 		errsMsg = append(errsMsg, "username length must be at least 5")
// 	}
// 	if len(userDto.Password) < 6 {
// 		errsMsg = append(errsMsg, "password length must be at least 6")
// 	}
// 	if userDto.Credit < 0 {
// 		errsMsg = append(errsMsg, "Credit cannot be negative at creation")
// 	}
// 	if !strings.ContainsAny(userDto.Password, "!-$+/") {
// 		errsMsg = append(errsMsg, "Password must contain at least 1 special char (!-$+/)")
// 	}
// 	if strings.Contains(userDto.Username, "TRANCHO") {
// 		errsMsg = append(errsMsg, "Username must not contain the forbidden word")
// 	}

// 	if len(errsMsg) > 0 {
// 		encoded, _ := json.Marshal(errsMsg)
// 		http.Error(w, string(encoded), http.StatusBadRequest)
// 		return
// 	}

// 	// si username est pas unique
// 	// http.Error(w, "username not unique", http.StatusConflict)

// 	existing ,err := db.GetUsersByUsername(userDto.Username)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		http.Error(w, "Error getting Users by username", http.StatusInternalServerError)
// 		return
// 	}

// 	if len(existing) > 0{
// 		http.Error(w, "Username must be unique", http.StatusConflict)
// 		return
// 	}
// 	// si pb insertion
// 	// http.Error(w, "pb d'insertion", http.StatusInternalServerError)

// 	w.WriteHeader(http.StatusCreated)
// }
