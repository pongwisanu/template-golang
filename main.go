package main

import (
	"nida/vmwareapi/handlers"
	"nida/vmwareapi/repository"
	"nida/vmwareapi/services"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// godotenv.Load()

	// app := fiber.New()

	// routes.Register(app)

	// log.Fatal(app.Listen(":8000"))

	cusRepo := repository.NewCustomerRepositoryConn()
	cusService := services.NewCustomerService(cusRepo)
	cusHandle := handlers.NewCustomerHandler(cusService)

	router := fiber.New()

	router.Get("/customer", cusHandle.GetCustomers)
	router.Get("/customer/:id", cusHandle.GetCustomer)

	router.Listen(":8000")
}
