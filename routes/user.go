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

// GetItem godoc
// @Summary Get a user
// @Description Get an user by its ID
// @ID get-item-by-int
// @Accept  json
// @Produce  json
// @Tags Item
// @Param id path int true "Item ID"
// @Success 200 {object} Item
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/users/{id} [get]
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

func UpdateUser(c *fiber.Ctx) error{
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindUser(id, &user); err != nil {
		c.Status(400).JSON(err.Error())
		
	}

	type Updateuser struct{
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`

	}
	var UpdataData Updateuser

	if err := c.BodyParser(&UpdataData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = UpdataData.FirstName
	user.LastName = UpdataData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)


}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil{
		c.Status(400).JSON("Please ensure that Id is an interger")
	}
	if err = FindUser(id, &user); err != nil {
		c.Status(400).JSON(err.Error())
		
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		c.Status(400).JSON(err.Error())

	}
	return c.Status(200).SendString("User successfully deleted")

}