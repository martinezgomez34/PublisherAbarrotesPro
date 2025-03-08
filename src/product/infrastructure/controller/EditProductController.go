package controllers

import (
	"encoding/json"
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
)

type EditProductController struct {
	UseCase *application.EditProductUseCase
}

func NewEditProductController(useCase *application.EditProductUseCase) *EditProductController {
	return &EditProductController{UseCase: useCase}
}

func (pc *EditProductController) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	var updatedProduct struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	product := domain.Product{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	}

	message := domain.Message{
		Type:    domain.MessageTypeUpdateProduct,
		Product: product,
	}

	if err := pc.UseCase.UpdateProduct(message); err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto actualizado exitosamente"})
}
