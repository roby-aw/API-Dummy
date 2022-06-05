package dummy

import (
	"api-redeem-point/api/dummy/response"
	"api-redeem-point/business/dummy"
	dummyBussiness "api-redeem-point/business/dummy"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
		ID:                1,
		Customer_id:       1,
		Transaction_id:    "EM25434353",
		Jenis_transaction: "Redeem Cashout/Emoney",
		An_bank:           "BRI",
		No_rekening:       "2563532554",
		Poin_account:      52000,
		Poin_redeem:       50000,
		Amount:            50000,
		Keterangan:        "Redeem Cashout - 10000",
		Status:            "COMPLETED",
	}
	transaction2 := dummy.DetailTransaction{
		ID:                2,
		Customer_id:       1,
		Transaction_id:    "P45574568",
		Jenis_transaction: "Redeem Pulsa/PaketData",
		Poin_account:      70000,
		Poin_redeem:       10000,
		Keterangan:        "Redeem Pulsa - 10000",
		Status:            "COMPLETED",
	}
	DetailTransaction = append(DetailTransaction, transaction1)
	DetailTransaction = append(DetailTransaction, transaction2)
}

// Create godoc
// @Summary Login
// @description Login user
// @tags User
// @Accept json
// @Produce json
// @Param user body dummy.AuthLogin true "user"
// @Success 200 {object} dummy.Customer
// @Router /v1/login [post]
func (Controller *Controller) Login(c echo.Context) error {
	var req dummyBussiness.AuthLogin
	var tmpCustomer dummy.Customer
	var err error
	c.Bind(&req)
	for _, v := range Customer {
		if v.Email == req.Email && v.Password == req.Password {
			tmpCustomer = v
		}
	}
	if tmpCustomer.Email == "" && tmpCustomer.Password == "" {
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
// @description History User
// @tags User
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
			var dethistory dummy.History
			dethistory.ID = v.ID
			dethistory.Keterangan = v.Keterangan
			dethistory.Tanggal = v.Tanggal
			dethistory.Status = v.Status

			History = append(History, dethistory)
		}
	}
	if History[0].Keterangan == "" {
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
// @description History/transaction User
// @tags User
// @Accept json
// @Produce json
// @Param id path int true "id detail history"
// @Success 200 {object} dummy.DetailTransaction
// @Router /v1/detailhistory/{id} [get]
func (Controller *Controller) DetailTransaction(c echo.Context) error {
	var err error
	detailtransaction := dummyBussiness.DetailTransaction{
		ID:                1,
		Jenis_transaction: "Redeem Cashout",
		An_bank:           "BNI",
		Poin_account:      500000,
		Poin_redeem:       100000,
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if id != 1 {
		err = fmt.Errorf("ID tidak ditemukan")
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
		"result":   detailtransaction,
	})
}

// Create godoc
// @Summary Register
// @description Register user
// @tags User
// @Accept json
// @Produce json
// @Param RegisterUser body dummy.Register true "Register"
// @Success 200 {object} dummy.Register
// @Router /v1/register [post]
func (Controller *Controller) Register(c echo.Context) error {
	var req dummyBussiness.Register
	var tmpCustomer dummy.Customer
	var err error
	c.Bind(&req)
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
// @Summary Order Cashout
// @description Order Cashout
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param InputDataCashout body dummy.Bank true "inputdatacashout"
// @Success 200 {object} dummy.Bank
// @Router /v1/order/cashout [post]
func (Controller *Controller) OrderCashout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order Emoney
// @description Emoney user
// @tags UserOrder
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
// @tags UserOrder
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/pulsa [post]
func (Controller *Controller) OrderPulsa(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order PaketData
// @description PaketData user
// @tags UserOrder
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
// @Summary UpdateUser
// @description UpdateUser
// @tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} dummy.Register
// @Router /v1/account [put]
func (Controller *Controller) UpdateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
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
