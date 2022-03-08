package main

import (
	"BulindingGoRestAPI/Database"
	"BulindingGoRestAPI/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello This is my first Restful API in GoLang")

}

func SetUpRoutes(app *fiber.App){
	app.Get("/", hello)
	// users endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)

}


func main() {
	database.ConnectDb()
	app := fiber.New()
	SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))

}