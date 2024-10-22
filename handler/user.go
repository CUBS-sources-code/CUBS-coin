package handler

import (
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) GetUsers(c *fiber.Ctx) error {
	
	users, err := h.userService.GetUsers()
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func (h userHandler) GetUser(c *fiber.Ctx) error {
	
	student_id := c.Params("student_id")

	user, err := h.userService.GetUser(student_id)
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(user)
}


func (h userHandler) GetMyUser(c *fiber.Ctx) error {
	
	token := c.Locals("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	id := claims["user"].(string)

	user, err := h.userService.GetUser(id)
	if err != nil {
		return handlerError(c, err)
	}
	return c.JSON(user)
}

func (h userHandler) CreateUser(c *fiber.Ctx) error {
	
	var request service.NewUserRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	user, err := h.userService.CreateUser(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(user)
}