package controllers

import (
"PPL2-ITBCareerCenter/app/models"
"github.com/revel/revel"
//models "PPL2-ITBCareerCenter/app/models"
    // "encoding/json"
"github.com/go-gorp/gorp"
    // "time"
"log"
//"time"
//"strconv"
//"math/rand"
)

type UsersInBundle struct {
    *revel.Controller
}

func InsertUsersInBundle(dbm *gorp.DbMap, uib *models.UsersInBundle){
    err := dbm.Insert(uib)
    checkErr(err, "Insert failed")
}

func SelectUIBByBundleId(dbm *gorp.DbMap, bundleid int) []models.UsersInBundle {
    var uib []models.UsersInBundle

    _, err := dbm.Select(&uib, "SELECT * FROM usersinbundle WHERE bundleid=?", bundleid)
    checkErr(err, "Select failed")
    log.Println("User Range rows:")
    for x, p := range uib {
        log.Printf("    %d: %v\n", x, p)
    }
    return uib 
}

func DeleteUIBByUserId(dbm *gorp.DbMap, userid int) {
    _, err := dbm.Exec("DELETE FROM UsersInBundle WHERE userid=?", userid)
    checkErr(err, "Delete failed")
}