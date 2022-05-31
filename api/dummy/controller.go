package dummy

import (
	dummyBussiness "api-redeem-point/business/dummy"
	"fmt"
	"net/http"

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

func (Controller *Controller) Login(c echo.Context) error {
	result := &dummyBussiness.Login{
		ID:       1,
		Email:    "test@gmail.com",
		Password: "testpassword",
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
