package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
	"time"

)

type Event struct {
	EventId 		int64 	`db:"eventid"`
	EventTitle		string 	`db:"event_title"`
	BannerPath		string 	`db:"banner_path"`
	EventStart		int64	`db:"event_start"`
	EventEnd		int64	`db:"event_end"`
	EventLocation	string  `db:"event_location"`
	EventDescription	string 	`db:"event_description"`
    CreatedAt       int64   `db:"news_created_at"`
    UpdatedAt		int64	`db:"news_updated_at"`
}

func CreateDefaultEvent(event_title string) Event {
	event_dummy := Event {
	EventId 	: 0,
	EventTitle	: "DEFAULT_TITLE",
	BannerPath	: "/public/images/event1.jpg",
	EventStart	: time.Now().Unix(),
	EventEnd	: time.Now().Unix(),
	EventLocation : "-",
	EventDescription	: "DEFAULT_DESCRIPTION",
    CreatedAt       : time.Now().UnixNano(), 
    UpdatedAt		: 0,
	}
    return event_dummy
}

func CreateDefaultEvent2(event_title string) Event {
	event_dummy := Event {
	EventId 	: 0,
	EventTitle	: "DEFAULT_TITLE",
	BannerPath	: "/public/images/event2.jpg",
	EventStart	: time.Now().Unix(),
	EventEnd	: time.Now().Unix(),
	EventLocation : "-"	,
	EventDescription	: "DEFAULT_DESCRIPTION",
    CreatedAt       : time.Now().UnixNano(), 
    UpdatedAt		: 0,
	}
    return event_dummy
}

func CreateDefaultEvent3(event_title string) Event {
	event_dummy := Event {
	EventId 	: 0,
	EventTitle	: "DEFAULT_TITLE",
	BannerPath	: "/public/images/event3.jpg",
	EventStart	: time.Now().Unix(),
	EventEnd	: time.Now().Unix(),
	EventLocation : "-"	,
	EventDescription	: "DEFAULT_DESCRIPTION",
    CreatedAt       : time.Now().UnixNano(), 
    UpdatedAt		: 0,
	}
    return event_dummy
}