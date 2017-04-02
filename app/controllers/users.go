package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
)

type UserCtrl struct {
    GorpController
}

func addUsers(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
	admin := &models.Users{
		Id: 1,
		Username: "admin",
		Name: "admin",
		Password: "password",
		Role: 1,
	}
    dbm.Insert(admin)
}