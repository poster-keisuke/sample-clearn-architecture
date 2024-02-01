package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/order"
	"log"
	"net/http"
)

type orderHandler struct {
	createOrderUseCase *order.CreteOrderUseCase
	cancelOrderUseCase *order.CancelOrderUseCase
	//processOrderUseCase *order.OrderProcessUseCase
}

func NewOrderHandler(
	createOrderUseCase *order.CreteOrderUseCase,
	cancelOrderUseCase *order.CancelOrderUseCase,
	// processOrderUseCase *order.OrderProcessUseCase,
) orderHandler {
	return orderHandler{
		createOrderUseCase: createOrderUseCase,
		cancelOrderUseCase: cancelOrderUseCase,
		//processOrderUseCase: processOrderUseCase,
	}
}

func (h orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

}

func (h orderHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.cancelOrderUseCase.Run(ctx, id); err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	message := struct {
		Message string `json:"message"`
	}{
		Message: "success",
	}
	response, err := json.Marshal(message)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}

func (h orderHandler) ProcessOrder() {

}
