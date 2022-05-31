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
// @Summary Register
// @description Register user
// @tags User
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
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

// Create godoc
// @Summary Order Emoney
// @description Emoney user
// @tags UserOrder
// @Accept json
// @Produce json
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
// @Success 200 {object} map[string]interface{}
// @Router /v1/order/paketdata [post]
func (Controller *Controller) OrderPaketData(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success register",
	})
}
