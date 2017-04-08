package controllers

import (
    "PPL2-ITBCareerCenter/app/models"
    // "github.com/revel/revel"
    // "encoding/json"
    "github.com/go-gorp/gorp"
    "log"
)

func InsertUserSocialMedia(dbm *gorp.DbMap, p models.UserSocialMedia) {
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func SelectAllUserSocialMedia(dbm *gorp.DbMap) []models.UserSocialMedia {
	var p []models.UserSocialMedia

    _, err := dbm.Select(&p, "SELECT * FROM usersocialmedia")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range p {
        log.Printf("    %d: %v\n", x, p)
    }
    return p 	
}

func SelectUserSocialMediaByUserSocialMediaId(dbm *gorp.DbMap, socialmediaaid int) models.UserSocialMedia {
	var p models.UserSocialMedia
    err := dbm.SelectOne(&p, "SELECT * FROM usersocialmedia WHERE socialmediaid=?", socialmediaaid)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func UpdateUserSocialMedia(dbm *gorp.DbMap, p models.UserSocialMedia) {
	count, err := dbm.Update(&p)
	checkErr(err, "Update failed")	
    log.Println("Rows updated:", count)
}


func DeleteUserSocialMediaByUserSocialMediaid(dbm *gorp.DbMap, socialmediaaid int) {
    _, err := dbm.Exec("DELETE FROM usersocialmedia WHERE socialmediaid=?", socialmediaaid)
    checkErr(err, "Delete failed")
}
