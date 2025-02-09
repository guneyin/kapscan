package mw

import (
	"github.com/gofiber/fiber/v2"
)

func GetPaginate(c *fiber.Ctx) (int16, int16) {
	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", 15)

	if size > 15 {
		size = 15
	}

	return int16(page), int16(size)
}
