package routes

import (
	"BulindingGoRestAPI/Database"
	"BulindingGoRestAPI/models"
	"github.com/gofiber/fiber/v2"
	"errors"
)
type User struct{
	// Serializer
	ID        uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

}
func CreateResponseUser(usermodel models.User) User{
	return User{ID: usermodel.ID, FirstName: usermodel.FirstName, LastName:usermodel.LastName}
}
func CreateUser(c *fiber.Ctx) error{
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error{

	users := []models.User{} 

	database.Database.Db.Find(&users)
	responseUsers := [] User{}
	for _, users := range users {
		responseUser := CreateResponseUser(users)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)



}

func FindUser(id int, user *models.User) error {

	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}
	return nil

}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindUser(id, &user); err != nil {
		c.Status(400).JSON(err.Error())
		
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}