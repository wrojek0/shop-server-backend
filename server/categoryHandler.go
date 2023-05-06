package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)


const categoryNotFoundMsg = "Category not found"

func createCategory(db *gorm.DB) HandlerFunc {
    return func(c echo.Context) error {
        var category Category
        if err := c.Bind(&category); err != nil {
            return c.JSON(http.StatusBadRequest, err.Error())
        }

        db.Create(&category)

        return c.JSON(http.StatusCreated, category)
    }
}

func getAllCategories(db *gorm.DB) HandlerFunc {
    return func(c echo.Context) error {
        var categories []Category

        db.Find(&categories)

        return c.JSON(http.StatusOK, categories)
    }
}

func getCategoryById(db *gorm.DB) HandlerFunc {
    return func(c echo.Context) error {
        var category Category
        id := c.Param("id")

        if err := db.First(&category, id).Error; err != nil {
            return c.JSON(http.StatusNotFound, categoryNotFoundMsg)
        }

        return c.JSON(http.StatusOK, category)
    }
}

func updateCategory(db *gorm.DB) HandlerFunc {
    return func(c echo.Context) error {
        var category Category
        id := c.Param("id")

        if err := db.First(&category, id).Error; err != nil {
            return c.JSON(http.StatusNotFound, categoryNotFoundMsg)
        }

        if err := c.Bind(&category); err != nil {
            return c.JSON(http.StatusBadRequest, err.Error())
        }

        db.Save(&category)

        return c.JSON(http.StatusOK, category)
    }
}

func deleteCategory(db *gorm.DB) HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		var category Category
		if err := db.First(&category, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, categoryNotFoundMsg)
		}

		if err := db.Where("category_id = ?", id).Delete(&Product{}).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if err := db.Delete(&category).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, "Category deleted")
	}
}
