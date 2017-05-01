package controllers

import (
    //"PPL2-ITBCareerCenter/app/models"
"github.com/revel/revel"
    // "encoding/json"
    //"github.com/go-gorp/gorp"
//"log"
)
import "strconv"

type Auth struct {
	*revel.Controller
}

func (c Auth) Login() revel.Result {
	loginFailedMsg := "Login Failed: "
	//log.Println(getCurrentUserId(c));

	uname := c.Params.Form.Get("username")
	pwd := c.Params.Form.Get("password")

	user := SelectUserByUsername(Dbm, uname)

	if (u.IsPasswordChanged == true) {
		pwd = EncryptSHA256(pwd)
	}

	if (uname == "") {
		c.Flash.Error(loginFailedMsg + "Username harus diisi");
		return c.Redirect("/Login")
	}

	if (user.Username == uname && user.Password == pwd) {
		//Saves currently logged in user id in session, NIL OTHERWISE
		c.Session["cUserId"] = strconv.FormatInt(user.UserId, 10)
		//Saves currently logged in user's role in session, NIL OTHERWISE
		user := SelectUsersByUserid(Dbm, int(user.UserId))
		c.Session["cUserRole"] = strconv.Itoa(user.Role)
		c.Flash.Success("Login Successful");
		return c.Redirect("/")
	} else {
		c.Flash.Error(loginFailedMsg + "Username atau Password salah");
	}
	return c.Redirect("/Login")
}

func (c Auth) LogoutView() revel.Result {
	return c.Render(true)
}

func (c Auth) Logout() revel.Result {
	delete(c.Session, "cUserId")
	delete(c.Session, "cUserRole")
	return c.Redirect("/")
}

func (c Auth) LoginView() revel.Result {
	userid := c.Session["cUserId"]
	if (userid == "") {
		return c.Render()
	} else {
		return c.Redirect("/Logout")
	}
}

/*
func getCurrentUserId(c Auth) int64 {
	result, err := strconv.ParseInt(c.Session["cUserId"], 10, 64)
	if (err != nil) {
		return -1;
	}
	return result;
}

func getCurrentUserRole(c Auth) int {
	cUserId := getCurrentUserId(c);
	if (cUserId == -1) {
		return -1;
	}
	user := SelectUserByUserId(Dbm, cUserId)
	return user.Role
}*/