package main

import (
	"fmt"

	"github.com/Anuragdubeyy/ticket-booking-api/config"
	"github.com/Anuragdubeyy/ticket-booking-api/db"
	"github.com/Anuragdubeyy/ticket-booking-api/handlers"
	repository "github.com/Anuragdubeyy/ticket-booking-api/repositories"
	"github.com/gofiber/fiber/v2"
)

func main(){
	EnvConfig := config.NewEnvConfig()
	db := db.Init(EnvConfig, db.DBMigrate)
	app := fiber.New(fiber.Config{
		AppName: "Ticket Booking API",
		ServerHeader: "Fiber",
	})

	// repository
	eventRepository := repository.NewEventRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	authRepository := repository.NewAuthRepository(db)

	// service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")

	// handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/ticket"), ticketRepository)
	handlers.NewAuthHandler(server.Group("/auth"), authService)

	privateRoutes := server.Use(middleware.AuthProtected(db))
	// 
	app.Listen(fmt.Sprint(":" + EnvConfig.ServerPort))
}