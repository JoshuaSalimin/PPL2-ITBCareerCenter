package models 

import (
    "time"
)

type Contact struct {
	ContactID               int64   `db:"contactid"` 
    Content		            string 	`db:"about_content, size:1000000"`
    CreatedAt		        int64	`db:"contact_created_at"`
    UpdatedAt		        int64	`db:"contact_updated_at"`		
}

func CreateDefaultContact() Contact {
    contact_dummy := Contact {
        ContactID          : 0,
        Content            : "DEFAULT_CONTENT",  
        CreatedAt          : time.Now().UnixNano(), 
        UpdatedAt          : time.Now().UnixNano(),     
    }
    return contact_dummy
}