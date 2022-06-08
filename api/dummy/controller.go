package dummy

import (
	"api-redeem-point/api/dummy/response"
	"api-redeem-point/business/dummy"
	dummyBussiness "api-redeem-point/business/dummy"
	"api-redeem-point/business/mitra"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service dummyBussiness.Service
}

func NewController(service dummyBussiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

var AccountMitra []mitra.Mitra
var Stockproduct []dummy.StockProduct
var Customer []dummy.Customer
var DetailTransaction []dummy.DetailTransaction

func InitiateDB() {
	product1 := dummy.StockProduct{
		ID:      1,
		Product: "Pulsa dan Paket Data",
		Stock:   800000,
	}
	product2 := dummy.StockProduct{
		ID:      2,
		Product: "EMoney dan Cashout",
		Stock:   800000,
	}
	Stockproduct = append(Stockproduct, product1)
	Stockproduct = append(Stockproduct, product2)

	Customer1 := dummy.Customer{
		ID:       1,
		Name:     "testname",
		Email:    "test@gmail.com",
		Password: "testpassword",
		No_hp:    "0856356214",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE2ODczODh9.dw2WuBIDJcb5dVT8iF_63POdhZFYOq4D1-kZTiCyo7c",
		Poin:     500000,
		Pin:      1234,
	}
	Customer = append(Customer, Customer1)

	transaction1 := dummy.DetailTransaction{
		ID:                 1,
		Customer_id:        1,
		Transaction_id:     "EM" + randomstring(),
		Jenis_transaction:  "Redeem Cashout/Emoney",
		Bank_Provider:      "BRI",
		AN_Rekening:        "Yudo",
		Nomor:              "2563532554",
		Poin_account:       120000,
		Poin_redeem:        50000,
		Amount:             50000,
		Keterangan:         "BRI - 50000",
		Status_transaction: "COMPLETED",
		Status_poin:        "OUT",
	}
	transaction2 := dummy.DetailTransaction{
		ID:                 2,
		Customer_id:        1,
		Transaction_id:     "P" + randomstring(),
		Jenis_transaction:  "Redeem Pulsa/PaketData",
		Bank_Provider:      "TELKOMSEL",
		Nomor:              "085756545325",
		Poin_account:       70000,
		Poin_redeem:        10000,
		Keterangan:         "TELKOMSEL - 10000",
		Status_transaction: "COMPLETED",
		Status_poin:        "OUT",
	}
	transaction3 := dummy.DetailTransaction{
		ID:                 3,
		Customer_id:        1,
		Transaction_id:     "P" + randomstring(),
		Jenis_transaction:  "Redeem Pulsa/PaketData",
		Bank_Provider:      "TELKOMSEL",
		Nomor:              "085756545325",
		Poin_account:       60000,
		Poin_redeem:        10000,
		Keterangan:         "TELKOMSEL - 10000",
		Status_transaction: "PENDING",
		Status_poin:        "OUT",
	}
	transaction4 := dummy.DetailTransaction{
		ID:                 4,
		Customer_id:        1,
		Transaction_id:     "P" + randomstring(),
		Jenis_transaction:  "Redeem Pulsa/PaketData",
		Bank_Provider:      "TELKOMSEL",
		Nomor:              "085756545325",
		Poin_account:       60000,
		Poin_redeem:        10000,
		Keterangan:         "TELKOMSEL - 3GB",
		Status_transaction: "PENDING",
		Status_poin:        "OUT",
	}
	DetailTransaction = append(DetailTransaction, transaction1)
	DetailTransaction = append(DetailTransaction, transaction2)
	DetailTransaction = append(DetailTransaction, transaction3)
	DetailTransaction = append(DetailTransaction, transaction4)

	mitra1 := mitra.Mitra{
		ID:        1,
		Nama_toko: "Mitra1",
		Email:     "mitra1@gmail.com",
		Password:  "mitrapassword",
		Alamat:    "Jl. Mitra",
	}

	mitra2 := mitra.Mitra{
		ID:        1,
		Nama_toko: "Mitra2",
		Email:     "mitra2@gmail.com",
		Password:  "mitrapassword",
		Alamat:    "Jl. Mitra",
	}

	AccountMitra = append(AccountMitra, mitra1)
	AccountMitra = append(AccountMitra, mitra2)
}

// Create godoc
// @Summary Login
// @description Login Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Customer body dummy.AuthLogin true "Customer"
// @Success 200 {object} dummy.Customer
// @Router /v1/login [post]
func (Controller *Controller) Login(c echo.Context) error {
	var req dummyBussiness.AuthLogin
	var tmpCustomer dummy.Customer
	var err error
	c.Bind(&req)
	err = validator.New().Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	for _, v := range Customer {
		if v.Email == req.Email && v.Password == req.Password {
			tmpCustomer = v
		}
	}
	if tmpCustomer.Email == "" || tmpCustomer.Password == "" {
		err = errors.New("Email Atau Password salah")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"result":   tmpCustomer,
	})
}

