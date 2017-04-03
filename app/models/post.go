package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"

)

type Post struct {
    PostId          int64   `db:"postid"`
    UserId          int64   `db:"userid"`
    MediaType       string  `db:"media_type"`
    PathFile        string  `db:"path_file"`
    CreatedAt       int64   `db:"post_created_at"`
}