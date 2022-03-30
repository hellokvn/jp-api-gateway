package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth"
	"github.com/hellokvn/jp-api-gateway/pkg/card"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	authSvc := auth.RegisterRoutes(app, c)
	card.RegisterRoutes(app, c, authSvc)

	app.Listen(c.Port)
}
