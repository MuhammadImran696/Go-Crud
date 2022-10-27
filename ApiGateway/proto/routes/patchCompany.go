package routes

import (
	"context"
	"net/http"

	pb "example.com/authProject/ApiGateway/proto/pb"
	"github.com/gin-gonic/gin"
	// "github.com/YOUR_USERNAME/go-grpc-api-gateway/pkg/product/pb"
)

type UpdateProductRequestBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Employees   int64  `json:"employees"`
	Registered  bool   `json:"registered"`
	Type        string `json:"type"`
}

func UpdateCompany(ctx *gin.Context, c pb.CompanyServiceClient) {
	id := ctx.Param("id")
	body := UpdateProductRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := c.UpdateCompany(context.Background(), &pb.PatchRequest{
		Id:          id,
		Name:        body.Name,
		Description: body.Description,
		Employees:   body.Employees,
		Registered:  body.Registered,
		Type:        body.Type,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
