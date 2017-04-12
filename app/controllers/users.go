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

    _, err := dbm.Select(&u, "SELECT * FROM Users")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range u {
        log.Printf("    %d: %v\n", x, p)
    }
    return u 	
}

func SelectLatestUsersInRange(dbm *gorp.DbMap, start int, count int) []models.Users {
    var u []models.Users

    _, err := dbm.Select(&u, "SELECT * FROM Users ORDER BY users_created_at DESC LIMIT ?, ?", start, count)
    checkErr(err, "Select failed")
    log.Println("User Range rows:")
    for x, p := range u {
        log.Printf("    %d: %v\n", x, p)
    }
    return u    
}

func CountUsers(dbm *gorp.DbMap) int {
    count, err := dbm.SelectInt("SELECT COUNT(*) FROM Users")
    checkErr(err, "Select failed")
    log.Println("User count:", count)
    return int(count)
}

func SelectUsersByUserid(dbm *gorp.DbMap, userid int) models.Users {
	var u models.Users
    err := dbm.SelectOne(&u, "SELECT * FROM Users WHERE userid=?", userid)
    checkErr(err, "SelectOne failed")
    log.Println("u :", u)
    return u
}

func UpdateUsers(dbm *gorp.DbMap, u models.Users) {
	count, err := dbm.Update(&u)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}

func DeleteUsersByUserid(dbm *gorp.DbMap, userid int) {
    _, err := dbm.Exec("DELETE FROM Users WHERE userid=?", userid)
    checkErr(err, "Delete failed")
}

