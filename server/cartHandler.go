package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	// "strconv"
	"fmt"
)

const productCategory = "Product.Category"

func addCartItem(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		var item CartItem
		if err := c.Bind(&item); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		fmt.Println("Call addCartItem was made")
		fmt.Println(item)
		
		// Check if the product exists
		var product Product
		if err := db.Preload("Category").First(&product, item.ProductID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, "Product not found")
		}
		
		// Create the cart item
		item.Product = product
		db.Create(&item)
		
		return c.JSON(http.StatusCreated, item)
	}
}

func updateCartItem(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var cartItem CartItem
		if err := db.Preload(productCategory).First(&cartItem, id).Error; err != nil {
            return c.JSON(http.StatusBadRequest, "Cart item not found")
        }
		
		var reqCartItem CartItem
		if err := c.Bind(&reqCartItem); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		cartItem.Quantity = reqCartItem.Quantity
		if err := db.Save(&cartItem).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, cartItem)
	}
}

func getCartItemById(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var item CartItem
		if err := db.Preload(productCategory).Find(&item,id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		
		return c.JSON(http.StatusOK, item)
	}
}

func getCartItems(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		var items []CartItem
		if err := db.Preload(productCategory).Find(&items).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		
		return c.JSON(http.StatusOK, items)
	}
}

func deleteCartItem(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		
		var item CartItem
		if err := db.First(&item, id).Error; err != nil {
			return c.JSON(http.StatusBadRequest, "Cart item not found")
		}
		db.Delete(&item)	

		return c.JSON(http.StatusOK, "Item deleted")
	}
}


