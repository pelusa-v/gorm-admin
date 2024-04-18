package samples

import (
	"fmt"
	"reflect"

	"github.com/pelusa-v/gorm-admin/src/pkg/first"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbInstance() *gorm.DB {
	dsn := "admon:admon@tcp(127.0.0.1:3306)/gorm-admin-sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	db.AutoMigrate(&User{}, &Product{})
	return db
}

func TestListHandler(appDb *gorm.DB) {
	handler := first.NewCrudHandler(reflect.TypeOf(Product{}), appDb)

	objectsList := handler.ListOjects()

	for _, v := range objectsList {
		p := v.(Product)
		fmt.Println(p.Id)
		fmt.Println(p.Name)
		fmt.Println(p.Email)
		fmt.Println("..................")
	}
}
