package controllers

import (
	"encoding/json"
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
)

type DeleteProductController struct {
	UseCase *application.DeleteProductUseCase
}

func NewDeleteProductController(useCase *application.DeleteProductUseCase) *DeleteProductController {
	return &DeleteProductController{UseCase: useCase}
}

func (pc *DeleteProductController) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	var deleteRequest struct {
		ID string `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deleteRequest); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	product := domain.Product{
		ID: deleteRequest.ID,
	}

	message := domain.Message{
		Type:    domain.MessageTypeDeleteProduct,
		Product: product,
	}

	if err := pc.UseCase.DeleteProduct(message); err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Producto eliminado exitosamente"})
}
