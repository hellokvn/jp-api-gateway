package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/card/pb"
)

type AddLevelBody struct {
	Level int32 `json:"level"`
}

func AddLevel(ctx *fiber.Ctx, c pb.CardServiceClient) error {
	body := AddLevelBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.AddLevel(context.Background(), &pb.AddLevelRequest{
		Level:  body.Level,
		UserId: ctx.Locals("userId").(int32),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.JSON(&res)
}
