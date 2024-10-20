package handler

import (
	"net/http"

	"github.com/CUBS-sources-code/CUBS-coin/errs"
	"github.com/gofiber/fiber/v2"
)

func handlerError(c *fiber.Ctx, err error) error {
	switch e := err.(type) {
	case errs.AppError:
		return c.Status(e.Code).SendString(e.Message)
	}
	return c.Status(http.StatusInternalServerError).SendString(err.Error())
}