package main

import (
	"os"

	"github.com/ilhamnyto/url-shortener-go/config"
	"github.com/ilhamnyto/url-shortener-go/controller"
	"github.com/ilhamnyto/url-shortener-go/pkg/database"
	"github.com/ilhamnyto/url-shortener-go/repositories"
	"github.com/ilhamnyto/url-shortener-go/routes"
	"github.com/ilhamnyto/url-shortener-go/services"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig(".env")

	db := database.ConnectDB()

	e := echo.New()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	routes.UserRouter(e, userController)

	urlRepository := repositories.NewUrlRepository(db)
	urlService := services.NewUrlServices(urlRepository)
	urlController := controller.NewUrlController(urlService)
	routes.UrlRouter(e, *urlController)

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}