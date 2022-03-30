package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *fiber.Ctx) error {
	authorization := ctx.Get("authorization")

	if authorization == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	ctx.Locals("userId", res.UserId)

	return ctx.Next()
}
