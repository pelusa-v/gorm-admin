package samples

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Product struct {
	Id    uint
	Name  string
	Email string
}
