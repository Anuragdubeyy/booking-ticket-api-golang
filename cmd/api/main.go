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

	server := app.Group("/api")

	// handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprint(":" + EnvConfig.ServerPort))
}