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

func (p Partnership) SavePartnership(partnershipID []int, partnershipName []string, partnershipLink []string) revel.Result {
    //Update Product Photos
   if ( (len(p.Params.Files["partnershipImg[]"]) != 0) || (len(partnershipID) != 0) || (len(partnershipName) != 0) || (len(partnershipLink) != 0) ) {
       var partnershipImg [][]byte
       p.Params.Bind(&partnershipImg, "partnershipImg")
       count := len(partnershipID)
       if(len(partnershipName) > count){
            count = len(partnershipName)
       }
       if(len(partnershipLink) > count){
            count = len(partnershipLink)
       }
       if(len(p.Params.Files["partnershipImg[]"]) > count){
            count = len(p.Params.Files["partnershipImg[]"])
       }
       for i := 0; i<count; i++ {
           var newPartnership models.Partnership
           if(len(p.Params.Files["partnershipImg[]"]) != 0){
                filename := p.Params.Files["partnershipImg[]"][i].Filename
                relativePath := p.UploadImagePartnership(partnershipImg[i], filename)
                newPartnership.ImgPath = relativePath  
           }
           newPartnership.PartnershipName = partnershipName[i]
           newPartnership.PartnershipLink = partnershipLink[i]
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