package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type About struct {
	AboutID                 int64   `db:"aboutid"` 
    CkeditorAbout		    string 	`db:"ckeditor_about, size:1000000"`
    CreatedAt		        int64	`db:"about_created_at"`
    UpdatedAt		        int64	`db:"about_updated_at"`		
}

func CreateDefaultAbout() About {
    about_dummy := About {
        AboutID            : 0,
        CkeditorAbout      : "Default_ABOUT",  
        CreatedAt          : time.Now().UnixNano(), 
        UpdatedAt          : time.Now().UnixNano(),     
    }
    return about_dummy
}