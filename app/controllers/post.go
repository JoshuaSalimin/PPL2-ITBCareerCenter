package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)

func InsertPostAdmin(dbm *gorp.DbMap){
    // set "postid" as primary key and autoincrement
	adminPost := models.CreateDefaultPost("Admin Post")
    dbm.Insert(&adminPost)
}

func InsertPost(dbm *gorp.DbMap, p models.Post){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func SelectAllPost(dbm *gorp.DbMap) []models.Post {
	var p []models.Post

    _, err := dbm.Select(&p, "SELECT * FROM post")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func SelectPostByPostId(dbm *gorp.DbMap, postid int) models.Post {
	var p models.Post
    err := dbm.SelectOne(&p, "SELECT * FROM post WHERE postid=?", postid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func SelectPostByUserId(dbm *gorp.DbMap, userid int) models.Post {
    var p models.Post
    err := dbm.SelectOne(&p, "SELECT * FROM post WHERE userid=?", userid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func UpdatePost(dbm *gorp.DbMap, p models.Post) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}

func DeletePostByPostid(dbm *gorp.DbMap, postid int) {
    _, err := dbm.Exec("DELETE FROM post WHERE postid=?", postid)
    checkErr(err, "Delete failed")
}

func SelectUserImage(dbm *gorp.DbMap, id int) []models.Post {
	var p []models.Post
	image := "Image"
    _, err := dbm.Select(&p, "SELECT * FROM post WHERE media_type=? AND userid=?", image, id)
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func SelectVideoByUserId(dbm *gorp.DbMap, userid int) models.Post {
    var p models.Post
    video := "Video"
    err := dbm.SelectOne(&p, "SELECT * FROM post WHERE userid=? AND media_type=?", userid, video)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}
