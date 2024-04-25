package samples

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Product struct {
	ID    uint
	Name  string
	Email string
}

type Car struct {
	Id    uint `gorm:"primaryKey"`
	Name  string
	Email string
}
