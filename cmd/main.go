package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hellokvn/jp-api-gateway/pkg/common/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// auth.RegisterRoutes(r, c)

	r.Run(c.Port)
}
