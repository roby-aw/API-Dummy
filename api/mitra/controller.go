package mitra

import (
	mitraBusiness "api-redeem-point/business/mitra"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service mitraBusiness.Service
}

func NewController(service mitraBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

var AccountMitra []mitraBusiness.Mitra

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
	var req mitraBusiness.AuthMitra
	var tmpMitra mitraBusiness.Mitra
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
func (Controller *Controller) Register(c echo.Context) error {
	var req mitraBusiness.MitraRegister
	var tmpMitra mitraBusiness.Mitra
	var err error
	c.Bind(&req)
	for _, v := range AccountMitra {
		if v.Email == req.Email {
			err = errors.New("Email sudah digunakan")
		}
	}
	tmpMitra.ID = len(AccountMitra) + 1
	tmpMitra.Nama_toko = req.Nama_toko
	tmpMitra.Email = req.Email
	tmpMitra.Password = req.Password
	tmpMitra.Alamat = req.Alamat
	AccountMitra = append(AccountMitra, tmpMitra)
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
