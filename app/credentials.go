package app

import (
	"encoding/json"
	"fmt"
	"tpcours/db"
	"tpcours/models"
	"tpcours/utils"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

    var creds models.Credentials
    err := json.NewDecoder(r.Body).Decode(&creds)

    if err != nil { 
        http.Error(w, "Invalid JSON", http.StatusBadRequest) 
        return 
	}

    existing, err := db.GetUsersByUsername(creds.Username) 
    if err != nil { 
        http.Error(w, "Error requesting user by username", http.StatusInternalServerError) 
        return 
    }

    if existing == nil || creds.Username != existing[0].Username || creds.Password != existing[0].Password { 
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