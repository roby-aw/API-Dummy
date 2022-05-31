package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/dummy"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller *admin.Controller
	DummyController  *dummy.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// user
	c := e.Group("/v1")
	c.GET("/foods/:name", controller.DummyController.GetFoodByName)
	c.POST("/login", controller.DummyController.Login)
	c.GET("/history/:iduser", controller.DummyController.History)
	c.GET("/detailhistory/:id", controller.DummyController.DetailTransaction)
	// product
	// f := c.Group("/product")
	// f.GET("/cashout", controller.DummyController.ProductCashout)
	// f.GET("/emoney", controller.DummyController.ProductEmoney)
	// f.GET("/pulsa", controller.DummyController.ProductPulsa)
	// f.GET("/paketdata", controller.DummyController.ProductPaketData)
	//admin

}
