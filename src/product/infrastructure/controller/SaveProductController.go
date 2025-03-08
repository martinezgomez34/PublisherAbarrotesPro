package controllers

import (
	"encoding/json"
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
	"github.com/google/uuid"
)

type ProductController struct {
	UseCase *application.ProductUseCase
}

func NewProductController(useCase *application.ProductUseCase) *ProductController {
	return &ProductController{UseCase: useCase}
}

func (pc *ProductController) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}


	product := domain.Product{
		ID:          uuid.New().String(), 
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
	}

	message := domain.Message{
		Type:    domain.MessageTypeCreateProduct,
		Product: product,
	}

	if err := pc.UseCase.CreateProduct(message); err != nil {
		http.Error(w, "Error al crear el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto creado exitosamente"})
}