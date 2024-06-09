package main

import (
	"golang-microservices/api-gateway/auth"
	"net/http"

	"github.com/ajayjadhav201/common"
	"github.com/gin-gonic/gin"
)

type OrderClient struct {
}

func NewOrderClient() *OrderClient {
	return &OrderClient{}
}

func (o *OrderClient) RegisterRoutes(r *gin.RouterGroup) {
	//
	r.Use(auth.AuthMiddleware)
	{
		r.GET("/orders/:id")
		r.GET("/orders")
		r.POST("/orders")
		r.PUT("/orders/:id")
		r.DELETE("/orders/:id")
	}
}

func (o *OrderClient) GetOrderByIdHandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
}

func (o *OrderClient) GetOrdersHandler(c *gin.Context) {
	//
}

func (o *OrderClient) CreateOrderHandler(c *gin.Context) {
	//
}

func (o *OrderClient) UpdateOrderhandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
}

func (o *OrderClient) DeleteProductsHandler(c *gin.Context) {
	//
	pid := common.Atoi(c.Param("id"))
	if pid == -1 {
		common.WriteError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
}
