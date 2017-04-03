package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"

)

type Users struct {
	UserId 			int64 	`db:"userid"`
	Username		string 	`db:"username"`
    Name            string  `db:"name"`
    Password        string  `db:"password"`
    Role            int  	`db:"role"`
    CreatedAt		int64	`db:"users_created_at"`
    ShowProfile		bool 	`db:"show_profile"`
}

