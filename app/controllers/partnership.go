package controllers

import (
    "github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    "github.com/go-gorp/gorp"
     "log"
     //"time"
     //"strconv"
)

type Partnership struct {
    *revel.Controller
}

func (p Partnership) Partnership() revel.Result {
    partnership := true
    allPartnership := SelectAllPartnership(Dbm)
    return p.Render(partnership, allPartnership)
}

func (p Partnership) EditPartnership() revel.Result {
    partnership := true
    allPartnership := SelectAllPartnership(Dbm)
    return p.Render(partnership, allPartnership)
}

func InsertPartnership(dbm *gorp.DbMap, p models.Partnership){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func PartnershiptoDB(dbm *gorp.DbMap, p models.Partnership){
    count, err := dbm.Update(&p)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}


func (p Partnership) SavePartnership() revel.Result {
    return p.Redirect(Partnership.Partnership);
}

/*
func (a About) SavePartnership() revel.Result {
    aboutid,_ := strconv.ParseInt(a.Request.Form.Get("aboutid"),0,64);
    newAbout := models.About{
        AboutID            : aboutid,
        CkeditorAbout      : a.Request.Form.Get("aboutcontent"),  
        UpdatedAt          : time.Now().UnixNano(),     
    }
    PartnershiptoDB(Dbm, newAbout);
    return a.Redirect(About.About);
}*/

func SelectPartnershipByPartnershipID(dbm *gorp.DbMap, id int) models.Partnership {
    var p models.Partnership
    err := dbm.SelectOne(&p, "SELECT * FROM Partnership WHERE partnershipid=?", id)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func CountPartnership(dbm *gorp.DbMap) int {
    count, err := dbm.SelectInt("SELECT COUNT(*) FROM Partnership")
    checkErr(err, "Select failed")
    log.Println("Partnership count:", count)
    return int(count)
}

func SelectAllPartnership(dbm *gorp.DbMap) []models.Partnership {
    var p []models.Partnership
    _, err := dbm.Select(&p, "SELECT * FROM Partnership")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, a := range p {
        log.Printf("    %d: %v\n", x, a)
    }
    return p   
}