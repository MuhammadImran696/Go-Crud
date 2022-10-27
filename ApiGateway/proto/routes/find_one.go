package routes

import (
	"context"
	"net/http"

	pb "example.com/authProject/ApiGateway/proto/pb"
	"github.com/gin-gonic/gin"
	// "github.com/YOUR_USERNAME/go-grpc-api-gateway/pkg/product/pb"
)

func FineOne(ctx *gin.Context, c pb.CompanyServiceClient) {
	id := ctx.Param("id")
	// fmt.Print(id)
	res, err := c.FindOne(context.Background(), &pb.GetOneRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