// Create godoc
// @Summary History
// @description History Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param idcustomer path int true "id customer"
// @Success 200 {object} dummy.History
// @Router /v1/history/{idcustomer} [get]
func (Controller *Controller) History(c echo.Context) error {
	var err error
	var History []dummy.History
	iduser, _ := strconv.Atoi(c.Param("idcustomer"))
	for _, v := range DetailTransaction {
		if v.Customer_id == iduser {
			if v.Status_poin == "OUT" {
				var dethistory dummy.History
				dethistory.ID = v.ID
				dethistory.Transaction_id = v.Transaction_id
				dethistory.Keterangan = v.Keterangan
				dethistory.Tanggal = time.Now()
				dethistory.Status = v.Status_transaction

				History = append(History, dethistory)
			}
		}
	}
	if len(History) == 0 {
		err = errors.New("ID User tidak ada history")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get history",
		"result":   History,
	})
}

// Create godoc
// @Summary Detail History/transaction
// @description History/transaction Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param transactionid path string true "id detail history"
// @Success 200 {object} dummy.DetailTransactionCustomer
// @Router /v1/detailhistory/{transactionid} [get]
func (Controller *Controller) DetailTransaction(c echo.Context) error {
	var err error
	var tmpHistory dummy.DetailTransactionCustomer
	id := c.Param("transactionid")
	for _, v := range DetailTransaction {
		if v.Transaction_id == id {
			tmpHistory.Transaction_id = v.Transaction_id
			tmpHistory.Jenis_transaction = v.Jenis_transaction
			tmpHistory.CreatedAt = v.CreatedAt
			tmpHistory.Bank_Provider = v.Bank_Provider
			tmpHistory.Nomor = v.Nomor
			tmpHistory.Poin_account = v.Poin_account
			tmpHistory.Poin_redeem = v.Poin_redeem
		}
	}
	if tmpHistory.Transaction_id == "" {
		err = errors.New("Data tidak ditemukan")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get detail transaction",
		"result":   tmpHistory,
	})
}

// Create godoc
// @Summary Register
// @description Register Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param RegisterCustomer body dummy.Register true "Register"
// @Success 200 {object} dummy.Register
// @Router /v1/register [post]
func (Controller *Controller) Register(c echo.Context) error {
	var req dummyBussiness.Register
	var tmpCustomer dummy.Customer
	var err error
	c.Bind(&req)
	err = validator.New().Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	for _, v := range Customer {
		if v.Email == req.Email {
			err = errors.New("Email sudah digunakan")
		}
	}
	tmpCustomer.ID = len(Customer) + 1
	tmpCustomer.Name = req.Name
	tmpCustomer.Email = req.Email
	tmpCustomer.Password = req.Password
	tmpCustomer.No_hp = req.No_hp
	Customer = append(Customer, tmpCustomer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
		"result":   req,
	})
}

