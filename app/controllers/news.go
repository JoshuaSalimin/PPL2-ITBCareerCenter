package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)

func createNews(dbm *gorp.DbMap, p models.News){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func selectAllNews(dbm *gorp.DbMap) []models.News {
	var p []models.News

    _, err := dbm.Select(&p, "SELECT * FROM news")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func selectNewsByNewsId(dbm *gorp.DbMap, newsid int) models.News {
	var p models.News
    err := dbm.SelectOne(&p, "SELECT * FROM news WHERE newsid=?", newsid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func updateNews(dbm *gorp.DbMap, p models.News) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func deleteNewsByNewsid(dbm *gorp.DbMap, newsid int) {
    _, err := dbm.Exec("DELETE FROM news WHERE newsid=?", newsid)
    checkErr(err, "Delete failed")
}
