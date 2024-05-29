package products

import (
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
}
