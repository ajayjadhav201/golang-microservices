package products

import (
	"golang-microservices/api-gateway/auth"
	"net/http"

	"github.com/ajayjadhav201/common"
	"github.com/gin-gonic/gin"
)

type ProductClient struct {
	//
}

func NewProductClient() *ProductClient {
	return &ProductClient{}
}

func (p *ProductClient) RegisterRoutes(r *gin.RouterGroup) {
	//
	r.Use(auth.AuthMiddleware)
	{
		r.GET("/products", p.GetProductsHandler)
		r.GET("/products/:id", p.GetProductByIdHandler)
		r.POST("/addproduct", p.AddProductHandler)
		r.POST("/addproduct/uploadImages", func(c *gin.Context) {})
		r.PUT("/products/:id", p.UpdateProductHandler)
		r.DELETE("/products/:id", p.DeleteProductHandler)
	}
}

func (p *ProductClient) GetProductsHandler(c *gin.Context) {
	//

}

func (p *ProductClient) GetProductByIdHandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	//
}

func (p *ProductClient) AddProductHandler(c *gin.Context) {
	//
}

func (p *ProductClient) UpdateProductHandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
}

func (p *ProductClient) DeleteProductHandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	//

}
