package controllers

import (
	"github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    "os"
    "github.com/go-gorp/gorp"
    "log"
    "time"
    "strconv"
    "html"
    "path/filepath"
    "math/rand" 
    "encoding/hex"    
    "fmt"
)

type Event struct {
	*revel.Controller
}


func (e Event) Event() revel.Result {
    events := true
    list := SelectAllEvent(Dbm)
    var isAuthorizedAsAdmin bool
    if (e.Session["cUserId"] == "") {
        isAuthorizedAsAdmin = false
    } else {
        if (e.Session["cUserRole"] == "1") {
            isAuthorizedAsAdmin = true
        } else {
            isAuthorizedAsAdmin = false
        }
    }
    return e.Render(events, list, isAuthorizedAsAdmin)
}

func (e Event) AddEvent() revel.Result {
    isAuthorizedAsAdmin := true
    return e.Render(isAuthorizedAsAdmin)
}


func (e Event) AddEventToDB(EventTitle string, 
                            EventStart_RFC3339 string, 
                            EventEnd_RFC3339 string, 
                            EventLocation string, 
                            EventBanner []byte,
                            EventDescription string ) revel.Result {

    layout := "2006-01-02T15:04"      

    EventStart, _ := time.Parse(layout, EventStart_RFC3339)
    EventEnd, _ := time.Parse(layout, EventEnd_RFC3339)

    trunc := -7 * time.Hour

    EventStart = EventStart.Add(trunc)
    EventEnd = EventEnd.Add(trunc)

    ev := models.CreateDefaultEvent3(EventTitle)

    if (len(EventBanner) != 0) {
        filename := e.Params.Files["EventBanner"][0].Filename
        relativePath := e.UploadBanner(EventBanner, filename)
        ev.BannerPath = relativePath
    } else {
        log.Println("STILL FAILED. Files : " + strconv.Itoa(len(e.Params.Files["eventbanner"])))
    }

    ev.EventStart = EventStart.Unix()
    ev.EventEnd = EventEnd.Unix()
    ev.EventLocation = EventLocation
    ev.EventDescription = EventDescription

    InsertEvent(Dbm, ev)


    return e.Redirect("/Events")
}

func (e Event) EventDetail(id int) revel.Result {
    ev := SelectEventByEventId(Dbm, id)
    EventTitle := ev.EventTitle
    EventBannerPath := ev.BannerPath
    _EventStart := time.Unix(ev.EventStart, 0)
    _EventEnd := time.Unix(ev.EventEnd, 0)
    EventStart := TimeToString(_EventStart)
    EventEnd := TimeToString(_EventEnd)
    EventLocation := ev.EventLocation
    EventDescription := ev.EventDescription
    EventDescription = html.UnescapeString(EventDescription)
    log.Println(EventDescription)
    EventCreatedAt := time.Unix(0, ev.CreatedAt)
    EventUpdatedAt := time.Unix(0, ev.UpdatedAt)
    var isAuthorizedAsAdmin bool
    if (e.Session["cUserId"] == "") {
        isAuthorizedAsAdmin = false
    } else {
        if (e.Session["cUserRole"] == "1") {
            isAuthorizedAsAdmin = true
        } else {
            isAuthorizedAsAdmin = false
        }
    }
    events := true
    return e.Render(events, id, EventTitle, EventBannerPath, EventStart, EventEnd, 
        EventLocation, EventDescription, EventCreatedAt, EventUpdatedAt, isAuthorizedAsAdmin)
}

func (e Event) DeleteEvent(id int) revel.Result {
    DeleteEventByEventId(Dbm, id)
    return e.Redirect("/Events")
}

func (e Event) EditEvent(id int) revel.Result {
    ev := SelectEventByEventId(Dbm, id)
    EventTitle := ev.EventTitle
    EventBannerPath := ev.BannerPath
    EventStart_ := time.Unix(ev.EventStart, 0)
    EventStart := TimeToStringHTML(EventStart_)
    EventEnd_ := time.Unix(ev.EventEnd, 0)
    EventEnd := TimeToStringHTML(EventEnd_)
    EventLocation := ev.EventLocation
    EventDescription := ev.EventDescription
    EventCreatedAt := time.Unix(0, ev.CreatedAt)
    EventUpdatedAt := time.Unix(0, ev.UpdatedAt)
    var isAuthorizedAsAdmin bool
    if (e.Session["cUserId"] == "") {
        isAuthorizedAsAdmin = false
    } else {
        if (e.Session["cUserRole"] == "1") {
            isAuthorizedAsAdmin = true
        } else {
            isAuthorizedAsAdmin = false
        }
    }
    events := true
    return e.Render(events, id, EventTitle, EventBannerPath, EventStart, EventEnd, 
        EventLocation, EventDescription, EventCreatedAt, EventUpdatedAt, isAuthorizedAsAdmin)
}

