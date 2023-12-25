package main

import (
	"log"
	"time"

	"github.com/abcdan/nxt/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load .env file, using environment variables.")
	}

	app.Static("/", "./public")
	app.Static("/404", "./public/404.html")

	app.Use("/api", limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
	}))

	routes.LinkRoutes(app)
	routes.RedirectRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "*" {
			return c.SendFile("./public/404.html")
		}
		return c.Next()
	})

	app.Listen(":3000")
}
