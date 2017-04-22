package controllers

import (
    "github.com/revel/revel"
    "github.com/go-gorp/gorp"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    "PPL2-ITBCareerCenter/app/models"
    "time"
    "math/rand"
    // "log"
)

func init(){
	revel.OnAppStart(InitDb)	
    revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
    revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
    revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
    revel.TemplateFuncs["convert_unix_time"] = func(unixtime int64) string {
        return time.Unix(0,unixtime).Format("2006-01-02, 15:04:05");
    } 
}

func getParamString(param string, defaultValue string) string {
    p, found := revel.Config.String(param)
    if !found {
        if defaultValue == "" {
            revel.ERROR.Fatal("Cound not find parameter: " + param)
        } else {
            return defaultValue
        }
    }
    return p
}

func getConnectionString() string {
    host := getParamString("db.host", "")
    port := getParamString("db.port", "3306")
    user := getParamString("db.user", "")
    pass := getParamString("db.password", "")
    dbname := getParamString("db.name", "career_center")
    protocol := getParamString("db.protocol", "tcp")
    dbargs := getParamString("dbargs", " ")

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s", 
        user, pass, protocol, host, port, dbname, dbargs)
}

var InitDb func() = func(){
    rand.Seed(time.Now().UTC().UnixNano())
    connectionString := getConnectionString()
    if db, err := sql.Open("mysql", connectionString); err != nil {
        revel.ERROR.Fatal(err)
    } else {
        Dbm = &gorp.DbMap{
            Db: db, 
            Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
    }
    // Defines the table for use by GORP
    // This is a function we will create soon.
    defineUserTable(Dbm)
    definePostTable(Dbm)
    defineNewsTable(Dbm)
    defineUserSocialMediaTable(Dbm)
    defineUserContactTable(Dbm)
    defineAboutTable(Dbm)
    defineContactTable(Dbm)
    definePartnershipTable(Dbm)
    defineUsersInBundleTable(Dbm)
    defineBundlesTable(Dbm)
    defineEventTable(Dbm)

    err := Dbm.CreateTablesIfNotExists()
    checkErr(err, "Create Table failed")
    err = Dbm.CreateIndex();
    checkErr(err, "Create Index Failed")
    // content of news must be changed to text instead of varchar(255) because 
    // varchar(255) is not enough to contain it
    _, err = Dbm.Exec(" ALTER TABLE News MODIFY content text")
    checkErr(err, "ALTER TABLE News FAILED")
    _, err = Dbm.Exec("ALTER TABLE users ADD UNIQUE (username)")
    checkErr(err, "ALTER TABLE users FAILED")
    _, err = Dbm.Exec("ALTER TABLE event MODIFY event_description text")
    checkErr(err, "ALTER TABLE event FAILED")
    countAbout := CountAbout(Dbm)
    if(countAbout >= 1){
        // do nothing
    } else{
        newAbout := models.CreateDefaultAbout()
        newAbout.AboutID = 0
        InsertAbout(Dbm, newAbout)
    }
    countContact := CountContact(Dbm)
    if(countContact >= 1){
        // do nothing
    } else{
        newContact := models.CreateDefaultContact()
        newContact.ContactID = 0
        InsertContact(Dbm, newContact)
    }

    // e1 := models.CreateDefaultEvent("Event1")
    // e2 := models.CreateDefaultEvent2("Event2")
    // e3 := models.CreateDefaultEvent3("Event3")
    // InsertEvent(Dbm, e1)
    // InsertEvent(Dbm, e2)
    // InsertEvent(Dbm, e3)
}


func defineUserTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    t := dbm.AddTable(models.Users{}).SetKeys(true, "userid")
    //t.AddIndex("username_idx","BTree",[]string{"Username"}).SetUnique(true);

    // e.g. VARCHAR(25)
    t.ColMap("name").SetMaxSize(25)
}

func defineBundlesTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    dbm.AddTable(models.Bundles{}).SetKeys(true, "bundleid")
}

func defineUsersInBundleTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    dbm.AddTable(models.UsersInBundle{}).SetKeys(false, "userid", "bundleid")
}

func defineUserSocialMediaTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.UserSocialMedia{}).SetKeys(true, "socialmediaid")    
}

func defineUserContactTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.UserContact{}).SetKeys(true, "contactid")    
}

func definePostTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    dbm.AddTable(models.Post{}).SetKeys(true, "postid") 
}

func defineNewsTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.News{}).SetKeys(true, "newsid")    
}

func defineAboutTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.About{}).SetKeys(true, "aboutid")    
}

func defineContactTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.Contact{}).SetKeys(true, "contactid")    
}

func definePartnershipTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.Partnership{}).SetKeys(true, "partnershipid")    
}

func defineEventTable(dbm *gorp.DbMap) {
    dbm.AddTable(models.Event{}).SetKeys(true,"eventid")
}
