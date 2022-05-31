package dummy

import (
	"api-redeem-point/business/dummy"
	"fmt"

	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo *MysqlRepository) FindFoodByName(name string) (foods []dummy.Food, err error) {
	fmt.Println("repo jalan")
	result := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(foods) == 0 {
		return nil, fmt.Errorf("data tidak ditemukan")
	}
	return foods, nil
}
