package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := InitDB()
	if err != nil {
		panic(err)
	}
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000"},
        AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))

	
	//routes for products
	e.POST("/products", RouteHandler(createProduct(db)))
	e.GET("/products", RouteHandler(getAllProducts(db)))
	e.GET("/products/:id",RouteHandler(getProductById(db)))
	e.PUT("/products/:id",RouteHandler(updateProduct(db)))
	e.DELETE("/products/:id",RouteHandler(deleteProduct(db)))

	// //routes for categories
	e.POST("/categories", RouteHandler(createCategory(db)))
	e.GET("/categories", RouteHandler(getAllCategories(db)))
	e.GET("/categories/:id",RouteHandler(getCategoryById(db)))
	e.PUT("/categories/:id",RouteHandler(updateCategory(db)))
	e.DELETE("/categories/:id",RouteHandler(deleteCategory(db)))

	//routes for cart
	e.POST("/cart", RouteHandler(addCartItem(db)))
	e.GET("/cart", RouteHandler(getCartItems(db)))
	e.GET("/cart/:id", RouteHandler(getCartItemById(db)))
	e.PUT("/cart/:id",RouteHandler(updateCartItem(db)))
	e.DELETE("/cart/:id",RouteHandler(deleteCartItem(db)))


	//payments routes
	e.POST("/payment", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"Success")
    })
	e.Logger.Fatal(e.Start(":8080"))
}	