package controllers

import (
	"github.com/revel/revel"
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)

type News struct {
	*revel.Controller
}

func (c News) Form() revel.Result {
	return c.Render()
}

func (c News) List() revel.Result {
    return c.Render()
}

func (c News) Add() revel.Result {
    err := c.Request().ParseForm();
    if(err != nil){
        c.Redirect(News.Form);
    };    
    
    innews = News{
        NewsId : 0.
        NewsTitle : c.Request.Form["newsritle"],
        Content : c.Request.Form["newscontent"],
        CreatedAt : time.Now().UnixNano(),
        UpdatedAt : time.Now().UnixNano(), 
    }
    success = InsertNews(Dbm, innews);
    if (success){
        return c.Redirect(News.List);
    }else{
        return c.Redirect(News.Form);        
    }
}

func InsertNews(dbm *gorp.DbMap, p models.News) int{
    err := dbm.Insert(&p)
    if(checkErr(err, "Insert failed")){
        return 0;
    }else{
        return 1;
    }
}

func SelectAllNews(dbm *gorp.DbMap) []models.News {
	var p []models.News

    _, err := dbm.Select(&p, "SELECT * FROM news")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func SelectNewsByNewsId(dbm *gorp.DbMap, newsid int) models.News {
	var p models.News
    err := dbm.SelectOne(&p, "SELECT * FROM news WHERE newsid=?", newsid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func UpdateNews(dbm *gorp.DbMap, p models.News) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func DeleteNewsByNewsid(dbm *gorp.DbMap, newsid int) {
    _, err := dbm.Exec("DELETE FROM news WHERE newsid=?", newsid)
    checkErr(err, "Delete failed")
}
