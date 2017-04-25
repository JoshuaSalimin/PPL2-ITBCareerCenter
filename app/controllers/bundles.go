package controllers

import (
"PPL2-ITBCareerCenter/app/models"
"github.com/revel/revel"
//models "PPL2-ITBCareerCenter/app/models"
    // "encoding/json"
"github.com/go-gorp/gorp"
"time"
"log"
"strconv"
"os"
"encoding/csv"
//"math/rand"
)

type Bundles struct {
	*revel.Controller
}

func (c Bundles) Bundles() revel.Result {
	bundles := SelectAllBundles(Dbm)
	return c.Render(bundles)
}

func (c Bundles) AddView() revel.Result {
	return c.Render(true)
}

func (c Bundles) Add() revel.Result {
	timecreated := time.Now().UnixNano()
	jumlah,_ := strconv.Atoi(c.Request.Form.Get("jumlah"))
	namabundle := c.Request.Form.Get("namabundle")
	angkatan,_ := strconv.Atoi(c.Request.Form.Get("angkatan"))

    //Tambah bundle ke tabel bundle
	newbundle := models.Bundles {
		BundleId : 0,
		BundleName : namabundle,
	}
	InsertBundles(Dbm, &newbundle)
    //Tambah user-user ke tabel users
	for i:=1; i<=jumlah; i++ {
		username := namabundle + "-" + strconv.Itoa(i)
		password := generateRandomPassword(10);
		password = EncryptSHA256(password)
		newuser := models.Users {
			UserId: 0,
			Username: username,
			Password: password,
			Name            : "DEFAULT_NAME",  
			ProductName     : "DEFAULT_PRODUCT_NAME",  
			CompanyName     : "DEFAULT_COMPANY_NAME",  
			CompanyDescription     : "DEFAULT_DESCRIPTION",  
			Visi            : "DEFAULT_VISI",   
			Misi            : "DEFAULT_MISI",
			Jurusan         : "DEFAULT_JURUSAN",
			Angkatan        : angkatan,       
			LogoPath        : "DEFAULT_LOGO_PATH",
			CreatedAt       : timecreated, 
			UpdatedAt       : timecreated,
			ShowProfile     : false,    
			Role            : 0,     
		}

		InsertUsers(Dbm, &newuser)

        //Tambah user baru ke usersinbundle
		newUIB := models.UsersInBundle {
			UserId : newuser.UserId,
			BundleId : newbundle.BundleId,
		}
		InsertUsersInBundle(Dbm, &newUIB)
		log.Println(newuser.UserId)
	}

	c.Flash.Success(c.Request.Form.Get("jumlah") + " Users added successfully");
	return c.Redirect("/Users/Bundles")
}

func (c Bundles) Delete() revel.Result {
	bundleid,_ := strconv.Atoi(c.Request.Form.Get("id"))
	users := SelectUIBByBundleId(Dbm, bundleid)
	for _, user := range users {
		DeleteUIBByUserId(Dbm, int(user.UserId))
		DeleteUsersByUserId(Dbm, int(user.UserId))
	}
	DeleteBundleByBundleId(Dbm, bundleid)
	c.Flash.Success("Bundle deleted successfully")
	return c.Redirect("/Users/Bundles")
}

func (c Bundles) DownloadCSV() revel.Result {
	bundleid,_ := strconv.Atoi(c.Request.Form.Get("id"))

	//create parent folder
	parentfolder := "csv/"
	os.Mkdir(parentfolder, os.FileMode(0522))

	//generate filename
	bundle := SelectBundlesByBundleid(Dbm, bundleid)
	prefix := "Bundle-"
	filename := parentfolder + prefix + bundle.BundleName + ".csv"

	//data yang akan ditulis di dalam csv
	var data = [][]string{{"UserId", "Username", "Password"}}
	uibs := SelectUIBByBundleId(Dbm, bundleid)

	//generate data yang akan ditulis di dalam csv
	for _, uib := range uibs {
		user := SelectUsersByUserid(Dbm, int(uib.UserId))
		datum := []string{strconv.FormatInt(user.UserId,10), user.Username}
		if (user.IsPasswordChanged) {
			datum = append(datum, "-")
		} else {
			datum = append(datum, user.Password)
		}
		data = append(data, datum)
	}

	//buat file CSV nya
	file, err := os.Create(filename)
	checkErr(err, "cannot create file")

	//tulis file CSV nya
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, datum := range data {
		err := writer.Write(datum)
		checkErr(err, "Cannot write to file")
	}

	return c.RenderFile(file, "attachment")
}

func SelectAllBundles(dbm *gorp.DbMap) []models.Bundles {
	var bundles_dummy []models.Bundles

	_, err := dbm.Select(&bundles_dummy, "SELECT * FROM bundles")
	checkErr(err, "Select failed")
	log.Println("All rows:")
	for x, p := range bundles_dummy {
		log.Printf("    %d: %v\n", x, p)
	}
	return bundles_dummy
}

func SelectBundlesByBundleid(dbm *gorp.DbMap, bundleid int) models.Bundles {
	var b models.Bundles
	err := dbm.SelectOne(&b, "SELECT * FROM Bundles WHERE bundleid=?", bundleid)
	checkErr(err, "SelectOne failed")
	log.Println("b :", b)
	return b
}

func InsertBundles(dbm *gorp.DbMap, bundles *models.Bundles){
	err := dbm.Insert(bundles)
	checkErr(err, "Insert failed")
}

func DeleteBundleByBundleId(dbm *gorp.DbMap, bundleid int) {
	_, err := dbm.Exec("DELETE FROM Bundles WHERE bundleid=?", bundleid)
	checkErr(err, "Delete failed")
}