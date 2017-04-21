package models 

import (

	"time"
)

type Post struct {
    PostId          int64   `db:"postid"`
    PostTitle		string	`db:"post_title"`
    UserId          int64   `db:"userid"`
    MediaType       string  `db:"media_type"`
    PathFile        string  `db:"path_file"`
    CreatedAt       int64   `db:"post_created_at"`
    UpdatedAt		int64	`db:"post_updated_at"`
}

func CreateDefaultPost(post_title string) Post {
    post_dummy := Post{
	    PostId          : 0,
	    PostTitle		: post_title,
	    UserId          : 0,
	    MediaType       : "DEFAULT_MEDIA_TYPE",
	    PathFile        : "DEFAULT_PATH_FILE",
	    CreatedAt       : time.Now().UnixNano(), 
	    UpdatedAt		: 0,
    }
    return post_dummy
}