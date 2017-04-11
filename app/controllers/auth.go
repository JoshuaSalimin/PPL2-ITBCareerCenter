package controllers

import (
    //"PPL2-ITBCareerCenter/app/models"
"github.com/revel/revel"
    // "encoding/json"
    //"github.com/go-gorp/gorp"
"log"
)
import "strconv"

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	uname := c.Params.Form.Get("username")
	pwd := c.Params.Form.Get("password")
	log.Println("uname: " + uname + " | pwd: " + pwd)
	users := SelectAllUsers(Dbm)
	for _,user := range users {
		if (user.Username == uname) {
			if (user.Password == pwd) {
				c.Session["cUserId"] = strconv.FormatInt(user.UserId, 10)
				log.Println("MATCH !! => " + user.Username)
				log.Println("session: " + c.Session["cUserId"])
				return c.Redirect("/Login")
			} else {
				return c.Redirect("/Login")
			}
		} else {
			log.Println("NO MATCH => " + user.Username)
		}
	}
	return c.Redirect("/Login")
}

func (c Auth) Logout() revel.Result {
	delete(c.Session, "cUserId")
	return c.Redirect("/Login")
}

/*
func getCurrentUserId() int64 {
	result, err := strconv.ParseInt(Controller.Session["cUserId"], 10, 64)
	if (err != nil) {
		return -1;
	}
	return result;
}

func getCurrentUserRole() int {
	cUserId := getCurrentUserId();
	if (cUserId == -1) {
		return -1;
	}
	user := SelectUsersByUserid(Dbm, cUserId)
	return user.Role
}*/