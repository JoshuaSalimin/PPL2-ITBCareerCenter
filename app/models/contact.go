package models 

import (
    // "github.com/go-gorp/gorp"
    //"regexp"
    "time"
)

type Contact struct {
    ContactID               int64   `db:"contactid"`
	Deskripsi		        string 	`db:"deskripsi"`
    AlamatITB               string  `db:"alamatITB"`
    Alamat                  string  `db:"alamat"`
    Email                   string  `db:"email"`		
	Telepon		            string 	`db:"telepon"`
    Facebook		        string  `db:"facebook"`
    LinkFacebook            string  `db:"link_facebook"`
    Twitter		            string	`db:"twitter"`
    LinkTwitter             string  `db:"link_twitter"`
    Youtube                 string  `db:"youtube"`
    LinkYoutube             string  `db:"link_youtube"`
    Google                  string  `db:"google"`
    LinkGoogle              string  `db:"link_google"`
    Linkedin                string  `db:"linkedin"`
    LinkLinkedin            string  `db:"link_linkedin"`
    Instagram               string  `db:"instagram"`
    LinkInstagram           string  `db:"link_instagram"`
    Line				    string  `db:"line"`
    LinkLine                string  `db:"link_line"`
    CreatedAt		        int64	`db:"contact_created_at"`
    UpdatedAt		        int64	`db:"contact_updated_at"`		
}

func CreateDefaultContact() Contact {
    contact_dummy := Contact {
        ContactID          : 1,
        Deskripsi          : "Bila Anda membutuhkan informasi mengenai program-program wirausaha mahasiswa dari ITB Career Center, hubungi kami melalui telepon, email, atau akun media sosial kami.",   
        AlamatITB          : "GKU Timur Lt. Dasar, Kampus ITB",
        Alamat             : "Jl. Ganesa 10, Bandung, 40132",  
        Email              : "career@itb.ac.id ",  
        Telepon            : "(62-22) 250 9177",  
        Facebook           : "itb.career.center",  
        LinkFacebook       : "https://www.facebook.com/itb.career.center/",
        Twitter            : "@ITBCareerCenter",  
        LinkTwitter        : "https://twitter.com/ITBCareerCenter",
        Youtube            : "ITBCareerCenter",  
        LinkYoutube        : "https://www.youtube.com/user/ITBCareerCenter",
        Google             : "ITB CareerCenter",   
        LinkGoogle         : "https://plus.google.com/u/0/106590411857705007612/posts",
        Linkedin           : "ITB Career Center",
        LinkLinkedin       : "https://www.linkedin.com/company-beta/4788319/",
        Instagram          : "@itbcareercenter",
        LinkInstagram      : "https://www.instagram.com/itbcareercenter/",
        Line               : "@itbcareercenter",       
        LinkLine           : "https://line.me/ti/p/%40yyc8475l",
        CreatedAt          : time.Now().UnixNano(), 
        UpdatedAt          : time.Now().UnixNano(),     
    }
    return contact_dummy
}