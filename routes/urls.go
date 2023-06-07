package routes

import (
	"github.com/ilhamnyto/url-shortener-go/controller"
	"github.com/ilhamnyto/url-shortener-go/middleware"
	"github.com/labstack/echo/v4"
)

func UrlRouter(e *echo.Echo, c controller.UrlController) {
	urlGroup := e.Group("api/v1/url", middleware.ValidateAuth)


	urlGroup.GET("", c.GetUserUrls)
	urlGroup.GET("/:username", c.GetUserUrlsByUsername)
	urlGroup.POST("/create", c.CreateUrl)

	e.GET("/:short_url", c.Redirect)
}