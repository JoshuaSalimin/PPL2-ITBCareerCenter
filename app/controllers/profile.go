package controllers

import (
	"github.com/revel/revel"
	"PPL2-ITBCareerCenter/app/models"
	"html"
	"time"
	"log"
	"os"
	"fmt"
	"path/filepath"
)

type Profile struct {
	App
}

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func (c Profile) List(page int) revel.Result {
	profiles := true
	numUserPerPage := 6
	if (page == 0) {
		page = 1
	}
	startUserLimit := (page-1)*numUserPerPage
	endUserLimit := page*numUserPerPage

	userCount := CountUsers(Dbm)

	startUserLimit = max(startUserLimit, 0)

	if (startUserLimit >= userCount) {
		return c.NotFound("Invalid Page: ", page);
	}

	endUserLimit = min(userCount, endUserLimit)

	users := SelectLatestUsersInRange(Dbm, startUserLimit, endUserLimit - startUserLimit)
	currentPageNum := page
	return c.Render(profiles, page, users, userCount, numUserPerPage, currentPageNum)
}

func (c Profile) UploadLogo(id int, companylogo []byte) string {
	c.Validation.MaxSize(companylogo, 2*MB).
		Message("File cannot be larger than 2MB")

	filename := c.Params.Files["companylogo"][0].Filename
	fileExt := filepath.Ext(filename)
	randFilename := randString() + fileExt
	relativePath := fmt.Sprintf("/public/images/user/%d/%s", id, randFilename)
	dstPath := fmt.Sprintf("%s/public/images/user/%d", revel.BasePath, id)
	if _, err := os.Stat(dstPath); os.IsNotExist(err) {
	    os.Mkdir(dstPath, 0777)
	}
	dstPath = dstPath + "/" + randFilename
	dstFile, _ := os.Create(dstPath)
	defer dstFile.Close()
	defer os.Chmod(dstPath, (os.FileMode)(0644))

	dstFile.Write(companylogo)
	return relativePath
}

func (c Profile) UpdateSocialMedia(id int, socialMediaTypes []string,  socialMediaURLs []string) {
	oldUserSocialMedias := SelectAllUserSocialMediaByUserID(Dbm, id)
	for _,oldUserSocialMedia := range oldUserSocialMedias {
	  DeleteUserSocialMediaByUserSocialMediaid(Dbm, oldUserSocialMedia.UserSocialMediaId)
	}
	newUserSocialMedias := make([]models.UserSocialMedia, len(socialMediaTypes))
	for index, socialMediaType := range socialMediaTypes {
	  socialMediaType = html.EscapeString(socialMediaType)
	  newUserSocialMedias[index] = models.CreateDefaultUserSocialMedia()
	  newUserSocialMedias[index].SocialMediaName = socialMediaType
	  newUserSocialMedias[index].UserId = id
	}
	for index, socialMediaURL := range socialMediaURLs {
	  socialMediaURL = html.EscapeString(socialMediaURL)
	  newUserSocialMedias[index].SocialMediaURL = socialMediaURL
	}
	for _,newUserSocialMedia := range newUserSocialMedias {
		InsertUserSocialMedia(Dbm, newUserSocialMedia)
	}

}

func (c Profile) Edit(id int, user models.Users, socialMediaTypes []string,  socialMediaURLs []string, companylogo []byte) revel.Result {
	//TODO sanitize input
	oldUser := SelectUsersByUserid(Dbm, id)
	oldUser.LogoPath = c.UploadLogo(id, companylogo)
	oldUser.CompanyName = user.CompanyName
	oldUser.Name = user.Name
	oldUser.CompanyDescription = user.CompanyDescription
	oldUser.Visi = user.Visi
	oldUser.Misi = user.Misi
	oldUser.Jurusan = user.Jurusan
	oldUser.Angkatan = user.Angkatan
	oldUser.UpdatedAt = time.Now().UnixNano()
	c.UpdateSocialMedia(id, socialMediaTypes, socialMediaURLs)
	UpdateUsers(Dbm, oldUser)

	var productphotos [][]byte
	c.Params.Bind(&productphotos, "productphotos")

	for k, v := range c.Params.Files { 
	    log.Println("key[%s] value[%s]\n", k, v)
	}

	// for i, _ := range productphotos {
	// 	// contentType := c.Params.Files["productphotos[]"][i].Header.Get("Content-Type")
	// 	filename := c.Params.Files["productphotos[]"][i].Filename
	// 	// size := len(productphotos[i])

	// 	c.Validation.MaxSize(productphotos[i], 2*MB).
	// 		Message("File cannot be larger than 2MB")

	// 	ioutil.WriteFile("tmp/" + filename, productphotos[i], 0644)
	// }

	return c.Redirect("/ProfilePage/%d", id)
}

func (c Profile) Form(id int) revel.Result {
	profiles := true
	user := SelectUsersByUserid(Dbm, id)
	userSocialMedias := SelectAllUserSocialMediaByUserID(Dbm, id)
	return c.Render(user, profiles, userSocialMedias)
}

func (c Profile) Page(id int) revel.Result {
	profiles := true
	authorized := false
	user := SelectUsersByUserid(Dbm, id)
	namaPerusahaan := user.CompanyName
	deskripsiPerusahaan := user.CompanyDescription
	visiPerusahaan := user.Visi
	misiPerusahaan := user.Misi
	namaPemilik := user.Name
	jurusan := user.Jurusan
	angkatanPMW := user.Angkatan
	userUpdatedAt := time.Unix(0, user.UpdatedAt)
	userContact := SelectAllUserContactByUserId(Dbm, id)
	userSocialMedias := SelectAllUserSocialMediaByUserID(Dbm, id)
	userVideo := SelectVideoByUserId(Dbm, id)
	userImage := SelectUserImage(Dbm, id)
	return c.Render(id, profiles, namaPerusahaan, deskripsiPerusahaan, visiPerusahaan, misiPerusahaan, namaPemilik, jurusan, angkatanPMW, userContact, userSocialMedias, authorized, userVideo, userImage, userUpdatedAt)
}
