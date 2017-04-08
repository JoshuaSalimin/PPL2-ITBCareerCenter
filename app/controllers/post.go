package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)



func createPostAdmin(dbm *gorp.DbMap){
    // set "postid" as primary key and autoincrement
	adminPost := models.CreateDefaultPost("Admin Post")
    dbm.Insert(&adminPost)
}

func createPost(dbm *gorp.DbMap, p models.Post){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func selectAllPost(dbm *gorp.DbMap) []models.Post {
	var p []models.Post

    _, err := dbm.Select(&p, "SELECT * FROM posts")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func selectPostByPostId(dbm *gorp.DbMap, postid int) models.Post {
	var p models.Post
    err := dbm.SelectOne(&p, "SELECT * FROM posts WHERE postid=?", postid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func selectPostByUserId(dbm *gorp.DbMap, userid int) models.Post {
    var p models.Post
    err := dbm.SelectOne(&p, "SELECT * FROM posts WHERE userid=?", userid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func updatePost(dbm *gorp.DbMap, p models.Post) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func deletePostByPostid(dbm *gorp.DbMap, postid int) {
    _, err := dbm.Exec("DELETE FROM posts WHERE postid=?", postid)
    checkErr(err, "Delete failed")
}
