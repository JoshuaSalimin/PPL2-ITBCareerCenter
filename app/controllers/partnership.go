package controllers

import (
    "github.com/revel/revel"
    models "PPL2-ITBCareerCenter/app/models"
    "github.com/go-gorp/gorp"
     "log"
     //"time"
     //"strconv"
     "path/filepath"
     "os"
     "fmt"
)

const (
   _      = iota
   KB int = 1 << (10 * iota)
   MB
   GB
)

type Partnership struct {
    *revel.Controller
}

func (p Partnership) Partnership() revel.Result {
    partnership := true
    allPartnership := SelectAllPartnership(Dbm)
    return p.Render(partnership, allPartnership)
}

func (p Partnership) EditPartnership() revel.Result {
    partnership := true
    allPartnership := SelectAllPartnership(Dbm)
    return p.Render(partnership, allPartnership)
}

func InsertPartnership(dbm *gorp.DbMap, p models.Partnership){
    err := dbm.Insert(&p)
    checkErr(err, "Insert failed")
}

func PartnershiptoDB(dbm *gorp.DbMap, p models.Partnership){
    count, err := dbm.Update(&p)
    checkErr(err, "Update failed")  
    log.Println("Rows updated:", count)
}

func (p Partnership) SavePartnership(partnershipName []string, partnershipLink []string) revel.Result {
    //Update Product Photos
   if (len(p.Params.Files["partnershipImg[]"]) != 0) {
       var partnershipImg [][]byte
       p.Params.Bind(&partnershipImg, "partnershipImg")

       for i, _ := range partnershipName {
           filename := p.Params.Files["partnershipImg[]"][i].Filename
           relativePath := p.UploadImagePartnership(partnershipImg[i], filename)
           newPartnership := models.CreateDefaultPartnership()
           newPartnership.PartnershipName = partnershipName[i]
           newPartnership.PartnershipLink = partnershipLink[i]
           newPartnership.ImgPath = relativePath  
           InsertPartnership(Dbm, newPartnership)
       }
   }
    return p.Redirect(Partnership.Partnership);
}

func (p Partnership) UploadImagePartnership(image []byte, filename string) string {
   p.Validation.MaxSize(image, 2*MB).
       Message("File cannot be larger than 2MB")
   fileExt := filepath.Ext(filename)
   randFilename := randString() + fileExt
   relativePath := fmt.Sprintf("/public/images/partnership/%s", randFilename)
   dstPath := fmt.Sprintf("%s/public/images/partnership", revel.BasePath)
   if _, err := os.Stat(dstPath); os.IsNotExist(err) {
       os.Mkdir(dstPath, 0777)
   }
   dstPath = dstPath + "/" + randFilename
   dstFile, _ := os.Create(dstPath)
   defer dstFile.Close()
   defer os.Chmod(dstPath, (os.FileMode)(0644))
   dstFile.Write(image)
   return relativePath
}

/*
func (p Partnership) SavePartnership() revel.Result {
    aboutid,_ := strconv.ParseInt(a.Request.Form.Get("aboutid"),0,64);
    newAbout := models.About{
        AboutID            : aboutid,
        CkeditorAbout      : a.Request.Form.Get("aboutcontent"),  
        UpdatedAt          : time.Now().UnixNano(),     
    }
    PartnershiptoDB(Dbm, newAbout);
    return a.Redirect(About.About);
}*/

func SelectPartnershipByPartnershipID(dbm *gorp.DbMap, id int) models.Partnership {
    var p models.Partnership
    err := dbm.SelectOne(&p, "SELECT * FROM Partnership WHERE partnershipid=?", id)
    checkErr(err, "SelectOne failed")
    log.Println("p :", p)
    return p
}

func CountPartnership(dbm *gorp.DbMap) int {
    count, err := dbm.SelectInt("SELECT COUNT(*) FROM Partnership")
    checkErr(err, "Select failed")
    log.Println("Partnership count:", count)
    return int(count)
}

func SelectAllPartnership(dbm *gorp.DbMap) []models.Partnership {
    var p []models.Partnership
    _, err := dbm.Select(&p, "SELECT * FROM Partnership")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, a := range p {
        log.Printf("    %d: %v\n", x, a)
    }
    return p   
}