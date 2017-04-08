package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
	"time"

)

type News struct {
	NewsId 			int64 	`db:"newsid"`
	NewsTitle		string 	`db:"news_title"`
	Content			string 	`db:"content"`
    CreatedAt       int64   `db:"news_created_at"`
    UpdatedAt		int64	`db:"news_updated_at"`
}

func CreateDefaultNews(news_title string) News {
    news_dummy := News{
		NewsId 		: 0,
		NewsTitle 	: news_title,
		Content 	: "DEFAULT_CONTENT",
	    CreatedAt       : time.Now().UnixNano(), 
	    UpdatedAt		: 0,
	}
    return news_dummy
}