package main

import (
	"project/config"
	id "project/features/instruments/data"
	ih "project/features/instruments/handler"
	is "project/features/instruments/service"
	pd "project/features/payments/data"
	ph "project/features/payments/handler"
	ps "project/features/payments/service"
	resd "project/features/reservations/data"
	resh "project/features/reservations/handler"
	ress "project/features/reservations/service"
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
	cfg, smtp := config.InitConfig()
	mdClient := config.MidtransConfig(cfg.MDServerKey)

	db, client, fcm := database.InitFirebaseApp(cfg.SDKPath, cfg.ProjectID, cfg.DatabaseURL)
	if db == nil {
		e.Logger.Fatal("db nil")
	}

	generator := helper.NewGenerator()
	jwt := helper.NewJWT(cfg.SECRET)

	userData := ud.NewUserData(client)
	userServices := us.NewUserService(userData, generator, jwt)
	userHandler := uh.NewUserHandler(userServices)

	roomData := rd.NewRoomData(client)
	roomServices := rs.NewRoomService(roomData, jwt)
	roomHandler := rh.NewRoomHandler(roomServices)

	instrumentData := id.NewInstrumentData(client)
	instrumentService := is.NewInstrumentService(instrumentData, jwt)
	instrumentHandler := ih.NewInstrumentHandler(instrumentService)

	reservationData := resd.NewReservationData(client)
	reservationService := ress.NewReservationService(reservationData, jwt)
	reservationHandler := resh.NewReservationHandler(reservationService)

	paymentData := pd.NewPaymentData(mdClient, client, fcm)
	paymentService := ps.NewPaymentService(paymentData, jwt, *smtp)
	paymentHandler := ph.NewPaymentHandler(paymentService)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.UserRoutes(e, userHandler)
	routes.RoomRoutes(e, roomHandler, cfg.SECRET)
	routes.InstrumentsRoutes(e, instrumentHandler, cfg.SECRET)
	routes.ReservationRoutes(e, reservationHandler, cfg.SECRET)
	routes.PaymentRoutes(e, paymentHandler, cfg.SECRET)

	e.Logger.Fatal(e.Start(":8080").Error())
}
