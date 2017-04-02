package models 

type users struct {
	Username		string 	`db:"username" json:"username"`
    Name            string  `db:"name" json:"name"`
    Password        string  `db:"password" json:"password"`
    Role            int64  `db:"role" json:"role"`

}