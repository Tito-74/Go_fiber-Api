package main

import (
	"BulindingGoRestAPI/Database"
	"BulindingGoRestAPI/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/arsmn/fiber-swagger/v2"
	_ "github.com/arsmn/fiber-swagger/v2/example/docs"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello This is my first Restful API in GoLang")

}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @contact.name API Support
// @contact.email youremail@provider.com
// @host localhost:3000
// @BasePath /

func SetUpRoutes(app *fiber.App){
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))


	app.Get("/", hello)
	// users endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// products endpoints
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.Getproduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	// order endpoints
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)



}


func main() {
	database.ConnectDb()
	app := fiber.New()
	SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))

}