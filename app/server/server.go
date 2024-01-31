package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poster-keisuke/sample-clearn-architecture/app/controller"
	"github.com/poster-keisuke/sample-clearn-architecture/app/infra/sqlite3/repository"
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
	h := controller.NewHandler(
		product.NewCreateProductUseCase(productRepository),
		product.NewUpdateProductUseCase(productRepository, transaction),
		product.NewGetProductUseCase(productRepository),
	)

	r.HandleFunc("/api/products", h.CreateProduct).Methods(http.MethodPost)
	r.HandleFunc("/api/products/{id}", h.UpdateProduct).Methods(http.MethodPut)
	r.HandleFunc("/api/products/{id}", h.GetProduct).Methods(http.MethodGet)
}
