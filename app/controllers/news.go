package controllers

import (
	"github.com/revel/revel"
)

type News struct {
	*revel.Controller
}

func (c News) Form() revel.Result {
	return c.Render()
}

func (c News) Add() revel.Result {

	return c.Redirect(News.Form)
}
