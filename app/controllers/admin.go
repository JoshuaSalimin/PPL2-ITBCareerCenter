package controllers

import (
	"github.com/revel/revel"
	// "log"
)

type Admin struct {
	*revel.Controller
}

func (c Admin) AdminNews() revel.Result {
	adminnews := true
	return c.Render(adminnews)
}