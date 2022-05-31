package dummy

import (
	dummyBussiness "api-redeem-point/business/dummy"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (Controller *Controller) GetFoodByName(c echo.Context) error {
	name := c.Param("name")
	fmt.Println(name)
	foods, err := Controller.service.GetFoodByName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get foods",
		"result":   foods,
	})
}

// Create godoc
// @Summary Login
// @description Login user
// @tags User
// @Accept json
// @Produce json
// @Param user body dummy.AuthLogin true "user"
// @Success 200 {object} dummy.Login
// @Router /v1/login [post]
func (Controller *Controller) Login(c echo.Context) error {
	result := &dummyBussiness.Login{
		ID:       1,
		Email:    "test@gmail.com",
		Password: "testpassword",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE2ODczODh9.dw2WuBIDJcb5dVT8iF_63POdhZFYOq4D1-kZTiCyo7c",
		Poin:     500000,
		Pin:      1234,
	}
	var req dummyBussiness.AuthLogin
	var err error
	c.Bind(&req)
	if (req.Email != result.Email) || (req.Password != result.Password) {
		err = fmt.Errorf("Email atau password salah")
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
		"result":   result,
	})
}

// Create godoc
// @Summary History
// @description History User
// @tags User
// @Accept json
// @Produce json
// @Param iduser path int true "id user"
// @Success 200 {object} dummy.History
// @Router /v1/history/{iduser} [get]
func (Controller *Controller) History(c echo.Context) error {
	History1 := &dummyBussiness.History{
		ID:             1,
		Tipe_transaksi: "Redeem CashOut",
		Tanggal:        time.Date(2022, 5, 16, 156, 24, 34, 534, time.UTC),
		Status:         "Sukses",
	}
	History2 := &dummyBussiness.History{
		ID:             5,
		Tipe_transaksi: "Redeem paket data",
		Tanggal:        time.Date(2022, 5, 17, 156, 24, 34, 534, time.UTC),
		Status:         "Sukses",
	}
	History3 := &dummyBussiness.History{
		ID:             7,
		Tipe_transaksi: "Redeem CashOut",
		Tanggal:        time.Date(2022, 5, 18, 156, 24, 34, 534, time.UTC),
		Status:         "Pending",
	}
	var arr []dummyBussiness.History
	arr = append(arr, *History1)
	arr = append(arr, *History2)
	arr = append(arr, *History3)
	var err error
	iduser, _ := strconv.Atoi(c.Param("iduser"))
	if iduser != 1 {
		err = fmt.Errorf("iduser tidak ditemukan")
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
		"result":   arr,
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
		Nama_bank:         "BNI",
		No_rekening:       12354665,
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
// @Summary Product CashOut
// @description Product CashOut user
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []dummy.ProductCashout
// @Router /v1/product/cashout [get]
func (Controller *Controller) ProductCashout(c echo.Context) error {
	var err error
	product1 := dummyBussiness.ProductCashout{
		ID:    1,
		Harga: 50000,
		Poin:  50000,
	}
	product2 := dummyBussiness.ProductCashout{
		ID:    2,
		Harga: 100000,
		Poin:  100000,
	}
	product3 := dummyBussiness.ProductCashout{
		ID:    3,
		Harga: 100000,
		Poin:  100000,
	}
	product4 := dummyBussiness.ProductCashout{
		ID:    4,
		Harga: 150000,
		Poin:  150000,
	}
	product5 := dummyBussiness.ProductCashout{
		ID:    5,
		Harga: 200000,
		Poin:  200000,
	}
	product6 := dummyBussiness.ProductCashout{
		ID:    6,
		Harga: 250000,
		Poin:  250000,
	}
	product7 := dummyBussiness.ProductCashout{
		ID:    7,
		Harga: 300000,
		Poin:  300000,
	}
	product8 := dummyBussiness.ProductCashout{
		ID:    8,
		Harga: 400000,
		Poin:  400000,
	}
	var arr []dummyBussiness.ProductCashout
	arr = append(arr, product1)
	arr = append(arr, product2)
	arr = append(arr, product3)
	arr = append(arr, product4)
	arr = append(arr, product5)
	arr = append(arr, product6)
	arr = append(arr, product7)
	arr = append(arr, product8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get product cashout",
		"result":   arr,
	})
}

// Create godoc
// @Summary Product EMoney
// @description Product EMoney user
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []dummy.ProductEmoney
// @Router /v1/product/emoney [get]
func (Controller *Controller) ProductEmoney(c echo.Context) error {
	var err error
	product1 := dummyBussiness.ProductEmoney{
		ID:    1,
		Harga: 50000,
		Poin:  50000,
	}
	product2 := dummyBussiness.ProductEmoney{
		ID:    2,
		Harga: 100000,
		Poin:  100000,
	}
	product3 := dummyBussiness.ProductEmoney{
		ID:    3,
		Harga: 150000,
		Poin:  150000,
	}
	product4 := dummyBussiness.ProductEmoney{
		ID:    4,
		Harga: 200000,
		Poin:  200000,
	}
	product5 := dummyBussiness.ProductEmoney{
		ID:    5,
		Harga: 250000,
		Poin:  250000,
	}
	product6 := dummyBussiness.ProductEmoney{
		ID:    6,
		Harga: 300000,
		Poin:  300000,
	}
	product7 := dummyBussiness.ProductEmoney{
		ID:    7,
		Harga: 350000,
		Poin:  350000,
	}
	product8 := dummyBussiness.ProductEmoney{
		ID:    8,
		Harga: 400000,
		Poin:  400000,
	}
	var arr []dummyBussiness.ProductEmoney
	arr = append(arr, product1)
	arr = append(arr, product2)
	arr = append(arr, product3)
	arr = append(arr, product4)
	arr = append(arr, product5)
	arr = append(arr, product6)
	arr = append(arr, product7)
	arr = append(arr, product8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get product emoney",
		"result":   arr,
	})
}

