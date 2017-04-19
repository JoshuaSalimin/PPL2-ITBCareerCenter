package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type UserContact struct {
    ContactID               int64   `db:"contactid"`
    UserId                  int64   `db:"userid"`
    ContactType             string  `db:"contact_type"`
    Contact                 string  `db:"contact"`
    CreatedAt               int64   `db:"contact_created_at"`
    UpdatedAt               int64   `db:"contact_updated_at"`     
}

func CreateDefaultUserContact() UserContact {
    tmp := UserContact{
        ContactID       : 0,
        UserId          : 0,
        ContactType     : "DEFAULT_CONTACT_TYPE",
        Contact         : "DEFAULT_CONTACT",
        CreatedAt       : time.Now().UnixNano(), 
        UpdatedAt       : time.Now().UnixNano(),
    }
    return tmp
}

