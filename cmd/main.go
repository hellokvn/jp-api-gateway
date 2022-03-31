package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hellokvn/jp-api-gateway/pkg/auth"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
	"github.com/hellokvn/jp-api-gateway/pkg/wanikani"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	authSvc := auth.RegisterRoutes(app, c)
	wanikani.RegisterRoutes(app, c, authSvc)

	app.Listen(c.Port)
}
