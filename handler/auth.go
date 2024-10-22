package handler

import (
	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/spf13/viper"
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

func (h authHandler) AuthErrorHandler(c *fiber.Ctx, err error) error {
	return handlerError(c, errs.NewUnAuthorizedError())
}

func (h authHandler) AuthSuccessHandler(c *fiber.Ctx) error {
	
	c.Next()
	return nil
}

func (h authHandler) AuthorizationRequired() fiber.Handler {
    return jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:   []byte(viper.GetString("app.jwt-secret")),
		ErrorHandler: h.AuthErrorHandler,
		SuccessHandler: h.AuthSuccessHandler,
	})
}