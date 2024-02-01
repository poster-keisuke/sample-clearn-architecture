package controller

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	orderDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/order"
	"io"
	"log"
	"net/http"
)

type orderHandler struct {
	createOrderUseCase  *order.CreateOrderUseCase
	cancelOrderUseCase  *order.CancelOrderUseCase
	processOrderUseCase *order.ProcessOrderUseCase
}

func NewOrderHandler(
	createOrderUseCase *order.CreateOrderUseCase,
	cancelOrderUseCase *order.CancelOrderUseCase,
	processOrderUseCase *order.ProcessOrderUseCase,
) orderHandler {
	return orderHandler{
		createOrderUseCase:  createOrderUseCase,
		cancelOrderUseCase:  cancelOrderUseCase,
		processOrderUseCase: processOrderUseCase,
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

func (h orderHandler) ProcessOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	var input struct {
		Type orderDomain.OrderProcessType `json:"type"`
	}

	if err = json.Unmarshal(body, &input); err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	o, err := h.processOrderUseCase.Run(ctx, id, input.Type)
	if err != nil {
		log.Printf("%+v\n", err)
		if errors.Is(err, errors.New("conflict between the process type and the status")) {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	response, err := json.Marshal(o)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
