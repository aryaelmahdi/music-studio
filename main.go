package main

import (
	"project/config"
	id "project/features/instruments/data"
	ih "project/features/instruments/handler"
	is "project/features/instruments/service"
	rd "project/features/rooms/data"
	rh "project/features/rooms/handler"
	rs "project/features/rooms/service"
	ud "project/features/users/data"
	uh "project/features/users/handler"
	us "project/features/users/service"
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

	generator := helper.NewGenerator()
	jwt := helper.NewJWT(config.SECRET)

	userData := ud.NewUserData(client)
	userServices := us.NewUserService(userData, generator, jwt)
	userHandler := uh.NewUserHandler(userServices)

	roomData := rd.NewRoomData(client)
	roomServices := rs.NewRoomService(roomData)
	roomHandler := rh.NewRoomHandler(roomServices)

	instrumentData := id.NewInstrumentData(client)
	instrumentService := is.NewInstrumentService(instrumentData, jwt)
	instrumentHandler := ih.NewInstrumentHandler(instrumentService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.UserRoutes(e, userHandler)
	routes.RoomRoutes(e, roomHandler, config.SECRET)
	routes.InstrumentsRoutes(e, instrumentHandler, config.SECRET)

	e.Logger.Fatal(e.Start(":8080").Error())
}
