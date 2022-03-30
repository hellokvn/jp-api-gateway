package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/auth/pb"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *fiber.Ctx, c pb.AuthServiceClient) error {
	body := LoginRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
