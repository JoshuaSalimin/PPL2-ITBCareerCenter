package controllers

import (
	"github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
    "time"
    "strconv"
)

type Event struct {
	*revel.Controller
}


func (c Event) EventDetail(id int) revel.Result {
    ev := SelectEventByEventId(Dbm, id)
    log.Println(ev)
    EventTitle := ev.EventTitle
    EventBannerPath := ev.BannerPath
    EventStart := time.Unix(ev.EventStart, 0)
    EventEnd := time.Unix(ev.EventEnd, 0)
    EventDescription := ev.EventDescription
    EventCreatedAt := time.Unix(0, ev.CreatedAt)
    EventUpdatedAt := time.Unix(0, ev.UpdatedAt)
    isAuthorizedAsAdmin := true
    return c.Render(true, id, EventTitle, EventBannerPath, EventStart, EventEnd, 
        EventDescription, EventCreatedAt, EventUpdatedAt, isAuthorizedAsAdmin)
}

func (c Event) DeleteEvent(id int) revel.Result {
    DeleteEventByEventId(Dbm, id)
    return c.Redirect("/Event")
}

func (c Event) EditEvent(id int) revel.Result {
    ev := SelectEventByEventId(Dbm, id)
    log.Println(ev)
    EventTitle := ev.EventTitle
    EventBannerPath := ev.BannerPath
    EventStart_ := time.Unix(ev.EventStart, 0)
    EventStart := TimeToStringHTML(EventStart_)
    EventEnd_ := time.Unix(ev.EventEnd, 0)
    EventEnd := TimeToStringHTML(EventEnd_)
    EventDescription := ev.EventDescription
    EventCreatedAt := time.Unix(0, ev.CreatedAt)
    EventUpdatedAt := time.Unix(0, ev.UpdatedAt)
    isAuthorizedAsAdmin := true
    return c.Render(true, id, EventTitle, EventBannerPath, EventStart, EventEnd, 
        EventDescription, EventCreatedAt, EventUpdatedAt, isAuthorizedAsAdmin)
}

func TimeToStringHTML(t time.Time) string {
    year, month_, day := time.Time.Date(t)
    var month int = int(month_) 
    hour, minute, second := time.Time.Clock(t)
    // year = 1990
    // month = 12
    // day = 12
    // hour = 12
    // minute = 12
    // second = 12

    yearString := strconv.Itoa(year)

    var monthString string
    var dayString string
    var hourString string
    var minuteString string
    var secondString string

    if (month < 10) {
        monthString = "0" + strconv.Itoa(month)
    } else {
        monthString = strconv.Itoa(month)    
    }

    if (day < 10) {
        dayString = "0" + strconv.Itoa(day)
    } else {
        dayString = strconv.Itoa(day)
    }

    if (hour < 10) {
        hourString = "0" + strconv.Itoa(hour)
    } else {
        hourString = strconv.Itoa(hour)
    }
   
    if (minute < 10) {
        minuteString = "0" + strconv.Itoa(minute)
    } else {
        minuteString = strconv.Itoa(minute)
    }

    if (second < 10) {
        secondString = "0" + strconv.Itoa(minute)
    } else {    
        secondString = strconv.Itoa(second)
    }



    stringTime := yearString + "-" + monthString + "-" + dayString + "T" + hourString + ":" + minuteString + ":" + secondString
    log.Println(stringTime)
    return stringTime
}


func InsertEvent(dbm *gorp.DbMap, p models.Event) bool {
    err := dbm.Insert(&p)    
    checkErr(err, "Insert failed")    
    if(err == nil){
        return true;
    } else{
        return false;
    }
}

func SelectAllEvent(dbm *gorp.DbMap) []models.Event {
    var p []models.Event

    _, err := dbm.Select(&p, "SELECT * FROM event")
    checkErr(err, "Select failed")
    return p     
}

func SelectEventByEventId(dbm *gorp.DbMap, eventid int) models.Event {
    var p models.Event
    err := dbm.SelectOne(&p, "SELECT * FROM event WHERE eventid=?", eventid)
    checkErr(err, "SelectOne failed")
    return p
}

func DeleteEventByEventId(dbm *gorp.DbMap, eventid int) bool {
    _, err := dbm.Exec("DELETE FROM event WHERE eventid=?", eventid)
    checkErr(err, "Delete failed")
    if(err == nil){
        return true;
    }else{
        return false;
    }
}


// func (c News) Detail() revel.Result {
//         var id int;
//         c.Params.Bind(&id,"id");
//         news :=  SelectNewsByNewsId(Dbm,id);
//         return c.Render(news);
// }

// func (c News) Delete() revel.Result {
//         id,_ := strconv.Atoi(c.Request.Form.Get("id"));
//         success := DeleteNewsByNewsid(Dbm,id);
//         if (success){
//             return c.Redirect(News.List);
//         }else{
//             return c.Redirect(App.Index);        
//         }
// }

// func (c News) EditForm() revel.Result {
//         var id int;
//         c.Params.Bind(&id,"id");
//         news :=  SelectNewsByNewsId(Dbm,id);
//         return c.Render(news);
// }

// func (c News) EditSubmit() revel.Result {
//     newsid,_ := strconv.ParseInt(c.Request.Form.Get("newsid"),0,64);
//     newscreatedat,_ := strconv.ParseInt(c.Request.Form.Get("newscreated"),0,64);
//     innews := models.News{
//         NewsId : newsid,
//         NewsTitle : c.Request.Form.Get("newstitle"),
//         Content : c.Request.Form.Get("newscontent"),
//         CreatedAt : newscreatedat,
//         UpdatedAt : time.Now().UnixNano(), 
//     }
//     success := UpdateNews(Dbm, innews);
//     if (success){
//         return c.Redirect(News.List);
//     }else{
//         return c.Redirect(App.Index);        
//     }
// }


// func InsertNews(dbm *gorp.DbMap, p models.News) bool{
//     err := dbm.Insert(&p)    
//     checkErr(err, "Insert failed")    
//     if(err == nil){
//         return true;
//     }else{
//         return false;
//     }
// }

// func SelectAllNews(dbm *gorp.DbMap) []models.News {
// 	var p []models.News

//     _, err := dbm.Select(&p, "SELECT * FROM News")
//     checkErr(err, "Select failed")
//     log.Println("All rows:")
//     return p 	
// }

// func SelectNewsByNewsId(dbm *gorp.DbMap, newsid int) models.News {
// 	var p models.News
//     err := dbm.SelectOne(&p, "SELECT * FROM News WHERE newsid=?", newsid)
//     checkErr(err, "SelectOne failed")
//     return p
// }

// func UpdateNews(dbm *gorp.DbMap, p models.News) bool{
// 	count, err := dbm.Update(&p)
// 	checkErr(err, "Update failed")	
//     log.Println("Rows updated:", count)
//     if(err == nil){
//         return true;
//     }else{
//         return false;
//     }
// }


