package db

import (
	// "database/sql"
	"database/sql"
	"fmt"
	"tpcours/models"
)

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	var rows *sql.Rows

	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS")
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USERS | %v", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)

		if err != nil {
			return nil, fmt.Errorf("(DATABASE) | GET USERS | %v", err.Error())
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USERS | %v", err.Error())
	}

	return users, nil
}
func GetUser(id int) (*models.User, error) {
	user := models.User{}

	var rows *sql.Rows

	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USER | %v", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		
		
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)

		if err != nil {
			return nil, fmt.Errorf("(DATABASE) | GET USER | %v", err.Error())
		}


	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USERS | %v", err.Error())
	}

	return &user, nil
}

func GetUsersByUsername(username string) ([]models.User, error) {
	users := []models.User{}

	var rows *sql.Rows

	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS WHERE username = ?", username)
	if err != nil {
		
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)

		if err != nil {
			return nil, fmt.Errorf("(DATABASE) | GET USER BY USERNAME | %v", err.Error())
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USER BY USERNAME | %v", err.Error())
	}

	return users, nil
}

func GetUsersById(id int) ([]models.User, error) {
	users := []models.User{}

	var rows *sql.Rows

	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USER BY ID | : %v", err.Error())
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)

		if err != nil {
			return nil, fmt.Errorf("(DATABASE) | GET USER BY ID | %v", err.Error())
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USER BY ID | %v", err.Error())
	}

	return users, nil
}

func CreateUser(user models.User) error {
	_, err := Conn.Exec("INSERT INTO goapi_USERS (username, password, credit) VALUES (?, ?, ?)", user.Username, user.Password, user.Credit)

	if err != nil {
		return fmt.Errorf("(DATABASE) | POST USER | %v", err.Error())
	}

	return nil
}

func GetUsersByUsernameWithoutCurrentUser(username string, id int) ([]models.User, error) {
	users := []models.User{}

	var rows *sql.Rows

	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS WHERE username = ? AND id <> ?", username, id)
	if err != nil {
		
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)

		if err != nil {
			return nil, fmt.Errorf("(DATABASE) | GET USER BY USERNAME WITHOUT CURRENT USERNAME | %v", err.Error())
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("(DATABASE) | GET USER BY USERNAME WITHOUT CURRENT USERNAME | %v", err.Error())
	}

	return users, nil
}

func UpdateUser(user models.User) error {
	_, err := Conn.Exec("UPDATE goapi_USERS SET username = ?, password = ?, credit = ? WHERE id = ?", user.Username, user.Password, user.Credit, user.Id)

	if err != nil {
		return fmt.Errorf("(DATABASE) | PUT USER | %v", err.Error())
	}

	return nil
}


func UpdateUserCredit(user models.User) error {
	_, err := Conn.Exec("UPDATE goapi_USERS SET credit = ? WHERE id = ?", user.Credit, user.Id)

	if err != nil {
		return fmt.Errorf("(DATABASE) | PATCH USER CREDIT | %v", err.Error())
	}

	return nil
}

func DeleteUser(user models.User) error {
	_, err := Conn.Exec("DELETE FROM goapi_USERS WHERE id = ?", user.Id)

	if err != nil {
		return fmt.Errorf("(DATABASE) | DELETE USER | %v", err.Error())
	}

	return nil
}
