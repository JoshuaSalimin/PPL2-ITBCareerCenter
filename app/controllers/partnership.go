package controllers

import (
    "github.com/revel/revel"
    //models "PPL2-ITBCareerCenter/app/models"
    //"github.com/go-gorp/gorp"
    // "log"
    // "time"
    // "strconv"
)

type Partnership struct {
    *revel.Controller
}

func (p Partnership) Partnership() revel.Result {
    partnership := true
    return p.Render(partnership)
}

func (p Partnership) EditPartnership() revel.Result {
    partnership := true
    return p.Render(partnership)
}

/*
func InsertAbout(dbm *gorp.DbMap, a models.About){
    err := dbm.Insert(&a)
    checkErr(err, "Insert failed")
}

func SavetoDB(dbm *gorp.DbMap, a models.About){
    count, err := dbm.Update(&a)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}

func (a About) Save() revel.Result {
    aboutid,_ := strconv.ParseInt(a.Request.Form.Get("aboutid"),0,64);
    newAbout := models.About{
        AboutID            : aboutid,
        CkeditorAbout      : a.Request.Form.Get("aboutcontent"),  
        UpdatedAt          : time.Now().UnixNano(),     
    }
    SavetoDB(Dbm, newAbout);
    return a.Redirect(About.About);
}

func SelectAboutByAboutID(dbm *gorp.DbMap, id int) models.About {
    var a models.About
    err := dbm.SelectOne(&a, "SELECT * FROM About WHERE aboutid=?", id)
    checkErr(err, "SelectOne failed")
    log.Println("a :", a)
    return a
}

func CountAbout(dbm *gorp.DbMap) int {
    count, err := dbm.SelectInt("SELECT COUNT(*) FROM About")
    checkErr(err, "Select failed")
    log.Println("About count:", count)
    return int(count)
}

func SelectAllAbout(dbm *gorp.DbMap) []models.About {
    var a []models.About

    _, err := dbm.Select(&a, "SELECT * FROM About")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range a {
        log.Printf("    %d: %v\n", x, p)
    }
    return a   
}*/