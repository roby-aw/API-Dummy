package admin

import (
	adminBusiness "api-redeem-point/business/admin"
	"fmt"
	"net/http"

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
	result := adminBusiness.Dashboard{
		User:  15,
		Stock: 120,
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
// @Summary Manage Customer Point
// @description Manage Customer Point
// @tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} admin.User
// @Router /v1/admin/managecustomerpoint [get]
func (Controller *Controller) ManageCustomerPoint(c echo.Context) error {
	var err error
	user1 := adminBusiness.User{
		ID:       1,
		Email:    "testuser1@gmail.com",
		Password: "testpassworduser1",
		Alamat:   "testalamat",
		Poin:     420000,
	}
	user2 := adminBusiness.User{
		ID:       2,
		Email:    "testuser2@gmail.com",
		Password: "testpassworduser2",
		Alamat:   "testalamat",
		Poin:     503000,
	}
	user3 := adminBusiness.User{
		ID:       3,
		Email:    "testuser3@gmail.com",
		Password: "testpassworduser3",
		Alamat:   "testalamat",
		Poin:     520000,
	}
	user4 := adminBusiness.User{
		ID:       4,
		Email:    "testuser4@gmail.com",
		Password: "testpassworduser4",
		Alamat:   "testalamat",
		Poin:     560000,
	}
	var arr []adminBusiness.User
	arr = append(arr, user1)
	arr = append(arr, user2)
	arr = append(arr, user3)
	arr = append(arr, user4)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get manage point customer",
		"result":   arr,
	})
}
