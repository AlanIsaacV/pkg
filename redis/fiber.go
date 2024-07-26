package redis

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// Handler
// @Summary Flush Cache
// @Router /redis [DELETE]
// @Tags Cache
// @Param pattern query string false "Redis pattern"
// @Success 200 {string} null
// @Failure 400,500 {object} ehandler.ErrorMsg
func Handler(c *fiber.Ctx) error {
	pattern := c.Query("pattern")

	count, err := Flush(context.Background(), pattern)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"deleted": count})
}
