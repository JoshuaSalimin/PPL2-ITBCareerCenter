package controllers

import (
	"github.com/revel/revel"
	// "log"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	home := true
	usersShown := SelectAllShownUsers(Dbm)
	return c.Render(home, usersShown)
}

func (c App) Profiles() revel.Result {
	profiles := true
	return c.Render(profiles)
}

func (c App) Contact() revel.Result {
	contact := true
	return c.Render(contact)
}


