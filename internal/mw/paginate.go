package mw

import (
	"github.com/gofiber/fiber/v2"
)

func GetPaginate(c *fiber.Ctx) (int, int) {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 20)

	if pageSize > 20 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	return offset, pageSize
}
