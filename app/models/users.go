package models 

type Users struct {
	Id 				int64 	`db:"id" json:"id"`
	Username		string 	`db:"username" json:"username"`
    Name            string  `db:"name" json:"name"`
    Password        string  `db:"password" json:"password"`
    Role            int64  `db:"role" json:"role"`

}