package routes

import (
	"e-commerce/controllers"
	"e-commerce/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Thoha")
	})

	//USERS
	e.GET("/users", controllers.FetcUsers, middleware.IsAuthenticated)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser, middleware.IsAuthenticated)
	e.DELETE("/users/:id", controllers.DeleteUser, middleware.IsAuthenticated)
	e.POST("/login", controllers.LoginUser)

	//GENRES
	e.GET("/genres", controllers.FetchGenres, middleware.IsAuthenticated)
	e.POST("/genres", controllers.CreateGenre, middleware.IsAuthenticated)
	e.PUT("/genres/:id", controllers.UpdateGenre, middleware.IsAuthenticated)
	e.DELETE("/genres/:id", controllers.DeleteGenre, middleware.IsAuthenticated)

	//PRODUCTS
	e.GET("/products", controllers.FetchProductsWithGenres, middleware.IsAuthenticated)
	e.POST("/products", controllers.CreateProduct, middleware.IsAuthenticated)
	e.PUT("/products/:id", controllers.UpdateProduct, middleware.IsAuthenticated)
	e.DELETE("/products/:id", controllers.DeleteProduct, middleware.IsAuthenticated)

	//CART
	e.GET("/carts", controllers.FetchCartByUserId, middleware.IsAuthenticated)
	e.POST("/carts", controllers.AddToCart, middleware.IsAuthenticated)
	e.PUT("/carts/:id", controllers.EditQty, middleware.IsAuthenticated)
	e.DELETE("/carts/:id", controllers.DeleteCart, middleware.IsAuthenticated)

	//ORDERS
	e.GET("/orders", controllers.FetchOrderByUserId, middleware.IsAuthenticated)
	e.POST("/orders", controllers.MoveCartToOrder, middleware.IsAuthenticated)
	e.DELETE("/orders/:id", controllers.DeleteOrder, middleware.IsAuthenticated)
	e.PUT("/orders/:id", controllers.UpdateOrderStatus, middleware.IsAuthenticated)

	return e
}
