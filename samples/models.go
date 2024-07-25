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
	Hint Hint `gorm:"embedded"`
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded"`
	Upvotes int32
}

// Belongs to relation
type Employee struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

// Has one relation
type Person struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number   string
	PersonID uint
}

type Person1 struct {
	gorm.Model
	Name        string
	CreditCard1 CreditCard1 `gorm:"foreignKey:Person1Name"`
}

type CreditCard1 struct {
	gorm.Model
	Number      string
	Person1Name string
}

// Has many relation
type BeastType struct {
	gorm.Model
	Beasts []Beast
}

type Beast struct {
	gorm.Model
	Number      string
	BeastTypeID uint
}
