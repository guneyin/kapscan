package mw

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

type status string

const (
	statusSuccess  status = "SUCCESS"
	statusError    status = "ERROR"
	statusNotfound status = "NOT_FOUND"
)

type ResponseHTTP struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func OK(c *fiber.Ctx, msg string, data any) error {
	return c.Status(fiber.StatusOK).JSON(ResponseHTTP{
		Status:  string(statusSuccess),
		Message: msg,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, err error) error {
	statusCode := fiber.StatusInternalServerError
	statusMsg := statusError

	if errors.Is(err, ErrNotFound) {
		statusCode = fiber.StatusNotFound
		statusMsg = statusNotfound
	}

	return c.Status(statusCode).JSON(ResponseHTTP{
		Status:  string(statusMsg),
		Message: err.Error(),
		Data:    nil,
	})
}
