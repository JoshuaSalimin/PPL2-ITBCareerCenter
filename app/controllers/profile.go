package controllers

import (
	"github.com/revel/revel"
	"PPL2-ITBCareerCenter/app/models"
	"html"
)

type Profile struct {
	App
}

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

func (c Profile) Edit(id int, user models.Users, socialMediaTypes []string,  socialMediaURLs []string) revel.Result {

	//Update Social Media
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

	//TODO sanitize input
	oldUser := SelectUsersByUserid(Dbm, id)
	oldUser.CompanyName = user.CompanyName
	oldUser.Name = user.Name
	oldUser.CompanyDescription = user.CompanyDescription
	oldUser.Visi = user.Visi
	oldUser.Misi = user.Misi
	oldUser.Jurusan = user.Jurusan
	oldUser.Angkatan = user.Angkatan
	UpdateUsers(Dbm, oldUser)
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
	userContact := SelectAllUserContactByUserId(Dbm, id)
	userSocialMedias := SelectAllUserSocialMediaByUserID(Dbm, id)
	userVideo := SelectVideoByUserId(Dbm, id)
	userImage := SelectUserImage(Dbm, id)
	return c.Render(id, profiles, namaPerusahaan, deskripsiPerusahaan, visiPerusahaan, misiPerusahaan, namaPemilik, jurusan, angkatanPMW, userContact, userSocialMedias, authorized, userVideo, userImage)
}
