package models 

import (
    "time"
)

type Partnership struct {
	PartnershipID           int64   `db:"partnershipid"` 
    PartnershipName		    string 	`db:"partnership_name"`
    PartnershipLink         string  `db:"partnership_link"`
    ImgPath                 string  `db:"img_path"`
    CreatedAt		        int64	`db:"partnership_created_at"`
    UpdatedAt		        int64	`db:"partnership_updated_at"`		
}

func CreateDefaultPartnership() Partnership {
    partnership_dummy := Partnership {
        PartnershipID      : 0,
        PartnershipName    : "DEFAULT_PARTNERSHIP",
        PartnershipLink    : "www.google.com",
        ImgPath            : "DEFAULT_PATH",
        CreatedAt          : time.Now().UnixNano(), 
        UpdatedAt          : time.Now().UnixNano(),     
    }
    return partnership_dummy
}