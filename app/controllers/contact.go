package controllers

import (
    "github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
    "time"
    //"strconv"
)

type Contact struct {
    *revel.Controller
}

func (c Contact) Contact() revel.Result {
    contact := true
    content := SelectContactByContactid(Dbm, 1)
    return c.Render(contact, content)
}

func (c Contact) EditContact() revel.Result {
    contact := true
    content := SelectContactByContactid(Dbm, 1)
    return c.Render(contact, content)
}

func InsertContact(dbm *gorp.DbMap, c models.Contact){
    err := dbm.Insert(&c)
    checkErr(err, "Insert failed")
}

func SelectContactByContactid(dbm *gorp.DbMap, id int) models.Contact {
    var c models.Contact
    err := dbm.SelectOne(&c, "SELECT * FROM Contact WHERE contactid=?", id)
    checkErr(err, "SelectOne failed")
    log.Println("c :", c)
    return c
}

func UpdateContact(dbm *gorp.DbMap, c models.Contact) {
    count, err := dbm.Update(&c)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}

func (c Contact) Update(content models.Contact) revel.Result {
    content.ContactID = 1
    content.UpdatedAt = time.Now().UnixNano()
    UpdateContact(Dbm, content)
    return c.Redirect("/Contact")
}
