package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)

func InsertUserContact(dbm *gorp.DbMap, p models.UserContact){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func SelectAllUserContact(dbm *gorp.DbMap) []models.UserContact {
	var p []models.UserContact

    _, err := dbm.Select(&p, "SELECT * FROM usercontact")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func SelectUserContactByUserContactId(dbm *gorp.DbMap, contactid int) models.UserContact {
	var p models.UserContact
    err := dbm.SelectOne(&p, "SELECT * FROM usercontact WHERE contactid=?", contactid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func UpdateUserContact(dbm *gorp.DbMap, p models.UserContact) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func DeleteUserContactByUserContactid(dbm *gorp.DbMap, contactid int) {
    _, err := dbm.Exec("DELETE FROM usercontact WHERE contactid=?", contactid)
    checkErr(err, "Delete failed")
}

func SelectAllUserContactByUserId(dbm *gorp.DbMap, id int) []models.UserContact {
    var p []models.UserContact

    _, err := dbm.Select(&p, "SELECT * FROM usercontact WHERE userid=?", id)
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p    
}