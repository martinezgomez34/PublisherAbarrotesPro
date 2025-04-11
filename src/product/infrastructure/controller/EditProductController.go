package controllers

import (
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
	"github.com/gin-gonic/gin"
)

type EditProductController struct {
	UseCase *application.EditProductUseCase
	Repo    domain.IProductRepository
}

func NewEditProductController(useCase *application.EditProductUseCase, repo domain.IProductRepository) *EditProductController {
	return &EditProductController{
		UseCase: useCase,
		Repo:    repo,
	}
}

func (ec *EditProductController) UpdateProductHandler(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar si el producto existe
	if _, err := ec.Repo.GetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	product := domain.Product{
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
	}

	if err := ec.Repo.EditProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el producto"})
		return
	}

	message := domain.Message{
		Type:    domain.MessageTypeNotification,
		Product: product,
	}

	if err := ec.UseCase.UpdateProduct(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar notificación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado exitosamente"})
}