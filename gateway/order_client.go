package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
)

type OrderClient struct {
	Client pb.OrderServiceClient
}

func NewHandler(service pb.OrderServiceClient) *OrderClient {
	return &OrderClient{service}
}

func (o *OrderClient) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", o.CreateOrderHandler)
}

//
//
//
//
//

func (o *OrderClient) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	//
	customerID := r.PathValue("customerID")
	common.Println("ajaj customer id is ", customerID)

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}

	// pb.NewOrderServiceClient()
	order, err := o.Client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, common.InternalServerErr)
		return
	}
	common.WriteJSON(w, http.StatusOK, order)
}
