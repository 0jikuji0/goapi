package db

import "database/sql"

var Conn *sql.DB

const(
	driver	= "postgres"
	host	= "localhost"
)


func NewDB() *sql.DB{

	conn, err := sql.Open(driver, sqlinfo)
}