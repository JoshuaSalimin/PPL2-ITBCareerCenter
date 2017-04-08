package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type Users struct {
	UserId 			int64 	`db:"userid"`
	Username		string 	`db:"username"`		
    Password        string  `db:"password"`
    Name            string  `db:"name"`		
	ProductName		string 	`db:"product_name"`
    CompanyName		string  `db:"company_name"`
    Description		string	`db:"description"`
    CreatedAt		int64	`db:"users_created_at"`
    UpdatedAt		int64	`db:"users_updated_at"`		
    ShowProfile		bool 	`db:"show_profile"`
    Role            int  	`db:"role"`
}

func CreateDefaultUser(username string) Users {
    user_dummy := Users{
        UserId          : 0,   
        Username        : username,  
        Password        : "password",  
        Name            : "DEFAULT_NAME",  
        ProductName     : "DEFAULT_PRODUCT_NAME",  
        CompanyName     : "DEFAULT_COMPANY_NAME",  
        Description     : "DEFAULT_DESCRIPTION",  
        CreatedAt       : time.Now().UnixNano(), 
        UpdatedAt       : 0,   
        ShowProfile     : false,    
        Role            : 0,     
    }
    return user_dummy
}

