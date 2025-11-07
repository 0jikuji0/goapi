package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

const(
	driver	 = "mysql"
	host	 = "localhost"
	port 	 = 3306
	user 	 = "user_goapi"
	password = "password"
	dbname	 = "goapi_db"
)

// var sqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)


func NewDB() *sql.DB {

    var sqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

    conn, err := sql.Open(driver, sqlInfo);

    if err != nil {
        panic(err.Error())
    }

    fmt.Println("(DATABASE) Connected !")

    return conn
}