package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poster-keisuke/sample-clearn-architecture/app/controller"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/repository"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/order"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/product"
	"log"
	"net/http"
	"os"
)

const (
	exitOK int = iota
	exitError
)

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "OK",
	}
	_ = json.NewEncoder(w).Encode(response)
}

func Run(ctx context.Context) {
	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health_check", healthCheckHandler)
	// Product Router
	productRouter(router)

	log.Println("listening on port :8000")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Printf("%q\n", err)
		os.Exit(exitError)
	}

	os.Exit(exitOK)
}

func productRouter(r *mux.Router) {
	productRepository := repository.NewProductRepository()
	transaction := repository.NewTransaction()
	h := controller.NewProductHandler(
		product.NewCreateProductUseCase(productRepository),
		product.NewUpdateProductUseCase(productRepository, transaction),
		product.NewGetProductUseCase(productRepository),
	)

	r.HandleFunc("/api/products", h.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods(http.MethodGet)
}

func orderRouter(r *mux.Router) {
	orderRepository := repository.NewOrderRepository()
	productRepository := repository.NewProductRepository()
	transaction := repository.NewTransaction()

	//transaction := repository.NewTransaction()
	h := controller.NewOrderHandler(
		order.NewCreteOrderUseCase(orderRepository, productRepository),
		order.NewCancelOrderUseCase(orderRepository, productRepository, transaction),
		order.NewProcessOrderUseCase(orderRepository, productRepository, transaction),
	)

	r.HandleFunc("/api/orders", h.CreateOrder).Methods(http.MethodPost)
	r.HandleFunc("/api/orders/{id}", h.CancelOrder).Methods(http.MethodPut)
	r.HandleFunc("/api/orders/{id}/process", h.ProcessOrder).Methods(http.MethodPost)
}
