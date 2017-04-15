package controllers

import (
	"github.com/revel/revel"
	"PPL2-ITBCareerCenter/app/models"
	"html"
	"time"
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

func (c Profile) UploadImage(id int, image []byte, filename string) string {
	c.Validation.MaxSize(image, 2*MB).
		Message("File cannot be larger than 2MB")

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

	dstFile.Write(image)
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

func (c Profile) UpdateVideoPost(id int, videoID string) {
	oldVideoPost := SelectVideoByUserId(Dbm, id)
	newVideoPost := models.CreateDefaultPost(videoID)
	newVideoPost.MediaType = "Video"
	newVideoPost.UserId = int64(id)
	newVideoPost.PathFile = "https://www.youtube.com/embed/" + videoID
	
	if (oldVideoPost.PathFile != "DEFAULT_PATH_FILE") {
		newVideoPost.CreatedAt = oldVideoPost.CreatedAt
		DeletePostByPostid(Dbm, oldVideoPost.PostId)
	}
	InsertPost(Dbm, newVideoPost)
}

func (c Profile) Edit(id int, user models.Users, socialMediaTypes []string,  socialMediaURLs []string, companylogo []byte, videoID string) revel.Result {
	//TODO sanitize input
	oldUser := SelectUsersByUserid(Dbm, id)
	if (len(companylogo) != 0) {
		filename := c.Params.Files["companylogo"][0].Filename
		oldUser.LogoPath = c.UploadImage(id, companylogo, filename)
	}
	oldUser.CompanyName = user.CompanyName
	oldUser.Name = user.Name
	oldUser.CompanyDescription = user.CompanyDescription
	oldUser.Visi = user.Visi
	oldUser.Misi = user.Misi
	oldUser.Jurusan = user.Jurusan
	oldUser.Angkatan = user.Angkatan
	oldUser.UpdatedAt = time.Now().UnixNano()
	UpdateUsers(Dbm, oldUser)

	c.UpdateSocialMedia(id, socialMediaTypes, socialMediaURLs)
	if (videoID != "") {
		c.UpdateVideoPost(id, videoID)
	} else {
		oldVideoPost := SelectVideoByUserId(Dbm, id)
		DeletePostByPostid(Dbm, oldVideoPost.PostId)
	}


	var productphotos [][]byte
	c.Params.Bind(&productphotos, "productphotos")

	for i, _ := range productphotos {
		filename := c.Params.Files["productphotos[]"][i].Filename
		relativePath := c.UploadImage(id, productphotos[i], filename)
		newImagePost := models.CreateDefaultPost(filename)
		newImagePost.MediaType = "Image"
		newImagePost.UserId = int64(id)
		newImagePost.PathFile = relativePath		
		InsertPost(Dbm, newImagePost)
	}

	return c.Redirect("/ProfilePage/%d", id)
}

func (c Profile) Form(id int) revel.Result {
	profiles := true
	user := SelectUsersByUserid(Dbm, id)
	userSocialMedias := SelectAllUserSocialMediaByUserID(Dbm, id)
	userVideo := SelectVideoByUserId(Dbm, id)
	userImages := SelectUserImage(Dbm, id)
	return c.Render(user, profiles, userSocialMedias, userVideo, userImages)
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
