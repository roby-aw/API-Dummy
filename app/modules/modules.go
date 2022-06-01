package modules

import (
	"api-redeem-point/api"
	adminApi "api-redeem-point/api/admin"
	dummyApi "api-redeem-point/api/dummy"
	adminBusiness "api-redeem-point/business/admin"
	dummyBusiness "api-redeem-point/business/dummy"
	"api-redeem-point/config"
	dummyRepo "api-redeem-point/repository/dummy"
	"api-redeem-point/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	dummyPermitRepository := dummyRepo.RepositoryFactory(dbCon)
	dummyPermitService := dummyBusiness.NewService(dummyPermitRepository)
	dummyPermitController := dummyApi.NewController(dummyPermitService)

	controller := api.Controller{
		AdminControlller: adminPermitController,
		DummyController:  dummyPermitController,
	}
	return controller
}
