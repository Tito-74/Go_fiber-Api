package routes

import (
	"errors"
	"BulindingGoRestAPI/models"
	"BulindingGoRestAPI/Database"
	"github.com/gofiber/fiber/v2"
	
)

type Product struct{
	// Serializer
	ID           uint `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productmodels models.Product) Product {
	return Product{ID: productmodels.ID, Name: productmodels.Name, SerialNumber:productmodels.SerialNumber}

}

func CreateProduct(c *fiber.Ctx) error{
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error{
	products := [] models.Product{}

	database.Database.Db.Find(&products)
	responseProducts := [] Product{}
	for _, product := range products{
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func FindProduct(id int, product *models.Product) error {

	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil

}

func Getproduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
		
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
		
	}

	type UpdateProduct struct{
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`

	}
	var UpdataData UpdateProduct

	if err := c.BodyParser(&UpdataData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = UpdataData.Name
	product.SerialNumber = UpdataData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)


}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
		
	}

	if err = database.Database.Db.Delete(&product).Error; err != nil {
		c.Status(400).JSON(err.Error())

	}
	return c.Status(200).SendString("Product successfully deleted")

}
