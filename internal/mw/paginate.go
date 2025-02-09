package mw

import (
	"github.com/gofiber/fiber/v2"
)

func GetPaginate(c *fiber.Ctx) (int16, int16) {
	page := c.QueryInt("page", 1)
	size := c.QueryInt("size", 20)

	if size > 20 {
		size = 20
	}

	//offset := (page - 1) * pageSize

	return int16(page), int16(size)
}
