package controllers

import (
    "github.com/revel/revel"
    "github.com/go-gorp/gorp"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    "PPL2-ITBCareerCenter/app/models"
    // "time"
)

func init(){
	revel.OnAppStart(InitDb)	
    revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
    revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
    revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
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

    err := Dbm.CreateTablesIfNotExists()
    checkErr(err, "Create Table failed")
    err = Dbm.CreateIndex();
    checkErr(err, "Create Index Failed")
    // content of news must be changed to text instead of varchar(255) because 
    // varchar(255) is not enough to contain it
    _, err = Dbm.Exec(" ALTER TABLE News MODIFY content text")
    checkErr(err, "ALTER TABLE News FAILED")
    _, err = Dbm.Exec("ALTER TABLE users ADD UNIQUE (username);")
    checkErr(err, "ALTER TABLE users FAILED")
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


    //u := models.CreateDefaultUser("ramos")

    // KEY := []byte("key")
    // unencryptedPassword := "this is password you will encrypt"
    // encryptedPassword, err := encrypt(KEY, unencryptedPassword) 
    // u.Password = encryptedPassword

    //InsertUsers(Dbm, u)

    // USAGE EXAMPLE --------------
    // InsertUsersAdmin(Dbm)
    // InsertPostAdmin(Dbm)

    // NewsTemp := models.CreateDefaultNews("News Title Temp")
    // InsertNews(Dbm, NewsTemp)
    // UsersSocialMedia := models.CreateDefaultUserSocialMedia()
    // n := SelectNewsByNewsId(Dbm, 2)
    // SelectAllNews(Dbm)
    // UpdateNews(Dbm, n)
    // DeleteNewsByNewsid(Dbm, 2)


    // InsertUserSocialMedia(Dbm, UsersSocialMedia)
    // UsersContact := models.CreateDefaultUserContact()
    // sm := SelectUserSocialMediaByUserSocialMediaId(Dbm, 2)
    // SelectAllUserSocialMedia(Dbm)
    // UpdateUserSocialMedia(Dbm, sm)
    // DeleteUserSocialMediaByUserSocialMediaid(Dbm, 2)


    // InsertUserContact(Dbm, UsersContact)
    // u := SelectUserContactByUserContactId(Dbm, 2)
    // SelectAllUserContact(Dbm)
    // UpdateUserContact(Dbm, u)
    // DeleteUserContactByUserContactid(Dbm, 2)

    // ----------------------------------------------
}


func defineUserTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    t := dbm.AddTable(models.Users{}).SetKeys(true, "userid")
    //t.AddIndex("username_idx","BTree",[]string{"Username"}).SetUnique(true);

    // e.g. VARCHAR(25)
    t.ColMap("name").SetMaxSize(25)
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