package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"strconv"
	"fmt"
)

//function to create new product
func createProduct(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		var product Product
		if err := c.Bind(&product); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if product.Name == "" && product.Price == 0 {
			return c.JSON(http.StatusBadRequest, "Name is required")
		}
		if product.CategoryID != 0 {
			var category Category
			if err := db.Where("id = ?", product.CategoryID).First(&category).Error; err != nil {
				return c.JSON(http.StatusBadRequest, "Category not found")
			}
			product.Category = category
		} else if product.Category.Name != "" {
			var category Category
			if err := db.Where("name = ?", product.Category.Name).First(&category).Error; err != nil {
				category = Category{Name: product.Category.Name}
				db.Create(&category)
			}
			product.Category = category
		} else {
			return c.JSON(http.StatusBadRequest, "Request incompatible with schema")
		}
	
		db.Create(&product)
	
		return c.JSON(http.StatusCreated, product)
	}
}



func updateProduct(db *gorm.DB) HandlerFunc {
    return func(c echo.Context) error {
        id := c.Param("id")
        var product Product
        if err := db.Preload("Category").First(&product, id).Error; err != nil {
            return c.JSON(http.StatusBadRequest, "Product not found")
        }

        if err := c.Bind(&product); err != nil {
            return c.JSON(http.StatusBadRequest, err.Error())
        }

        if product.Category.ID != 0 {
            var category Category
            if err := db.Where("id = ?", product.CategoryID).First(&category).Error; err != nil {
                return c.JSON(http.StatusBadRequest, "Category not found")
            }
            product.Category = category
        } else {
            return c.JSON(http.StatusBadRequest, "Request incompatible with schema")
        }

        db.Save(&product)

        return c.JSON(http.StatusOK, product)
    }
}

func getAllProducts(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		var products []Product
		category := c.QueryParam("category")
		fmt.Println(category)
		if category != "" {
			db.Preload("Category").Joins("left join categories on products.category_id = categories.id").Where("categories.name = ?", category).Find(&products)
		} else {
			db.Preload("Category").Find(&products)
		}

		return c.JSON(http.StatusOK, products)
	}
}


func getProductById(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var product Product
		if err := db.Preload("Category").First(&product, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, product)
	}
}

func deleteProduct(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		
		
		if err := db.Delete(&Product{},id).Error; err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, "Item deleted")
	}
}



