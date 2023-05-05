package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type HandlerFunc func(c echo.Context) error

func RouteHandler(handler HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return handler(c)
	}
}


func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// db.Exec("DROP TABLE IF EXISTS products")
	// db.Exec("DROP TABLE IF EXISTS categories")
	db.Exec("DROP TABLE IF EXISTS cart_items")
	db.AutoMigrate(&Product{}, &Category{},&CartItem{})
	return db, nil
}