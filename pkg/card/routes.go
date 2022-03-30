package card

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth"
	"github.com/hellokvn/jp-api-gateway/pkg/card/routes"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	r := app.Group("/card")
	r.Use(a.AuthRequired)
	r.Post("/level", svc.AddLevel)
}

func (svc *ServiceClient) AddLevel(ctx *fiber.Ctx) error {
	return routes.AddLevel(ctx, svc.Client)
}
