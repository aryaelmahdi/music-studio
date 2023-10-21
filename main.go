package main

import (
	"project/config"
	"project/features/users/data"
	"project/features/users/handler"
	"project/features/users/service"
	"project/helper"
	"project/routes"
	"project/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	config := config.InitConfig()

	db, client := database.InitFirebaseApp(config.SDKPath, config.ProjectID, config.DatabaseURL)
	if db == nil {
		e.Logger.Fatal("db nil")
	}
	userData := data.NewUserData(client)

	generator := helper.NewGenerator()
	jwt := helper.NewJWT(config.SECRET)
	userServices := service.NewUserService(userData, generator, jwt)

	userHandler := handler.NewUserHandler(userServices)
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))
	routes.RouteUser(e, userHandler)
	e.Logger.Fatal(e.Start(":8080").Error())
}
