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


func (c Users) Users(page int) revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    numUserPerPage := 20
    if (page == 0) {
        page = 1
    }
    startUserLimit := (page-1)*numUserPerPage
    endUserLimit := page*numUserPerPage

    userCount := CountUsers(Dbm)

    startUserLimit = max(startUserLimit, 0)

    if (startUserLimit >= userCount) {
        return c.NotFound("Invalid Page: ", page);
    }

    endUserLimit = min(userCount, endUserLimit)

    users := SelectLatestUsersInRange(Dbm, startUserLimit, endUserLimit - startUserLimit)
    currentPageNum := page
    return c.Render(page, users, userCount, numUserPerPage, currentPageNum)
}

func (c Users) RedirectToList() revel.Result {
    return c.Redirect("/Users/List/1")
}

func (c Users) AddView() revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    return c.Render(true)
}

func (c Users) Add() revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

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
    user.Password = EncryptSHA256(user.Password)
    InsertUsers(Dbm, &user)
    c.Flash.Success("User " + c.Request.Form.Get("username") + " added successfully");
    return c.Redirect("/Users")
}

func (c Users) Delete() revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    c.Flash.Success("User added successfully");
    id,_ := strconv.Atoi(c.Request.Form.Get("id"))
    DeleteUsersByUserId(Dbm,id)
    //c.Flash.Success("User " + "deleted successfully");
    return c.Redirect("/Users")
}

func (c Users) EditView() revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    id,_ := strconv.Atoi(c.Params.Query.Get("id"))
    user := SelectUsersByUserid(Dbm, id)
    return c.Render(user)
}

func (c Users) Edit() revel.Result {
    //Check Auth
    isAuthorized := c.IsAuthorized()
    if (!isAuthorized) {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login");
    }

    userid,_ := strconv.Atoi(c.Request.Form.Get("userid"))
    username := c.Request.Form.Get("username")
    password := c.Request.Form.Get("password")
    password = EncryptSHA256(password)
    angkatan,_ := strconv.Atoi(c.Request.Form.Get("angkatan"))

    UpdateUsersByUserid(Dbm, userid, username, password, angkatan)
    
    c.Flash.Success("User edited successfully");
    return c.Redirect("/Users")
}

func (c Users) IsAuthorized() bool {
    //Check Auth
    isAdmin := false
    if (c.Session["cUserRole"] == "1") {
        isAdmin = true
    }
    return isAdmin
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
    admin.Password = EncryptSHA256(admin.Password)
    admin.Role = 1;
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