package admin

import (
	"api-redeem-point/business/admin"
	adminBusiness "api-redeem-point/business/admin"
	"net/http"
	"strconv"

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

func (Controller *Controller) GetAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	admin, err := Controller.service.GetAdminByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success get data admin",
		"data":     admin,
	})
}

func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := admin.Admin{}
	c.Bind(&admin)
	admins, err := Controller.service.CreateAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":     201,
		"messages": "success create data",
		"data":     admins,
	})
}

func (Controller *Controller) GetToken(c echo.Context) error {
	var request adminBusiness.Admin

	c.Bind(&request)
	token, err := Controller.service.GetToken(&request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "success login",
		"token":   token,
	})
}

func (Controller *Controller) DeleteAdmin(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := Controller.service.DeleteAdmin(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success delete admin",
		"data id ": id,
	})
}

func (Controller *Controller) UpdateAdmin(c echo.Context) error {
	var admin *admin.Admin
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&admin)
	admin, err := Controller.service.UpdateAdmin(id, admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":     200,
		"messages": "success update data admin",
		"data":     admin,
	})
}
