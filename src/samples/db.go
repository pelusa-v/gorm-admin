package samples

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbInstance() *gorm.DB {
	dsn := "admon:admon@tcp(127.0.0.1:3306)/gorm-admin-sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db
	}
	return nil
}
