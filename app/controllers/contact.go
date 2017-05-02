package controllers

import (
    "github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    "github.com/go-gorp/gorp"
    "log"
    "time"
    "strconv"
)

type Contact struct {
    *revel.Controller
}

func (c Contact) Contact() revel.Result {
    //Check Auth
    isAdmin := false
    if (c.Session["cUserRole"] == "1") {
        isAdmin = true
    }
    
    contact := true
    allcontact := SelectAllContact(Dbm)
    contentcontact := allcontact[0]
    return c.Render(contact, contentcontact, isAdmin)
}

func (c Contact) EditContact() revel.Result {
    //Check Auth
    isAdmin := false
    if (c.Session["cUserRole"] == "1") {
        isAdmin = true
    }
    if (!isAdmin) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    about := true
    allcontactcontent := SelectAllContact(Dbm)
    contactcontent := allcontactcontent[0]
    return c.Render(about, contactcontent)
}

func InsertContact(dbm *gorp.DbMap, c models.Contact){
    err := dbm.Insert(&c)
    checkErr(err, "Insert failed")
}

func ContactSavetoDB(dbm *gorp.DbMap, c models.Contact){
    count, err := dbm.Update(&c)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}

func (c Contact) SaveContact() revel.Result {
    contactid,_ := strconv.ParseInt(c.Request.Form.Get("contactid"),0,64);
    newContact := models.Contact{
        ContactID            : contactid,
        Content              : c.Request.Form.Get("contactcontent"),  
        UpdatedAt            : time.Now().UnixNano(),     
    }
    ContactSavetoDB(Dbm, newContact);
    return c.Redirect(Contact.Contact);
}

func SelectContactByContactID(dbm *gorp.DbMap, id int) models.Contact {
    var c models.Contact
    err := dbm.SelectOne(&c, "SELECT * FROM Contact WHERE contactid=?", id)
    checkErr(err, "SelectOne failed")
    log.Println("c :", c)
    return c
}

func CountContact(dbm *gorp.DbMap) int {
    count, err := dbm.SelectInt("SELECT COUNT(*) FROM Contact")
    checkErr(err, "Select failed")
    log.Println("Contact count:", count)
    return int(count)
}

func SelectAllContact(dbm *gorp.DbMap) []models.Contact {
    var c []models.Contact
    _, err := dbm.Select(&c, "SELECT * FROM Contact")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range c {
        log.Printf("    %d: %v\n", x, p)
    }
    return c   
}