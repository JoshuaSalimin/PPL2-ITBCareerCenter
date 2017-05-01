package controllers

import (
	"github.com/revel/revel"
	"log"
	"strconv"
)

type ChangePassword struct {
	*revel.Controller
}

func (c ChangePassword) ChangePassword() revel.Result {
	changePassword := true
	return c.Render(changePassword)
}

func (c ChangePassword) Test() revel.Result {
	ChangePasswordFailedMessage := "Change Password Fail : "
	oldpassword := c.Params.Form.Get("oldpassword")
	newpassword := c.Params.Form.Get("newpassword")
	confirmpassword := c.Params.Form.Get("confirmpassword")
	useridString := c.Session["cUserId"]
	userid, _ := strconv.Atoi(useridString)

	if ((oldpassword == "") || (newpassword == "") || (oldpassword == "")) {
		c.Flash.Error(ChangePasswordFailedMessage + "Isi semua kolom yang diperlukan");
		return c.Redirect("/ChangePassword")
	} else {
		User := SelectUsersByUserid(Dbm, userid)
	
		if (User.IsPasswordChanged == true) {
			oldpassword = EncryptSHA256(oldpassword)
		} else if (User.IsPasswordChanged == false) {
			User.IsPasswordChanged == true
		}

		if (User.Password != oldpassword) {
			c.Flash.Error(ChangePasswordFailedMessage + "Password lama salah");
		return c.Redirect("/ChangePassword")
		} else {
			if (newpassword != confirmpassword) {
				c.Flash.Error(ChangePasswordFailedMessage + "Konfirmasi password tidak dilakukan dengan benar");
				return c.Redirect("/ChangePassword")
			} else {
				User.Password = EncryptSHA256(newpassword)
				UpdateUsers(Dbm, User)
				c.Flash.Success("Change Password Success")
				return c.Redirect("/ChangePassword")
			}
		}
	}
	// return c.Render(test)
}

