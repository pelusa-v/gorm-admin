package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pelusa-v/gorm-admin/src/pkg/admin"
	"github.com/pelusa-v/gorm-admin/src/samples"
)

func main() {
	r := gin.Default()
	db := samples.NewDbInstance()

	admin := admin.NewGinAdmin(r, db)
	admin.Register()
	admin.Configure("My awesome project")
	admin.RegisterModel(samples.User{})
	admin.RegisterModel(samples.Product{})
	admin.RegisterModel(samples.Car{})
	admin.RegisterModel(samples.Blog{})
	admin.RegisterModel(samples.Employee{})
	admin.RegisterModel(samples.Company{})
	admin.RegisterModel(samples.Person{})
	admin.RegisterModel(samples.CreditCard{})
	admin.RegisterModel(samples.Person1{})
	admin.RegisterModel(samples.CreditCard1{})

	r.Run()
}
