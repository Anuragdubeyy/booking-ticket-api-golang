package main

import (
	"github.com/Anuragdubeyy/ticket-booking-api/handlers"
	repository "github.com/Anuragdubeyy/ticket-booking-api/repositories"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New(fiber.Config{
		AppName: "Ticket Booking API",
		ServerHeader: "Fiber",
	})

	// repository
	eventRepository := repository.NewEventRepository(nil)

	server := app.Group("/api")

	// handler
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}