package db

import (
	// "database/sql"
	"fmt"
	"goapi/models"
)

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}
	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS")
	if err != nil {
		return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
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
	return users, nil
}

func GetUsersByUsername(username string) ([]models.User, error) {
	users := []models.User{}
	rows, err := Conn.Query("SELECT id, username, password, credit FROM goapi_USERS WHERE username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("package DB getUsersByUsername : %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)
		if err != nil {
			return nil, fmt.Errorf("package DB getUsersByUsername : %v", err.Error())
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("package DB getUsersByUsername : %v", err.Error())
	}
	return users, nil
}

func CreateUser(user models.User) error {

	_, err := Conn.Exec("INSERT INTO goapi_USERS (username, password, credit) VALUES (?,?,?)", user.Username, user.Password, user.Credit)
	if err != nil {
		return fmt.Errorf("package db CreateUser : %v", err.Error())
	}
	return nil
}

// package db

// import (
// 	// "database/sql"
// 	"goapi/models"
// 	"fmt"
// )

// func GetAllUsers () ([]models.User, error) {
// 	users := []models.User{}
// 	// users = append(users, models.User{Id : 1, Username: "Youssouf", Password: "secret", Credit: 30})
// 	// users = append(users, models.User{Id : 2, Username: "Paco", Password: "secret", Credit: 70})

// 	rows, err := Conn.Query("SELECT id, username, password, credit, FROM esgi.users")

// 	if err != nil {
// 		return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())

// 	}

// 	defer rows.Close()

// 	for rows.Next(){
// 		var user models.User
// 		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)
// 		if err != nil {
// 			return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())
// 		}
// 		users = append(users, user)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, fmt.Errorf("package DB getAllUsers : %v", err.Error())
// 	}

// 	rows.Close()

// 	return users, nil
// }

// func GetUsersByUsername(username string)([]models.User, error){

// 	users := []models.User{}
// 	// users = append(users, models.User{Id : 1, Username: "Youssouf", Password: "secret", Credit: 30})
// 	// users = append(users, models.User{Id : 2, Username: "Paco", Password: "secret", Credit: 70})

// 	rows, err := Conn.Query("SELECT id, username, password, credit, FROM esgi.users WHERE username = ?")

// 	if err != nil {
// 		return nil, fmt.Errorf("package DB GetUsersByUsername : %v", err.Error())

// 	}

// 	defer rows.Close()

// 	for rows.Next(){
// 		var user models.User
// 		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Credit)
// 		if err != nil {
// 			return nil, fmt.Errorf("package DB GetUsersByUsername : %v", err.Error())
// 		}
// 		users = append(users, user)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		return nil, fmt.Errorf("package DB GetUsersByUsername : %v", err.Error())
// 	}

// 	rows.Close()

// 	return users, nil

// }
