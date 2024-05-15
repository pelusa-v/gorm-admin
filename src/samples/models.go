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
	Id         uint `gorm:"primaryKey"`
	Name       string
	Email      string
	Color      string
	Number     string
	Passengers int
}

type Hint struct {
	Test string
}

type Author struct {
	Name  string
	Email string
	// Hint  Hint `gorm:"embedded;embeddedPrefix:hint_"`
	// Hint Hint `gorm:"embedded"`
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}
