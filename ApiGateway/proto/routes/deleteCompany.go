package routes

import (
	"context"
	"net/http"

	pb "example.com/authProject/ApiGateway/proto/pb"
	"github.com/gin-gonic/gin"
)

func DeleteCompany(ctx *gin.Context, c pb.CompanyServiceClient) {
	id := ctx.Param("id")

	res, err := c.DeleteCompany(context.Background(), &pb.DeleteCompanyRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
