package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    "github.com/revel/revel"
    "encoding/json"
)

type UserCtrl struct {
    GorpController
}


func (c UserCtrl) parseUsers() (models.Users, error) {
    users := models.Users{}
    err := json.NewDecoder(c.Request.Body).Decode(&users)
    return users, err
}

func (c UserCtrl) Add() revel.Result {
    if users, err := c.parseUsers(); err != nil {
        return c.RenderText("Unable to parse the User from JSON.")
    } else {
        // Validate the model
        users.Validate(c.Validation)
        if c.Validation.HasErrors() {
            // Do something better here!
            return c.RenderText("You have error in your User.")
        } else {
            if err := c.Txn.Insert(&users); err != nil {
                return c.RenderText(
                    "Error inserting record into database!")
            } else {
                return c.RenderJson(users)
            }
        }
    }
}

func (c UserCtrl) Get(id int64) revel.Result {
    users := new(models.Users)
    err := c.Txn.SelectOne(users, 
        `SELECT * FROM Users WHERE id = ?`, id)
    if err != nil {
        return c.RenderText("Error. Users doesn't exist.")
    }
    return c.RenderJson(users)
}

func (c UserCtrl) List() revel.Result {
    lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
    limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
    users, err := c.Txn.Select(models.Users{}, 
        `SELECT * FROM Users WHERE Id > ? LIMIT ?`, lastId, limit)
    if err != nil {
        return c.RenderText(
            "Error trying to get records from DB.")
    }
    return c.RenderJson(users)
}

func (c UserCtrl) Update(id int64) revel.Result {
    users, err := c.parseUsers()
    if err != nil {
        return c.RenderText("Unable to parse the Users from JSON.")
    }
    // Ensure the Id is set.
    users.Id = id
    success, err := c.Txn.Update(&users)
    if err != nil || success == 0 {
        return c.RenderText("Unable to update users.")
    }
    return c.RenderText("Updated %v", id)
}

func (c UserCtrl) Delete(id int64) revel.Result {
    success, err := c.Txn.Delete(&models.Users{Id: id})
    if err != nil || success == 0 {
        return c.RenderText("Failed to remove Users")
    }
    return c.RenderText("Deleted %v", id)
}