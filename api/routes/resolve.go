package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/prasanth-pn/UrlShortner_GO-Redis/database"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error:": "short not found"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to db",
		})
	}
	rIr := database.CreateClient(1)
	defer rIr.Close()
	_ = rIr.Incr(database.Ctx, "counter")
	return c.Redirect(value, 301)
}
