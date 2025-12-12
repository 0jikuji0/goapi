package models

type User struct {
	Id int				`json:"id"`
	Username string		`json:"username"`
	Password string		`json:"password"`
	Credit int			`json:"credit"`
}
type Credentials struct { 
	Username string `json:"username"` 
	Password string `json:"password"`
}