// Create godoc
// @Summary Order Emoney/Cashout
// @description Emoney user
// @tags CustomerOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/emoney [post]
func (Controller *Controller) OrderEmoney(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order Pulsa
// @description Pulsa user
// @tags CustomerOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param OrderPulsa body dummy.OrderPulsa true "OrderPulsa"
// @Success 200 {object} dummy.OrderPulsa
// @Router /v1/order/pulsa [post]
func (Controller *Controller) OrderPulsa(c echo.Context) error {
	var err error
	var req dummy.OrderPulsa
	c.Bind(&req)
	err = validator.New().Struct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	transaction := dummyBussiness.DetailTransaction{
		ID:                 len(DetailTransaction) + 1,
		Customer_id:        req.Customer_id,
		Transaction_id:     "P" + randomstring(),
		Bank_Provider:      req.Bank_Provider,
		Nomor:              req.Nomor,
		Poin_account:       req.Poin_account,
		Poin_redeem:        req.Poin_redeem,
		Keterangan:         req.Keterangan,
		Status_transaction: "PENDING",
		Status_poin:        "OUT",
	}
	DetailTransaction = append(DetailTransaction, transaction)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success order pulsa",
	})
}

// Create godoc
// @Summary Order PaketData
// @description PaketData user
// @tags CustomerOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/paketdata [post]
func (Controller *Controller) OrderPaketData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Update Customer
// @description Update Customer
// @tags Customer
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param Update Customer body dummy.UpdateCustomer true "UpdateCustomer"
// @Success 200 {object} dummy.UpdateCustomer
// @Router /v1/customer [put]
func (Controller *Controller) UpdateCustomer(c echo.Context) error {
	var req dummyBussiness.UpdateCustomer
	var err error
	c.Bind(&req)
	err = validator.New().Struct(req)
	if req.IDCustomer > len(Customer) {
		err = errors.New("user tidak ada")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	Customer[req.IDCustomer-1].Name = req.Name
	Customer[req.IDCustomer-1].No_hp = req.No_hp
	Customer[req.IDCustomer-1].Password = req.Password
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update customer",
	})
}

func (Controller *Controller) CallbackXendit(c echo.Context) error {
	req := c.Request()
	decoder := json.NewDecoder(req.Body)
	disbursermentData := dummyBussiness.Disbursement{}
	err := decoder.Decode(&disbursermentData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	defer req.Body.Close()
	disbursement, _ := json.Marshal(disbursermentData)
	var resbank dummyBussiness.Disbursement
	json.Unmarshal(disbursement, &resbank)
	responseWriter := c.Response().Writer
	responseWriter.Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	responseWriter.WriteHeader(200)
	fmt.Fprintf(responseWriter, "%s", disbursement)
	fmt.Println(resbank)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"res": responseWriter,
		"dis": disbursement,
	})
}

// Create godoc
// @Summary Get Stock Product
// @description Get Stock Product
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} response.ResponseGetProduct
// @Router /v1/stockproduct [get]
func (Controller *Controller) StockProduct(c echo.Context) error {
	var err error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.ResponseGetProduct{
		Code:    200,
		Message: "Success get stock product",
		Result:  Stockproduct,
	})
}

// Create godoc
// @Summary Update Stock Product
// @description Update Stock Product
// @tags Product
// @Accept json
// @Produce json
// @Param id path int true "id product"
// @Param Update Stock Product body dummy.InputStockProduct true "Update Stock Product"
// @Success 200 {object} response.ResponseGetProduct
// @Router /v1/stockproduct/{id} [put]
func (Controller *Controller) ManageStockProduct(c echo.Context) error {
	var stock *dummy.InputStockProduct
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&stock)
	id = id - 1
	Stockproduct[id].Stock = Stockproduct[id].Stock + stock.Stock
	product := Stockproduct[id]
	var err error
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.ResponsePutProduct{
		Code:    200,
		Message: "Success update stock product",
		Result:  product,
	})
}

