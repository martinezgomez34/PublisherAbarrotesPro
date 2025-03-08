package routes

import (
	"publisher/src/core"
	"publisher/src/product/application"
	"publisher/src/product/infrastructure/controller"
	"publisher/src/product/infrastructure"

	"github.com/gorilla/mux"
)

func SetupProductRoutes(r *mux.Router, rabbitMQ *core.RabbitMQ) {
	repo := &repositories.RabbitMQRepository{rabbitMQ}
	productUseCase := application.NewProductUseCase(repo)
	EditproductUseCase := application.NewEditProductUseCase(repo)
	DeleteProductUseCase := application.NewDeleteProductUseCase(repo)

	productController := controllers.NewProductController(productUseCase)
	EditproductController := controllers.NewEditProductController(EditproductUseCase)
	DeleteproductController := controllers.NewDeleteProductController(DeleteProductUseCase)

	r.HandleFunc("/products", productController.CreateProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", EditproductController.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", DeleteproductController.DeleteProductHandler).Methods("DELETE")
}