package controllers

import (
	"github.com/revel/revel"
	"PPL2-ITBCareerCenter/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	home := true
	return c.Render(home)
}

func (c App) News() revel.Result {
	news := true
	return c.Render(news)
}

func (c App) Articles() revel.Result {
	articles := true
	return c.Render(articles)
}

func (c App) Files() revel.Result {
	files := true
	return c.Render(files)
}

func (c App) Photos() revel.Result {
	photos := true
	return c.Render(photos)
}

func (c App) Videos() revel.Result {
	videos := true
	return c.Render(videos)
}

func (c App) Profiles() revel.Result {
	profiles := true
	return c.Render(profiles)
}

func (c App) About() revel.Result {
	about := true
	return c.Render(about)
}

func (c App) Contact() revel.Result {
	contact := true
	return c.Render(contact)
}

func (c App) Login() revel.Result {
	login := true
	return c.Render(login)
}

func (c App) ListProfiles(page int) revel.Result {
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

func (c App) EditProfiles(id int, user models.Users) revel.Result {
	oldUser := SelectUsersByUserid(Dbm, id)
	oldUser.CompanyName = user.CompanyName
	UpdateUsers(Dbm, oldUser)
	return c.Redirect("/ProfilePage/%d", id)
}

func (c App) ProfilesForm(id int) revel.Result {
	user := SelectUsersByUserid(Dbm, id)
	return c.Render(user)
}

func (c App) ProfilePage(id int) revel.Result {
	profiles := true
	user := SelectUsersByUserid(Dbm, id)
	namaPerusahaan := user.CompanyName
	deskripsiPerusahaan := user.CompanyDescription
	visiPerusahaan := user.Visi
	misiPerusahaan := user.Misi
	namaPemilik := user.Name
	jurusan := user.Jurusan
	angkatanPMW := user.Angkatan
	userContact := SelectAllUserContactByUserId(Dbm, id)
	userSocialMedia := SelectAllUserSocialMediaByUserID(Dbm, id)
	return c.Render(id, profiles, namaPerusahaan, deskripsiPerusahaan, visiPerusahaan, misiPerusahaan, namaPemilik, jurusan, angkatanPMW, userContact, userSocialMedia)

}