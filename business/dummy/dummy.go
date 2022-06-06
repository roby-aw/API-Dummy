package dummy

import (
	"time"

	"gorm.io/gorm"
)

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

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
}

type Customer struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
	Poin     int    `json:"poin"`
	Pin      int    `json:"pin"`
}
type AuthLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type History struct {
	ID             int       `json:"id"`
	Transaction_id string    `json:"transaction_id"`
	Keterangan     string    `json:"keterangan"`
	Tanggal        time.Time `json:"tanggal"`
	Status         string    `json:"status"`
}

type DetailTransaction struct {
	ID                 int       `json:"id"`
	Customer_id        int       `json:"customer_id"`
	Mitra_id           int       `json:"mitra_id"`
	Transaction_id     string    `json:"transaction_id"`
	CreatedAt          time.Time `json:"createdat"`
	Jenis_transaction  string    `json:"jenis_transaction"`
	Bank_Provider      string    `json:"bank_provider"`
	AN_Rekening        string    `json:"an_rekening"`
	Nomor              string    `json:"nomor"`
	Poin_account       int       `json:"poin_account"`
	Poin_redeem        int       `json:"poin_redeem"`
	Amount             int       `json:"amount"`
	Keterangan         string    `json:"keterangan"`
	Status_transaction string    `json:"status_transaction"`
	Status_poin        string    `json:"status_poin"`
}

type ProductCashout struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}
type ProductEmoney struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}
type ProductPulsa struct {
	ID    int `json:"id"`
	Harga int `json:"harga"`
	Poin  int `json:"poin"`
}

type ProductPaketData struct {
	ID    int    `json:"id"`
	Nama  string `json:"Internet"`
	Kuota string `json:"kuota"`
	Harga int    `json:"harga"`
	Poin  int    `json:"poin"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	No_hp    string `json:"no_hp"`
	Password string `json:"password"`
	Pin      int    `json:"pin"`
}

type Bank struct {
	BankCode    string `json:"bankcode"`
	No_rekening string `json:"no_rekening"`
	Amount      int    `json:"amount"`
}

type Disbursement struct {
	UserID                  string `json:"user_id"`
	ExternalID              string `json:"external_id"`
	Amount                  int    `json:"amount"`
	BankCode                string `json:"bank_code"`
	AccountHolderName       string `json:"account_holder_name"`
	DisbursementDescription string `json:"disbursement_description"`
	Status                  string `json:"status"`
	ID                      string `json:"id"`
}

type StockProduct struct {
	ID      int    `gorm:"primarykey"`
	Product string `json:"product"`
	Stock   int    `json:"stock"`
}

type InputStockProduct struct {
	Stock int `json:"stock"`
}

type InputPin struct {
	Pin int `json:"pin"`
}

type TransactionBank struct {
	An_bank     string `json:"an_bank"`
	No_rekening string `json:"no_rekening"`
}

type InputPoinMitra struct {
	IDCustomer   int `json:"idcustomer"`
	Poin_account int `json:"poin_account"`
	Amount       int `json:"amount"`
}
