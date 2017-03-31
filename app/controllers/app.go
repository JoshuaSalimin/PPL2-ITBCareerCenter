package controllers

import (
	"github.com/revel/revel"
	"fmt"
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

func (c App) Articles() revel.Result {
	articles := true
	return c.Render(articles)
}

func (c App) Files() revel.Result {
	files := true
	return c.Render(files)
}

func (c App) Photos() revel.Result {
	photos := true
	return c.Render(photos)
}

func (c App) Videos() revel.Result {
	videos := true
	return c.Render(videos)
}

func (c App) Profiles() revel.Result {
	profiles := true
	return c.Render(profiles)
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

func (c App) EditProfiles(myName1 string) revel.Result {
	fmt.Printf(myName1)
	return c.Redirect(App.Profiles)
}

func (c App) ProfilesForm() revel.Result {
	return c.Render()
}