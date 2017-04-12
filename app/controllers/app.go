package controllers

import (
	"github.com/revel/revel"
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
	return c.Render(news)
}

func (c App) Profiles() revel.Result {
	profiles := true
	return c.Render(profiles)
}

func (c App) Partner() revel.Result {
	partner := true
	return c.Render(partner)
}

func (c App) About() revel.Result {
	about := true
	return c.Render(about)
}

func (c App) Contact() revel.Result {
	contact := true
	return c.Render(contact)
}

func (c App) Login() revel.Result {
	login := true
	return c.Render(login)
}