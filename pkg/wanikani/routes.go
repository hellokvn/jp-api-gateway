package wanikani

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
	"github.com/hellokvn/jp-api-gateway/pkg/wanikani/routes"
)

func RegisterRoutes(app *fiber.App, c config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	r := app.Group("/wanikani")
	r.Use(a.AuthRequired)
	r.Post("/find", svc.Find)
}

func (svc *ServiceClient) Find(ctx *fiber.Ctx) error {
	return routes.Find(ctx, svc.Client)
}
