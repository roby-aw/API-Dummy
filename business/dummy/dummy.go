package dummy

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name        string `json:"name"`
	Picture_url string `json:"picture_url"`
	City        string `json:"city"`
	Price       string `json:"price"`
	Open_time   string `json:"open_time"`
	Latitude    string `json:"latitude"`
	Longtitude  string `json:"Longtitude"`
	Rating      string `json:"rating"`
	Visited     int    `json:"visited"`
}

type Login struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}
type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
