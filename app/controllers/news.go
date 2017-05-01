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

func (c News) Index() revel.Result {
    news := true
    list := SelectAllNews(Dbm);
    latest := SelectLatestNews(Dbm);
    return c.Render(news,list,latest);
}

func (c News) Form() revel.Result {
    isAdmin := false
    if (c.Session["cUserRole"] == "1") {
        isAdmin = true
    }
    if (isAdmin) {
        return c.Render()
    } else {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login")
    }
}

func (c News) List() revel.Result {
    news := SelectAllNews(Dbm);
    return c.Render(news);
}

func (c News) Add() revel.Result {
    if (c.Session["cUserRole"] == "1") {
        timecreated := time.Now().UnixNano();
        innews := models.News{
            NewsId : 0,
            NewsTitle : c.Request.Form.Get("newstitle"),
            Content : c.Request.Form.Get("newscontent"),
            CreatedAt : timecreated,
            UpdatedAt : timecreated, 
        }
        success := InsertNews(Dbm, innews);
        if (success){
            return c.Redirect(News.List);
        }else{
            return c.Redirect(News.Form);        
        }    
    } else {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login")
    }   
    
}

func (c News) Detail() revel.Result {
        var id int;
        c.Params.Bind(&id,"id");
        news :=  SelectNewsByNewsId(Dbm,id);
        return c.Render(news);
}

func (c News) Delete() revel.Result {
    if (c.Session["cUserRole"] == "1") {
        id,_ := strconv.Atoi(c.Request.Form.Get("id"));
        success := DeleteNewsByNewsid(Dbm,id);
        if (success){
            return c.Redirect(News.List);
        }else{
            return c.Redirect(App.Index);        
        }
    } else {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login")
    }
}

func (c News) EditForm() revel.Result {
    if (c.Session["cUserRole"] == "1") {
        var id int;
        c.Params.Bind(&id,"id");
        news :=  SelectNewsByNewsId(Dbm,id);
        return c.Render(news);
    } else {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login")
    }
}

func (c News) EditSubmit() revel.Result {
    if (c.Session["cUserRole"] == "1") {
        newsid,_ := strconv.ParseInt(c.Request.Form.Get("newsid"),0,64);
        newscreatedat,_ := strconv.ParseInt(c.Request.Form.Get("newscreated"),0,64);
        innews := models.News{
            NewsId : newsid,
            NewsTitle : c.Request.Form.Get("newstitle"),
            Content : c.Request.Form.Get("newscontent"),
            CreatedAt : newscreatedat,
            UpdatedAt : time.Now().UnixNano(), 
        }
        success := UpdateNews(Dbm, innews);
        if (success){
            return c.Redirect(News.List);
        }else{
            return c.Redirect(App.Index);        
        }
    } else {
        c.Flash.Error("You are not authorized!")
        return c.Redirect("/Login")
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

    _, err := dbm.Select(&p, "SELECT * FROM News ORDER BY news_updated_at DESC")
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

func SelectLatestNews(dbm *gorp.DbMap) models.News {
    var p models.News
    err := dbm.SelectOne(&p, "SELECT * FROM News ORDER BY news_updated_at DESC LIMIT 1")
    checkErr(err, "SelectOne failed")
    return p
}
