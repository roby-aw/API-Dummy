package api

import (
	"api-redeem-point/api/admin"
	"api-redeem-point/api/dummy"
	auth "api-redeem-point/api/middleware"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	AdminControlller *admin.Controller
	DummyController  *dummy.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	c := e.Group("/v1")
	c.GET("/foods/:name", controller.DummyController.GetFoodByName)
	c.POST("/login", controller.DummyController.Login)
	//admin
	g := e.Group("/admin")
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.POST("/token", controller.AdminControlller.GetToken)
	g.PUT("/:id", controller.AdminControlller.UpdateAdmin)
	g.GET("/:id", controller.AdminControlller.GetAdminByID)
	// admin using jwt
	g.DELETE("/:id", controller.AdminControlller.DeleteAdmin, auth.SetupAuthenticationJWT())
	g.GET("", controller.AdminControlller.GetAdmins, auth.SetupAuthenticationJWT())
}
