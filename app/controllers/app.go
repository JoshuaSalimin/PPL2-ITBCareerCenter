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
	return c.Render(home)
}

func (c App) News() revel.Result {
	news := true
    list := SelectAllNews(Dbm);
    return c.Render(news,list);
}

func (c App) Event() revel.Result {
	event := true
	list := SelectAllEvent(Dbm)
	return c.Render(event, list)
}

func (c App) Profiles() revel.Result {
	profiles := true
	return c.Render(profiles)
}

func (c App) Contact() revel.Result {
	contact := true
	return c.Render(contact)
}


