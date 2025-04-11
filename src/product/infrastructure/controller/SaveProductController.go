// controllers/product_controller.go
package controllers

import (
	"net/http"
	"publisher/src/product/application"
	"publisher/src/product/domain"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	UseCase *application.ProductUseCase
	Repo    domain.IProductRepository
}

func NewProductController(useCase *application.ProductUseCase, repo domain.IProductRepository) *ProductController {
	return &ProductController{
		UseCase: useCase,
		Repo:    repo,
	}
}

func (pc *ProductController) CreateProductHandler(c *gin.Context) {
	var newProduct struct {
		Name        string  `json:"name" binding:"required"`
		Description string  `json:"description"`
		Price       float64 `json:"price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := domain.Product{
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Price:       newProduct.Price,
	}

	if err := pc.Repo.SaveProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el producto en la BD"})
		return
	}

	message := domain.Message{
		Type:    domain.MessageTypeNotification,
		Product: product,
	}

	if err := pc.UseCase.CreateProduct(message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar notificación"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Producto creado exitosamente",
		"id":      product.ID,
	})
}