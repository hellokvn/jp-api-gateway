package card

import (
	"fmt"

	"github.com/hellokvn/jp-api-gateway/pkg/card/pb"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CardServiceClient
}

func InitServiceClient(c *config.Config) pb.CardServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCardServiceClient(cc)
}
