package controllers

import (
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	UseCase *application.DeleteProductUseCase
	Repo    domain.IProductRepository
}

func NewDeleteProductController(useCase *application.DeleteProductUseCase, repo domain.IProductRepository) *DeleteProductController {
	return &DeleteProductController{
		UseCase: useCase,
		Repo:    repo,
	}
}

func (dc *DeleteProductController) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")

	// Verificar si el producto existe
	if _, err := dc.Repo.GetByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	if err := dc.Repo.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el producto"})
		return
	}

	message := domain.Message{
		Type: domain.MessageTypeNotification,
		Product: domain.Product{
		},
	}

	if err := dc.UseCase.DeleteProduct(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar notificación"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado exitosamente"})
}