func (e Event) UpdateEvent(EventBanner []byte,) revel.Result{
    layout := "2006-01-02T15:04"       
    id, _ := strconv.Atoi(e.Request.Form.Get("id"))
    ev := SelectEventByEventId(Dbm, id)

    EventStart_RFC3339 := e.Request.Form.Get("EventStart")
    EventEnd_RFC3339 := e.Request.Form.Get("EventEnd")
    EventStart, _ := time.Parse(layout, EventStart_RFC3339)
    EventEnd, _ := time.Parse(layout, EventEnd_RFC3339)

    trunc := -7 * time.Hour

    EventStart = EventStart.Add(trunc)
    EventEnd = EventEnd.Add(trunc)


    EventLocation := e.Request.Form.Get("EventLocation")
    EventTitle := e.Request.Form.Get("EventTitle")

    if (len(EventBanner) != 0) {
        filename := e.Params.Files["EventBanner"][0].Filename
        relativePath := e.UploadBanner(EventBanner, filename)
        ev.BannerPath = relativePath
    } else {
        log.Println("STILL FAILED. Files : " + strconv.Itoa(len(e.Params.Files["eventbanner"])))
    }


    EventDescription := e.Request.Form.Get("EventDescription")

    ev.EventTitle = EventTitle
    ev.EventStart = EventStart.Unix()
    ev.EventEnd = EventEnd.Unix()
    ev.EventLocation = EventLocation
    ev.EventDescription = EventDescription
    ev.UpdatedAt = time.Now().UnixNano()

    UpdateEventDB(Dbm, ev)

    return e.Redirect("/Events")
}


func (e Event) UploadBanner(image []byte, filename string) string {
    e.Validation.MaxSize(image, 2*MB).
        Message("File cannot be larger than 2MB")

    fileExt := filepath.Ext(filename)
    randFilename := randBannerString() + fileExt
    relativePath := fmt.Sprintf("/public/images/banner/%s", randFilename)
    dstPath := fmt.Sprintf("%s/public/images/banner/", revel.BasePath)
    if _, err := os.Stat(dstPath); os.IsNotExist(err) {
        os.Mkdir(dstPath, 0777)
    }
    dstPath = dstPath + "/" + randFilename
    dstFile, _ := os.Create(dstPath)
    defer dstFile.Close()
    defer os.Chmod(dstPath, (os.FileMode)(0644))

    dstFile.Write(image)
    return relativePath
}


func TimeToString(t time.Time) string {
    year, month_, day := time.Time.Date(t)
    var month int = int(month_) 
    hour, minute, _ := time.Time.Clock(t)

    yearString := strconv.Itoa(year)

    var monthString string
    var dayString string
    var hourString string
    var minuteString string

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

    s := dayString + "-" + monthString + "-" + yearString + " " + hourString + ":" + minuteString
    return s
}

func TimeToStringHTML(t time.Time) string {
    year, month_, day := time.Time.Date(t)
    var month int = int(month_) 
    hour, minute, _ := time.Time.Clock(t)

    yearString := strconv.Itoa(year)

    var monthString string
    var dayString string
    var hourString string
    var minuteString string
    // var secondString string

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

    // if (second < 10) {
    //     secondString = "0" + strconv.Itoa(minute)
    // } else {    
    //     secondString = strconv.Itoa(second)
    // }

    stringTime := yearString + "-" + monthString + "-" + dayString + "T" + hourString + ":" + minuteString
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

    _, err := dbm.Select(&p, "SELECT * FROM event ORDER BY news_created_at DESC")
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

func UpdateEventDB(dbm *gorp.DbMap, u models.Event) {
    count, err := dbm.Update(&u)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}

func randBannerString() string {
    randBytes := make([]byte, 16)
    rand.Read(randBytes)
    return hex.EncodeToString(randBytes)
}