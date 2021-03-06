package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type LoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Token    string `json:"token"`
}
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
}

type UserPoin struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Alamat   string `json:"alamat"`
	Poin     int    `json:"poin"`
}

type ManageCustomer struct {
	IDCustomer int    `json:"id"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	No_hp      string `json:"no_hp" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

type ManageCustomerPoint struct {
	IDCustomer   int    `json:"id"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	No_hp        string `json:"no_hp" validate:"required"`
	Poin_account int    `json:"poin_account" validate:"required"`
}

type UpdateCustomerPoint struct {
	IDCustomer   int `json:"id"`
	Poin_account int `json:"poin_account" validate:"required"`
}

type Kota struct {
	ID                 int    `json:"id"`
	Rajaongkir_city_id int    `json:"rajaongkir_city_id" validate:"required"`
	Kota_Nama          string `json:"kota_nama" validate:"required"`
	Postal_code        int    `json:"postal_code" validate:"required"`
	Tipe               string `json:"tipe" validate:"required"`
	Province_ID        int    `json:"province_id" validate:"required"`
}

type TestAdmin struct {
	ID       int    `gorm:"primaryKey"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Auth struct {
	Token string
}

type Claims struct {
	Username string
	jwt.StandardClaims
}

type GetCityById struct {
	Nama_kota string `json:"nama_kota"`
	Tipe_kota string `json:"tipe_kota"`
}

type Dashboard struct {
	Transaction_id     string `json:"transaction_id"`
	Nomor              string `json:"nomor"`
	Customer_id        int    `json:"customer_id"`
	Keterangan         string `json:"keterangan"`
	Status_transaction string `json:"status_transaction"`
}
type ApproveTransaction struct {
	Status_transaction string `json:"status_transaction"`
}

type CustomerHistory struct {
	Customer_id int       `json:"customer_id"`
	Keterangan  string    `json:"keterangan"`
	Nomor       string    `json:"nomor"`
	Tanggal     time.Time `json:"tanggal"`
	Status      string    `json:"status"`
	Poin_redeem int       `json:"poin_redeem"`
}
