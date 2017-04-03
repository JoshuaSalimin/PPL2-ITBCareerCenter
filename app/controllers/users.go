package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "time"
    "log"
)

func createUsersAdmin(dbm *gorp.DbMap){
    // set "userid" as primary key and autoincrement
	admin := &models.Users{
		UserId: 1,
		Username: "admin",
		Name: "admin",
		Password: "password",
		Role: 1,
		CreatedAt: time.Now().UnixNano(),
		ShowProfile: false,
	}
    dbm.Insert(admin)
}

func createUsers(dbm *gorp.DbMap, u models.Users){
    err := dbm.Insert(&u)
    checkErr(err, "Insert failed")
}


func selectAllUsers(dbm *gorp.DbMap) []models.Users {
	var u []models.Users

    _, err := dbm.Select(&u, "SELECT * FROM users")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range u {
        log.Printf("    %d: %v\n", x, p)
    }
    return u 	
}

func selectUsersByUserid(dbm *gorp.DbMap, userid int) models.Users {
	var u models.Users
    err := dbm.SelectOne(&u, "SELECT * FROM users WHERE userid=?", userid)
    checkErr(err, "SelectOne failed")
    log.Println("u :", u)
    return u
}

func updateUsers(dbm *gorp.DbMap, u models.Users) {
	count, err := dbm.Update(&u)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func deleteUsersByUserid(dbm *gorp.DbMap, userid int) {
    _, err := dbm.Exec("DELETE FROM users WHERE userid=?", userid)
    checkErr(err, "Delete failed")
}
