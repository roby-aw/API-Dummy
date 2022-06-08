package admin

import (
	"api-redeem-point/api/dummy"
	"api-redeem-point/business/admin"
	adminBusiness "api-redeem-point/business/admin"
	dummyBusiness "api-redeem-point/business/dummy"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) GetAdmins(c echo.Context) error {
	admins, err := Controller.service.GetAdmins()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get all admin",
		"data":     admins,
	})
}

// Create godoc
// @Summary Login
// @description Login Admin
// @tags Admin
// @Accept json
// @Produce json
// @Param LoginAdmin body admin.LoginAdmin true "Admin"
// @Success 200 {object} admin.LoginAdmin
// @Router /v1/admin/login [post]
func (Controller *Controller) LoginAdmin(c echo.Context) error {
	var err error
	result := adminBusiness.Admin{
		ID:       1,
		Name:     "testadmin",
		Email:    "testadmin@gmail.com",
		Password: "testpassword",
		Token:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTE2ODczODh9.dw2WuBIDJcb5dVT8iF_63POdhZFYOq4D1-kZTiCyo7c",
	}
	var req adminBusiness.Admin
	c.Bind(&req)
	if (req.Email != result.Email) || (req.Password != result.Password) {
		err = fmt.Errorf("Email atau password salah")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success login admin",
		"result":   result,
	})
}

// Create godoc
// @Summary Dashboard Admin
// @description Dashboard Admin
// @tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} admin.Dashboard
// @Router /v1/admin/dashboard [get]
func (Controller *Controller) Dashboard(c echo.Context) error {
	var err error
	var result []adminBusiness.Dashboard
	for _, v := range dummy.DetailTransaction {
		if v.Status_transaction == "PENDING" {
			var tmpDashboard adminBusiness.Dashboard
			tmpDashboard.Customer_id = v.Customer_id
			tmpDashboard.Transaction_id = v.Transaction_id
			tmpDashboard.Status_transaction = v.Status_transaction

			result = append(result, tmpDashboard)
		}
	}
	if len(result) == 0 {
		err = errors.New("Tidak ada pending data")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get dashboard",
		"result":   result,
	})
}

// Create godoc
// @Summary Approve Transaction
// @description Approve Transaction
// @tags Admin
// @Accept json
// @Produce json
// @Param transactionid path string true "transaction_id"
// @Success 200
// @Router /v1/admin/approve/{transactionid} [post]
func (Controller *Controller) ApproveTransaction(c echo.Context) error {
	var err error
	transactionid := c.Param("transactionid")
	var check dummyBusiness.DetailTransaction
	for _, v := range dummy.DetailTransaction {
		if v.Transaction_id == transactionid {
			check = v
			dummy.DetailTransaction[v.ID-1].Status_transaction = "COMPLETED"
		}
	}
	if check.Transaction_id == "" {
		err = errors.New("transaction id tidak ditemukan")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success approve transaction",
	})
}

// Create godoc
// @Summary Manage Customer Point
// @description Manage Customer Point
// @tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} admin.UserPoin
// @Router /v1/admin/customerpoint [get]
func (Controller *Controller) ManageCustomerPoint(c echo.Context) error {
	var err error
	var result []admin.ManageCustomerPoint
	var req admin.ManageCustomerPoint
	c.Bind(&req)
	for _, v := range dummy.Customer {
		var tmpCustomer admin.ManageCustomerPoint
		tmpCustomer.IDCustomer = v.ID
		tmpCustomer.Email = v.Email
		tmpCustomer.Name = v.Name
		tmpCustomer.No_hp = v.No_hp
		tmpCustomer.Poin_account = v.Poin

		result = append(result, tmpCustomer)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get manage point customer",
		"result":   result,
	})
}

// Create godoc
// @Summary Manage Customer
// @description Manage Customer
// @tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} admin.ManageCustomer
// @Router /v1/admin/customer [get]
func (Controller *Controller) ManageCustomer(c echo.Context) error {
	var err error
	var result []adminBusiness.ManageCustomer
	for _, v := range dummy.Customer {
		var data adminBusiness.ManageCustomer
		data.IDCustomer = v.ID
		data.Email = v.Email
		data.Name = v.Name
		data.Password = v.Password
		data.No_hp = v.No_hp
		result = append(result, data)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get manage customer",
		"result":   result,
	})
}

// Create godoc
// @Summary Customer history
// @description Customer history
// @tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} admin.CustomerHistory
// @Router /v1/admin/historycustomer [get]
func (Controller *Controller) CustomerHistory(c echo.Context) error {
	var err error
	var result []adminBusiness.CustomerHistory
	for _, v := range dummy.DetailTransaction {
		var tmpHistory adminBusiness.CustomerHistory
		tmpHistory.Customer_id = v.Customer_id
		tmpHistory.Keterangan = v.Keterangan
		tmpHistory.Nomor = v.Nomor
		tmpHistory.Tanggal = v.CreatedAt
		tmpHistory.Status = v.Status_transaction
		tmpHistory.Poin_redeem = v.Poin_redeem

		result = append(result, tmpHistory)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get history customer",
		"result":   result,
	})
}

// Create godoc
// @Summary Update customer data
// @description Update customer data
// @tags Admin
// @Accept json
// @Produce json
// @Param Update Customer body admin.ManageCustomer true "Customer"
// @Success 200 {object} admin.ManageCustomer
// @Router /v1/admin/customer [PUT]
func (Controller *Controller) UpdateCustomer(c echo.Context) error {
	var req adminBusiness.ManageCustomer
	var err error
	c.Bind(&req)
	err = validator.New().Struct(req)
	if req.IDCustomer > len(dummy.Customer) {
		err = errors.New("user tidak ada")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	dummy.Customer[req.IDCustomer-1].Name = req.Name
	dummy.Customer[req.IDCustomer-1].No_hp = req.No_hp
	dummy.Customer[req.IDCustomer-1].Email = req.Email
	dummy.Customer[req.IDCustomer-1].Password = req.Password
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update customer",
	})
}

// Create godoc
// @Summary Update customer point
// @description Update customer point
// @tags Admin
// @Accept json
// @Produce json
// @Param UpdateCustomerPoint body admin.UpdateCustomerPoint true "Customer Point"
// @Success 200
// @Router /v1/admin/customerpoint [PUT]
func (Controller *Controller) UpdateCustomerPoint(c echo.Context) error {
	var req adminBusiness.UpdateCustomerPoint
	var err error
	c.Bind(&req)
	err = validator.New().Struct(req)
	if req.IDCustomer > len(dummy.Customer) {
		err = errors.New("user tidak ada")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":     400,
			"messages": err.Error(),
		})
	}
	dummy.Customer[req.IDCustomer-1].Poin = dummy.Customer[req.IDCustomer-1].Poin + req.Poin_account
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update point customer",
	})
}
