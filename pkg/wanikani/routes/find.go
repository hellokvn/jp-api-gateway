package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/jp-api-gateway/pkg/wanikani/pb"
)

type CardType string

type FindBody struct {
	Level int32  `json:"level"`
	Type  string `json:"type"`
}

func Find(ctx *fiber.Ctx, c pb.WanikaniServiceClient) error {
	body := FindBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.Find(context.Background(), &pb.FindRequest{
		Level: body.Level,
		Type:  body.Type,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
