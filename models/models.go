package models

type Product struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64
	Stock       int64
}
