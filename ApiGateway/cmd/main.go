package main

import (
	"log"

	"example.com/authProject/ApiGateway/config"
	proto "example.com/authProject/ApiGateway/proto"

	"example.com/authProject/ApiGateway/pkg/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	proto.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
