package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type Users struct {
	UserId 			        int64 	`db:"userid"`
	Username		        string 	`db:"username"`	// unique
    Password                string  `db:"password"`
    Name                    string  `db:"name"`		
	ProductName		        string 	`db:"product_name"`
    CompanyName		        string  `db:"company_name"`
    CompanyDescription		string	`db:"description"`
    Visi                    string  `db:"visi"`
    Misi                    string  `db:"misi"`
    Jurusan                 string  `db:"jurusan"`
    Angkatan                int  `db:"angkatan"`
    LogoPath				string `db:"logo_path"`
    CreatedAt		        int64	`db:"users_created_at"`
    UpdatedAt		        int64	`db:"users_updated_at"`		
    ShowProfile		        bool 	`db:"show_profile"`  // show profile ato ngga
    Role                    int  	`db:"role"`   // 1 untuk admin, 0 untuk non-admin
}

func CreateDefaultUser(username string) Users {
    user_dummy := Users {
        UserId          : 0,   
        Username        : username,  
        Password        : "password",  
        Name            : "DEFAULT_NAME",  
        ProductName     : "DEFAULT_PRODUCT_NAME",  
        CompanyName     : "DEFAULT_COMPANY_NAME",  
        CompanyDescription     : "DEFAULT_DESCRIPTION",  
        Visi            : "DEFAULT_VISI",   
        Misi            : "DEFAULT_MISI",
        Jurusan         : "DEFAULT_JURUSAN",
        Angkatan        : 0,       
        LogoPath 		: "DEFAULT_LOGO_PATH",
        CreatedAt       : time.Now().UnixNano(), 
        UpdatedAt       : 0,   
        ShowProfile     : false,    
        Role            : 0,     
    }
    return user_dummy
}

