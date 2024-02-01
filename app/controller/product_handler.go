package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/poster-keisuke/sample-clearn-architecture/app/usacase/product"
	"io"
	"log"
	"net/http"
)

type productHandler struct {
	createProductUseCase *product.CreateProductUseCase
	updateProductUseCase *product.UpdateProductUseCase
	getProductUseCase    *product.GetProductUseCase
}

func NewProductHandler(
	createProductUseCase *product.CreateProductUseCase,
	updateProductUseCase *product.UpdateProductUseCase,
	getProductUseCase *product.GetProductUseCase,
) productHandler {
	return productHandler{
		createProductUseCase: createProductUseCase,
		updateProductUseCase: updateProductUseCase,
		getProductUseCase:    getProductUseCase,
	}
}

type ProductParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
	Stock       int    `json:"stock"`
}

func (h productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	var params ProductParams
	if err = json.Unmarshal(body, &params); err != nil {
		log.Printf("%q\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input := product.CreateProductUseCaseInputDto{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
		Category:    params.Category,
		Stock:       params.Stock,
	}

	p, err := h.createProductUseCase.Run(ctx, input)
	if err != nil {
		log.Printf("%q\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(p)
	if err != nil {
		log.Printf("%q\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}

func (h productHandler) UpdateProduct(w http.ResponseWriter, _ *http.Request) {

}

func (h productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Println("id is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.getProductUseCase.Run(ctx, id)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(p)
	if err != nil {
		log.Printf("%+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(response)
}
