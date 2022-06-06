package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/dummy"
	"api-redeem-point/api/mitra"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller *admin.Controller
	DummyController  *dummy.Controller
	MitraController  *mitra.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// user
	c := e.Group("/v1")
	c.POST("/login", controller.DummyController.Login)
	c.POST("/register", controller.DummyController.Register)
	c.GET("/history/:idcustomer", controller.DummyController.History)
	c.GET("/detailhistory/:id", controller.DummyController.DetailTransaction)
	c.GET("/stockproduct", controller.DummyController.StockProduct)
	c.PUT("/stockproduct/:id", controller.DummyController.ManageStockProduct)
	// product
	// f := c.Group("/product")
	// f.GET("/cashout", controller.DummyController.ProductCashout)
	// f.GET("/emoney", controller.DummyController.ProductEmoney)
	// f.GET("/pulsa", controller.DummyController.ProductPulsa)
	// f.GET("/paketdata", controller.DummyController.ProductPaketData)
	//admin
	g := c.Group("/admin")
	g.POST("/login", controller.AdminControlller.LoginAdmin)
	g.GET("/dashboard", controller.AdminControlller.Dashboard)
	g.GET("/managecustomerpoint", controller.AdminControlller.ManageCustomerPoint)
	g.GET("/managecustomer", controller.AdminControlller.ManageCustomer)
	g.GET("/historycustomer", controller.AdminControlller.CustomerHistory)
	g.PUT("/managecustomer/:id", controller.AdminControlller.UpdateCustomer)

	//mitra
	m := c.Group("/mitra")
	m.POST("/login", controller.MitraController.LoginMitra)
	m.POST("/register", controller.MitraController.Register)

	c.POST("/callback", controller.DummyController.CallbackXendit)
}