// Create godoc
// @Summary Product Pulsa
// @description Product Pulsa user
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []dummy.ProductPulsa
// @Router /v1/product/pulsa [get]
func (Controller *Controller) ProductPulsa(c echo.Context) error {
	var err error
	product1 := dummyBussiness.ProductPulsa{
		ID:    1,
		Harga: 10000,
		Poin:  10000,
	}
	product2 := dummyBussiness.ProductPulsa{
		ID:    2,
		Harga: 15000,
		Poin:  15000,
	}
	product3 := dummyBussiness.ProductPulsa{
		ID:    3,
		Harga: 20000,
		Poin:  20000,
	}
	product4 := dummyBussiness.ProductPulsa{
		ID:    4,
		Harga: 25000,
		Poin:  25000,
	}
	product5 := dummyBussiness.ProductPulsa{
		ID:    5,
		Harga: 30000,
		Poin:  30000,
	}
	product6 := dummyBussiness.ProductPulsa{
		ID:    6,
		Harga: 50000,
		Poin:  50000,
	}
	product7 := dummyBussiness.ProductPulsa{
		ID:    7,
		Harga: 75000,
		Poin:  75000,
	}
	product8 := dummyBussiness.ProductPulsa{
		ID:    8,
		Harga: 100000,
		Poin:  100000,
	}
	var arr []dummyBussiness.ProductPulsa
	arr = append(arr, product1)
	arr = append(arr, product2)
	arr = append(arr, product3)
	arr = append(arr, product4)
	arr = append(arr, product5)
	arr = append(arr, product6)
	arr = append(arr, product7)
	arr = append(arr, product8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get product pulsa",
		"result":   arr,
	})
}

// Create godoc
// @Summary Product PaketData
// @description Product PaketData user
// @tags Product
// @Accept json
// @Produce json
// @Success 200 {object} []dummy.ProductPaketData
// @Router /v1/product/paketdata [get]
func (Controller *Controller) ProductPaketData(c echo.Context) error {
	var err error
	product1 := dummyBussiness.ProductPaketData{
		ID:    1,
		Nama:  "Internet",
		Kuota: "1 GB",
		Harga: 10000,
		Poin:  10000,
	}
	product2 := dummyBussiness.ProductPaketData{
		ID:    2,
		Nama:  "Internet",
		Kuota: "2 GB",
		Harga: 15000,
		Poin:  15000,
	}
	product3 := dummyBussiness.ProductPaketData{
		ID:    3,
		Nama:  "Internet",
		Kuota: "5 GB",
		Harga: 20000,
		Poin:  20000,
	}
	product4 := dummyBussiness.ProductPaketData{
		ID:    4,
		Nama:  "Internet",
		Kuota: "8 GB",
		Harga: 25000,
		Poin:  25000,
	}
	product5 := dummyBussiness.ProductPaketData{
		ID:    5,
		Nama:  "Internet",
		Kuota: "10 GB",
		Harga: 30000,
		Poin:  30000,
	}
	product6 := dummyBussiness.ProductPaketData{
		ID:    6,
		Nama:  "Internet",
		Kuota: "12 GB",
		Harga: 50000,
		Poin:  50000,
	}
	product7 := dummyBussiness.ProductPaketData{
		ID:    7,
		Nama:  "Internet",
		Kuota: "14 GB",
		Harga: 75000,
		Poin:  75000,
	}
	product8 := dummyBussiness.ProductPaketData{
		ID:    8,
		Nama:  "Internet",
		Kuota: "16 GB",
		Harga: 100000,
		Poin:  100000,
	}
	var arr []dummyBussiness.ProductPaketData
	arr = append(arr, product1)
	arr = append(arr, product2)
	arr = append(arr, product3)
	arr = append(arr, product4)
	arr = append(arr, product5)
	arr = append(arr, product6)
	arr = append(arr, product7)
	arr = append(arr, product8)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get product paket data",
		"result":   arr,
	})
}

// Create godoc
// @Summary Register
// @description Register user
// @tags User
// @Accept json
// @Produce json
// @Router /v1/register [post]
func (Controller *Controller) Register(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}

// Create godoc
// @Summary Order Cashout
// @description Register user
// @tags UserOrder
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/cashout [post]
func (Controller *Controller) OrderCashout(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}
