package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.APP) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}
func main() {
	fmt.Println("Hello word")
	err := godotenv.Load() 

	if err != nil {
		fmt.Println("errot loading env file")
	}

	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	err=app.Listen(os.Getenv("APP_PORT"))
	log.Fatal(err)

}
