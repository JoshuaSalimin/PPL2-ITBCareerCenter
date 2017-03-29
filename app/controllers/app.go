package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) News() revel.Result {
	return c.Render()
}

func (c App) Articles() revel.Result {
	return c.Render()
}

func (c App) Files() revel.Result {
	return c.Render()
}

func (c App) Photos() revel.Result {
	return c.Render()
}

func (c App) Videos() revel.Result {
	return c.Render()
}

func (c App) Profiles() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}