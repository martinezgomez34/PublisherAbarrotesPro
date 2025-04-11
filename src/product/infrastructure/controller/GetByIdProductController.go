package controllers

import (
	"publisher/src/product/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetByIdProductController struct {
	service *application.GetByIdProductUseCase
}

func NewGetByIDProductController(service *application.GetByIdProductUseCase) *GetByIdProductController {
	return &GetByIdProductController{service: service}
}

func (pc *GetByIdProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}
