package main

import (
	"github.com/ifty123/crawl_web/crawler"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/links", crawler.GetCraw)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
