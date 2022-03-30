package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth/routes"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	r := app.Group("/auth")
	r.Post("/register", svc.Register)
	r.Post("/login", svc.Login)
}

func (svc *ServiceClient) Register(ctx *fiber.Ctx) error {
	return routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *fiber.Ctx) error {
	return routes.Login(ctx, svc.Client)
}