// Create godoc
// @Summary Login Mitra
// @description Login Mitra
// @tags Mitra
// @Accept json
// @Produce json
// @Param mitra body mitra.AuthMitra true "mitra"
// @Success 200 {object} mitra.Mitra
// @Router /v1/mitra/login [post]
func (Controller *Controller) LoginMitra(c echo.Context) error {
	var req mitra.AuthMitra
	var tmpMitra mitra.Mitra
	var err error
	c.Bind(&req)
	for _, v := range AccountMitra {
		if v.Email == req.Email && v.Password == req.Password {
			tmpMitra = v
		}
	}
	if tmpMitra.Email == "" && tmpMitra.Password == "" {
		err = errors.New("Email Atau Password salah")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login",
		"result":   tmpMitra,
	})
}

// Create godoc
// @Summary Register Mitra
// @description Register Mitra
// @tags Mitra
// @Accept json
// @Produce json
// @Param RegisterMitra body mitra.MitraRegister true "MitraRegister"
// @Success 200 {object} mitra.MitraRegister
// @Router /v1/mitra/register [post]
func (Controller *Controller) RegisterMitra(c echo.Context) error {
	var req mitra.MitraRegister
	var tmpMitra mitra.Mitra
	var err error
	c.Bind(&req)
	for _, v := range AccountMitra {
		if v.Email == req.Email {
			err = errors.New("Email sudah digunakan")
		}
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}

	tmpMitra.ID = len(AccountMitra) + 1
	tmpMitra.Nama_toko = req.Nama_toko
	tmpMitra.Email = req.Email
	tmpMitra.Password = req.Password
	tmpMitra.Alamat = req.Alamat
	AccountMitra = append(AccountMitra, tmpMitra)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
		"result":   req,
	})
}

// Create godoc
// @Summary Input Poin Mitra
// @description Input Poin Mitra
// @tags Mitra
// @Accept json
// @Produce json
// @Param InputPoin body mitra.InputPoinMitra true "InputPoinMitra"
// @Success 200
// @Router /v1/mitra/inputpoin [post]
func (Controller *Controller) InputPoin(c echo.Context) error {
	var req mitra.InputPoinMitra
	var data dummy.DetailTransaction
	var err error
	c.Bind(&req)
	if req.IDCustomer > len(Customer) {
		err = errors.New("user tidak ada")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	var i int
	price := req.Amount
	for i = 0; price > 10000; i = i + 100 {
		price = price - 10000
	}

	data.ID = len(DetailTransaction) + 1
	data.Customer_id = req.IDCustomer
	data.Mitra_id = req.IDMitra
	data.Amount = req.Amount
	data.Poin_redeem = i
	data.Status_poin = "IN"
	data.Poin_account = Customer[req.IDCustomer-1].Poin

	acc := Customer[req.IDCustomer-1].Poin
	Customer[req.IDCustomer-1].Poin = acc + i

	DetailTransaction = append(DetailTransaction, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success add poin",
	})
}

// Create godoc
// @Summary History Mitra
// @description History Mitra
// @tags Mitra
// @Accept json
// @Produce json
// @Success 200
// @Router /v1/mitra/history [get]
func (Controller *Controller) HistoryMitra(c echo.Context) error {
	var err error
	var History []*mitra.HistoryMitra
	for _, v := range DetailTransaction {
		if v.Status_poin == "IN" {
			var data mitra.HistoryMitra
			data.Customer_id = v.Customer_id
			data.Mitra_id = v.Mitra_id
			data.Amount = v.Amount
			data.Poin_redeem = v.Poin_redeem
			History = append(History, &data)
		}
	}
	if History == nil {
		err = errors.New("data tidak ada")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get mitra",
		"result":   History,
	})
}

func randomstring() string {
	rand.Seed(time.Now().UTC().UnixNano())
	random := rand.Int()
	angka := []rune(fmt.Sprintf("%d", random))
	b := make([]rune, 8)
	for i := range b {
		b[i] = angka[rand.Intn(len(angka))]
	}
	hasil := string(b)
	inthasil, _ := strconv.Atoi(hasil)
	strhasil := strconv.Itoa(inthasil)
	return strhasil
}
