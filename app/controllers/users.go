package controllers

import (
//"PPL2-ITBCareerCenter/app/models"
"github.com/revel/revel"
models "PPL2-ITBCareerCenter/app/models"
    // "encoding/json"
"github.com/go-gorp/gorp"
    // "time"
"log"
"time"
"strconv"
"math/rand"
)

type Users struct {
    *revel.Controller
}

func (c Users) Users() revel.Result {
    users := SelectAllUsers(Dbm)
    return c.Render(users)
}

func (c Users) AddView() revel.Result {
    return c.Render(true)
}

func (c Users) Add() revel.Result {
    timecreated := time.Now().UnixNano()
    angkatan,_ := strconv.Atoi(c.Request.Form.Get("angkatan"))
    user := models.Users {
        UserId: 0,
        Username: c.Request.Form.Get("username"),
        Password: c.Request.Form.Get("password"),
        Name: "",
        ProductName: "",
        CompanyName: "",
        CompanyDescription: "",
        Visi: "",
        Misi: "",
        Jurusan: "",
        Angkatan: angkatan,
        LogoPath: "",
        CreatedAt: timecreated,
        UpdatedAt: timecreated,
        ShowProfile: false,
        Role: 0,
    }
    InsertUsers(Dbm, &user)
    c.Flash.Success("User " + c.Request.Form.Get("username") + " added successfully");
    return c.Redirect("/Users")
}

func (c Users) Delete() revel.Result {
    c.Flash.Success("User added successfully");
    id,_ := strconv.Atoi(c.Request.Form.Get("id"))
    DeleteUsersByUserId(Dbm,id)
    //c.Flash.Success("User " + "deleted successfully");
    return c.Redirect("/Users")
}

func (c Users) EditView() revel.Result {
    id,_ := strconv.Atoi(c.Params.Query.Get("id"))
    user := SelectUsersByUserid(Dbm, id)
    return c.Render(user)
}

func (c Users) Edit() revel.Result {
    userid,_ := strconv.Atoi(c.Request.Form.Get("userid"))
    username := c.Request.Form.Get("username")
    password := c.Request.Form.Get("password")
    angkatan,_ := strconv.Atoi(c.Request.Form.Get("angkatan"))

    UpdateUsersByUserid(Dbm, userid, username, password, angkatan)
    
    c.Flash.Success("User edited successfully");
    return c.Redirect("/Users")
}

func generateRandomPassword(digits int) string {
    const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$^&*()_+-=/?,.<>"
    pwd := make([]byte, digits)

    rand.Seed(time.Now().UnixNano())
    for i := range pwd {
        pwd[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(pwd)
}

func InsertUsersAdmin(dbm *gorp.DbMap){
    // set "userid" as primary key and autoincrement
    var admin models.Users
    admin = models.CreateDefaultUser("admin")
    log.Println("u :", admin)
    dbm.Insert(&admin)
}

func InsertUsers(dbm *gorp.DbMap, u *models.Users){
    err := dbm.Insert(u)
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

func SelectUserByUsername(dbm *gorp.DbMap, username string) models.Users {
	var u models.Users
    dbm.SelectOne(&u, "SELECT * FROM users WHERE Username=?", username)
    return u
}

func SelectUserByUsernameAndPassword(dbm *gorp.DbMap, username string, password string) models.Users {
    var u models.Users
    dbm.SelectOne(&u, "SELECT * FROM users WHERE Username=? AND Password=?", username, password)
    return u
}

func UpdateUsers(dbm *gorp.DbMap, u models.Users) {
	count, err := dbm.Update(&u)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}

func UpdateUsersByUserid(dbm *gorp.DbMap, userid int, username string, password string, angkatan int) {
    _, err := dbm.Exec("UPDATE users SET username=?, password=?, angkatan=? WHERE userid=?", username, password, angkatan, userid)
    checkErr(err, "Update failed")
    log.Println("Updated")
}

func DeleteUsersByUserId(dbm *gorp.DbMap, userid int) {
    _, err := dbm.Exec("DELETE FROM Users WHERE userid=?", userid)
    checkErr(err, "Delete failed")
}