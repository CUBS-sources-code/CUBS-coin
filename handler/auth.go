package handler

import (
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) authHandler {
	return authHandler{authService: authService}
}

func (h authHandler) SignUp(c *fiber.Ctx) error {
	
	var request service.SignUpRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.authService.SignUp(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}

func (h authHandler) SignIn(c *fiber.Ctx) error {
	
	var request service.SignInRequest
    if err := c.BodyParser(&request); err != nil {
       return handlerError(c, err)
    }

	token, err := h.authService.SignIn(request)
	if err != nil {
		return handlerError(c, err)
	}

	return c.JSON(token)
}