package routes

import (
	"github.com/ilhamnyto/url-shortener-go/controller"
	"github.com/ilhamnyto/url-shortener-go/middleware"
	"github.com/labstack/echo/v4"
)

func UserRouter(e *echo.Echo, c *controller.UserController) {
	var (
		authGroup = e.Group("api/v1/auth")
		userGroup = e.Group("api/v1/users", middleware.ValidateAuth)
	)

	authGroup.POST("/register", c.CreateUser)
	authGroup.POST("/login", c.Login)
	userGroup.PUT("/update_password", c.UpdatePassword)
	userGroup.GET("/profile", c.UserProfile)
	userGroup.GET("/profile/:username", c.UserProfileByUsername)
}