package db

import (
	// "database/sql"
	"goapi/models"
	"fmt"
)

func GetAllUsers () ([]models.User, error) {
	users := []models.User{}
	// users = append(users, models.User{Id : 1, Username: "Youssouf", Password: "secret", Credit: 30})
	// users = append(users, models.User{Id : 2, Username: "Paco", Password: "secret", Credit: 70})

	rows, err := Conn.Query("SELECT id, username, password, credit, FROM esgi.users")

	if err != nil {
		return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())

	}

	defer rows.Close()

	for rows.Next(){
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)
		if err != nil {
			return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())
	}

	rows.Close()

	return users, nil

}