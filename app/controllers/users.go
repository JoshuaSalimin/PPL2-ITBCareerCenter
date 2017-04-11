package controllers

import (
"PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
"github.com/go-gorp/gorp"
    // "time"
"log"
)



func InsertUsersAdmin(dbm *gorp.DbMap){
    // set "userid" as primary key and autoincrement
    var admin models.Users
    admin = models.CreateDefaultUser("admin");
    log.Println("u :", admin)
    dbm.Insert(&admin)
}

func InsertUsers(dbm *gorp.DbMap, u models.Users){
    err := dbm.Insert(&u)
    checkErr(err, "Insert failed")
}


func SelectAllUsers(dbm *gorp.DbMap) []models.Users {
	var u []models.Users

    _, err := dbm.Select(&u, "SELECT * FROM users")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range u {
        log.Printf("    %d: %v\n", x, p)
    }
    return u 	
}

func SelectUserByUserId(dbm *gorp.DbMap, userid int64) models.Users {
    var u models.Users
    err := dbm.SelectOne(&u, "SELECT * FROM users WHERE UserId=?", userid)
    checkErr(err, "SelectOne failed")
    log.Println("u :", u)
    return u
}

func SelectUserByUsername(dbm *gorp.DbMap, username string) models.Users {
	var u models.Users
    err := dbm.SelectOne(&u, "SELECT * FROM users WHERE Username=?", username)
    //log.Println("Query: " + "SELECT * FROM users WHERE Username='" + username +"'");
    if (err != nil) {
        return u
    } else {
        return u
    }
    //checkErr(err, "SelectOne failed")
    //log.Println("u :", u)
}

func SelectUserByUsernameAndPassword(dbm *gorp.DbMap, username string, password string) models.Users {
    var u models.Users
    err := dbm.SelectOne(&u, "SELECT * FROM users WHERE Username=? AND Password=?", username, password)
    //log.Println("Query: " + "SELECT * FROM users WHERE Username='" + username +"'");
    if (err != nil) {
        return u
    } else {
        return u
    }
}

func UpdateUsers(dbm *gorp.DbMap, u models.Users) {
	count, err := dbm.Update(&u)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func DeleteUsersByUserid(dbm *gorp.DbMap, userid int) {
    _, err := dbm.Exec("DELETE FROM users WHERE userid=?", userid)
    checkErr(err, "Delete failed")
}

