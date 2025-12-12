package app

import (
	"encoding/json"
	"fmt"

	"strconv"
	"strings"
	"tpcours/db"
	"tpcours/models"

	// "goapi/models"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := db.GetAllUsers()

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(users)
	fmt.Fprintf(w, "%s", string(res))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Internal server error"}`)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("userId"))

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusBadRequest)
		return
	}

	users, err := db.GetUser(id)

	if users == nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusNotFound)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusInternalServerError)
		return
	}

	var res []byte
	res, err = json.Marshal(users)
	fmt.Fprintf(w, "%s", string(res))

	

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Internal server error"}`)
		return
	}
}

func createValidateUser(userDto models.User) []string {
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
	if strings.Contains(userDto.Username, "ANTOINE") {
		errsMsg = append(errsMsg, "Username must not contain the forbidden word")
	}
	
	existing, err := db.GetUsersByUsername(userDto.Username)
	
	if err != nil {
		errsMsg = append(errsMsg, "Error getting Users by username")
	}
	if len(existing) > 0 {
		errsMsg = append(errsMsg, "Username must be unique")
	}

	return errsMsg
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userDto models.User

	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Incorrect body format", http.StatusBadRequest)
		return
	}

	errMsg := createValidateUser(userDto);

	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	}

	

	err = db.CreateUser(userDto)

	if err != nil {
		http.Error(w, "pb insertion", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateValidateUser(userDto models.User) []string {
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
	if strings.Contains(userDto.Username, "ANTOINE") {
		errsMsg = append(errsMsg, "Username must not contain the forbidden word")
	}

	existing, err := db.GetUsersById(userDto.Id)
	
	if err != nil {
		errsMsg = append(errsMsg, "Error getting Users by id")
	}
	if len(existing) < 1 {
		errsMsg = append(errsMsg, "User not existing")
	}
	
	existing, err = db.GetUsersByUsernameWithoutCurrentUser(userDto.Username, userDto.Id)
	
	if err != nil {
		errsMsg = append(errsMsg, "Error getting Users by username")
	}
	if len(existing) > 0 {
		errsMsg = append(errsMsg, "Username must be unique")
	}

	return errsMsg
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userDto models.User

	id, err := strconv.Atoi(r.PathValue("userId"))

	userDto.Id = id

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Incorrect body format", http.StatusBadRequest)
		return
	}

	errMsg := updateValidateUser(userDto);

	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	}

	
	
	err = db.UpdateUser(userDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Internal server error"}`)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateValidateUserCredit(userDto models.User) []string {
	var errsMsg []string

	if userDto.Credit < 0 {
		errsMsg = append(errsMsg, "Credit cannot be negative")
	}

	existing, err := db.GetUsersById(userDto.Id)
	
	if err != nil {
		errsMsg = append(errsMsg, "Error getting Users by id")
	}
	if len(existing) < 1 {
		errsMsg = append(errsMsg, "User not existing")
	}

	return errsMsg
}

func UpdateUserCredit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userDto models.User

	id, err := strconv.Atoi(r.PathValue("userId"))

	userDto.Id = id

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusBadRequest)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		http.Error(w, "Incorrect body format", http.StatusBadRequest)
		return
	}

	errMsg := updateValidateUserCredit(userDto);

	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), http.StatusBadRequest)
		return
	}

	
	
	err = db.UpdateUserCredit(userDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Internal server error"}`)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteValidateUser(userDto models.User) []string {
	var errsMsg []string

	existing, err := db.GetUsersById(userDto.Id)
	
	if err != nil {
		errsMsg = append(errsMsg, "Error getting Users by id")
	}
	if len(existing) < 1 {
		errsMsg = append(errsMsg, "User not existing")
	}

	return errsMsg
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userDto models.User

	id, err := strconv.Atoi(r.PathValue("userId"))

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "(API) ", http.StatusBadRequest)
		return
	}

	userDto.Id = id

	errMsg := deleteValidateUser(userDto);

	if len(errMsg) > 0 {
		encoded, _ := json.Marshal(errMsg)
		http.Error(w, string(encoded), http.StatusNotFound)
		return
	}


	
	err = db.DeleteUser(userDto)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "Internal server error"}`)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}