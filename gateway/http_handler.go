package main

import (
	"net/http"

	"github.com/ajayjadhav201/common"
	pb "github.com/ajayjadhav201/common/api"
)

type Handler struct {
	//gateway
	Client pb.OrderServiceClient
}

func NewHandler(service pb.OrderServiceClient) *Handler {
	return &Handler{service}
}

func (h *Handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandlerCreateOrder)
}

func (h *Handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	//
	customerID := r.PathValue("customerID")
	common.Println("ajaj customer id is ", customerID)

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UnmarshalFailed)
		return
	}

	// pb.NewOrderServiceClient()
	order, err := h.Client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, common.InternalServerErr)
		return
	}
	common.WriteJSON(w, http.StatusOK, order)
}
