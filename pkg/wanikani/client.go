package wanikani

import (
	"fmt"

	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
	"github.com/hellokvn/jp-api-gateway/pkg/wanikani/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.WanikaniServiceClient
}

func InitServiceClient(c *config.Config) pb.WanikaniServiceClient {
	cc, err := grpc.Dial(c.WanikaniSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewWanikaniServiceClient(cc)
}
