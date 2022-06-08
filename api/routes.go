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
	c.PUT("/customer", controller.DummyController.UpdateCustomer)
	c.GET("/history/:idcustomer", controller.DummyController.History)
	c.GET("/detailhistory/:transactionid", controller.DummyController.DetailTransaction)
	c.GET("/stockproduct", controller.DummyController.StockProduct)
	c.PUT("/stockproduct/:id", controller.DummyController.ManageStockProduct)
	c.POST("/pulsa", controller.DummyController.RedeemPulsa)
	c.POST("/paketdata", controller.DummyController.RedeemPulsa)
	c.POST("/emoney", controller.DummyController.RedeemPulsa)
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
	g.GET("/customerpoint", controller.AdminControlller.ManageCustomerPoint)
	g.PUT("/customerpoint", controller.AdminControlller.UpdateCustomerPoint)
	g.GET("/customer", controller.AdminControlller.ManageCustomer)
	g.GET("/historycustomer", controller.AdminControlller.CustomerHistory)
	g.PUT("/customer", controller.AdminControlller.UpdateCustomer)
	g.POST("/approve/:transactionid", controller.AdminControlller.ApproveTransaction)

	//mitra
	m := c.Group("/mitra")
	m.POST("/login", controller.DummyController.LoginMitra)
	m.POST("/register", controller.DummyController.RegisterMitra)
	m.POST("/inputpoin", controller.DummyController.InputPoin)
	m.GET("/history", controller.DummyController.HistoryMitra)

	c.POST("/callback", controller.DummyController.CallbackXendit)
}
