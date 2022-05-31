package dummy

import (
	"api-redeem-point/business/dummy"
	"api-redeem-point/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) dummy.Repository {
	dummyRepo := NewMysqlRepository(dbCon.Mysql)
	return dummyRepo
}
