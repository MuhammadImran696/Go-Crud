package product

import (
	"example.com/authProject/ApiGateway/config"
	"example.com/authProject/ApiGateway/pkg/auth"
	_ "example.com/authProject/ApiGateway/proto/pb"
	"example.com/authProject/ApiGateway/proto/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateCompany)
	routes.GET("/:id", svc.FindOne)
	routes.DELETE("/:id", svc.DeleteCompany)
	routes.PATCH("/:id", svc.UpdateCompany)

}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateCompany(ctx *gin.Context) {
	routes.CreateCompany(ctx, svc.Client)
}
func (svc *ServiceClient) DeleteCompany(ctx *gin.Context) {
	routes.DeleteCompany(ctx, svc.Client)
}
func (svc *ServiceClient) UpdateCompany(ctx *gin.Context) {
	routes.UpdateCompany(ctx, svc.Client)
}
