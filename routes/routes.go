package routes

import (
	"go-training-restful/constants"
	"go-training-restful/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUsersController)

	eAuth := e.Group("")
	eAuth.Use(middleware.JWT([]byte(constants.SECRET)))
	eAuth.GET("/users", controllers.GetUserControllers)

	return e
}
