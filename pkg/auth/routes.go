package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/jp-api-gateway/pkg/auth/routes"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
)

func RegisterRoutes(r *gin.Engine, c config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
}

func (c *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, c.Client)
}

func (c *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, c.Client)
}
