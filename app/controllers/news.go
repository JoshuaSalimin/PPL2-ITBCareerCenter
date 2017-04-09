package controllers

import (
	"github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
    "time"
    "strconv"
)

type News struct {
	*revel.Controller
}

func (c News) Form() revel.Result {
	return c.Render()
}

func (c News) List() revel.Result {
    news := SelectAllNews(Dbm);
    return c.Render(news);
}

func (c News) Add() revel.Result {       
    innews := models.News{
        NewsId : 0,
        NewsTitle : c.Request.Form.Get("newstitle"),
        Content : c.Request.Form.Get("newscontent"),
        CreatedAt : time.Now().UnixNano(),
        UpdatedAt : time.Now().UnixNano(), 
    }
    success := InsertNews(Dbm, innews);
    if (success){
        return c.Redirect(News.List);
    }else{
        return c.Redirect(News.Form);        
    }
}

func (c News) Detail() revel.Result {
        var id int;
        c.Params.Bind(&id,"id");
        news :=  SelectNewsByNewsId(Dbm,id);
        return c.Render(news);
}

func (c News) Delete() revel.Result {
        id,_ := strconv.Atoi(c.Request.Form.Get("id"));
        success := DeleteNewsByNewsid(Dbm,id);
        if (success){
            return c.Redirect(News.List);
        }else{
            return c.Redirect(App.Index);        
        }
}

func (c News) EditForm() revel.Result {
        var id int;
        c.Params.Bind(&id,"id");
        news :=  SelectNewsByNewsId(Dbm,id);
        return c.Render(news);
}

func (c News) EditSubmit() revel.Result {
    newsid,_ := strconv.ParseInt(c.Request.Form.Get("newsid"),0,64);
    innews := models.News{
        NewsId : newsid,
        NewsTitle : c.Request.Form.Get("newstitle"),
        Content : c.Request.Form.Get("newscontent"),
        CreatedAt : time.Now().UnixNano(),
        UpdatedAt : time.Now().UnixNano(), 
    }
    success := UpdateNews(Dbm, innews);
    if (success){
        return c.Redirect(News.List);
    }else{
        return c.Redirect(App.Index);        
    }
}


func InsertNews(dbm *gorp.DbMap, p models.News) bool{
    err := dbm.Insert(&p)    
    checkErr(err, "Insert failed")    
    if(err == nil){
        return true;
    }else{
        return false;
    }
}

func SelectAllNews(dbm *gorp.DbMap) []models.News {
	var p []models.News

    _, err := dbm.Select(&p, "SELECT * FROM News")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    return p 	
}

func SelectNewsByNewsId(dbm *gorp.DbMap, newsid int) models.News {
	var p models.News
    err := dbm.SelectOne(&p, "SELECT * FROM News WHERE newsid=?", newsid)
    checkErr(err, "SelectOne failed")
    return p
}

func UpdateNews(dbm *gorp.DbMap, p models.News) bool{
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
    if(err == nil){
        return true;
    }else{
        return false;
    }
}


func DeleteNewsByNewsid(dbm *gorp.DbMap, newsid int) bool{
    _, err := dbm.Exec("DELETE FROM News WHERE newsid=?", newsid)
    checkErr(err, "Delete failed")
    if(err == nil){
        return true;
    }else{
        return false;
    }
}
