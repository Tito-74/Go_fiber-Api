package routes

import (
	"BulindingGoRestAPI/models"
	"time"
	"BulindingGoRestAPI/Database"
	"errors"

	"github.com/gofiber/fiber/v2"
)
type Order struct{
	ID           uint `json:"id"`
	User         User  `json:"user"`
	Product      Product `json:"product"`
	CreatedAt    time.Time `json:"order_date"`

}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error{
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	 
		
	database.Database.Db.Create(&order)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrders := CreateResponseOrder(order,responseUser,responseProduct)

	return c.Status(200).JSON(responseOrders)
}

func GetOrders(c *fiber.Ctx) error{
	orders := [] models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders{
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user),CreateResponseProduct(product))  
		responseOrders = append(responseOrders, responseOrder)




	}

	return c.Status(200).JSON(responseOrders)
	

}

func FindOrder(id int, order *models.Order) error {

	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil

}
func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var order models.Order

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindOrder(id, &order); err != nil {
		c.Status(400).JSON(err.Error())
		
	}
	var user models.User
	var product models.Product
	database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
	responseOrder := CreateResponseOrder(order, CreateResponseUser(user),CreateResponseProduct(product)) 

	// responseOrders := CreateResponseOrder(order)
	return c.Status(200).JSON(responseOrder)

}
