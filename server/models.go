package main

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

type Product struct {
	gorm.Model
	Name       string    `json:"name"`
	Price      float32   `json:"price"`
	Description string 	 `json:"description"` 
	CategoryID uint      `json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
}

type CartItem struct {
    gorm.Model
    Product   Product
    ProductID uint   	`json:"product_id"`
    Quantity  uint		`json:"quantity"` 
}