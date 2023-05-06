package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const productsWithIdEndpoint = "/products/:id"
const categoriesWithIdEndpoint = "categoriesWithIdEndpoint"
const cartWithIdEndpoint = "cartWithIdEndpoint"

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
	e.GET(productsWithIdEndpoint,RouteHandler(getProductById(db)))
	e.PUT(productsWithIdEndpoint,RouteHandler(updateProduct(db)))
	e.DELETE(productsWithIdEndpoint,RouteHandler(deleteProduct(db)))

	// //routes for categories
	e.POST("/categories", RouteHandler(createCategory(db)))
	e.GET("/categories", RouteHandler(getAllCategories(db)))
	e.GET(categoriesWithIdEndpoint,RouteHandler(getCategoryById(db)))
	e.PUT(categoriesWithIdEndpoint,RouteHandler(updateCategory(db)))
	e.DELETE(categoriesWithIdEndpoint,RouteHandler(deleteCategory(db)))

	//routes for cart
	e.POST("/cart", RouteHandler(addCartItem(db)))
	e.GET("/cart", RouteHandler(getCartItems(db)))
	e.GET(cartWithIdEndpoint, RouteHandler(getCartItemById(db)))
	e.PUT(cartWithIdEndpoint,RouteHandler(updateCartItem(db)))
	e.DELETE(cartWithIdEndpoint,RouteHandler(deleteCartItem(db)))


	//payments routes
	e.POST("/payment", func(c echo.Context) error {
		return c.JSON(http.StatusOK,"Success")
    })
	e.Logger.Fatal(e.Start(":8080"))
}	