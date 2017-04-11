package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type UserSocialMedia struct {
    UserSocialMediaId      int   `db:"socialmediaid"`
    UserId                  int   `db:"userid"`
    SocialMediaName         string  `db:"social_media_name"`
    SocialMediaURL          string  `db:"social_media_url"`
    CreatedAt               int64   `db:"social_media_created_at"`
    UpdatedAt               int64   `db:"social_media_updated_at"`     
}

func CreateDefaultUserSocialMedia() UserSocialMedia {
    tmp := UserSocialMedia{
        UserSocialMediaId: 0, 
        UserId: 0, 
        SocialMediaName: "DEFUALT_SOCIAL_MEDIA_NAME", 
        SocialMediaURL: "DEFAULT_SOCIAL_MEDIA_URL", 
        CreatedAt       : time.Now().UnixNano(), 
        UpdatedAt       : time.Now().UnixNano(),
    }
    return tmp
}

