package routes

import (
	"publisher/src/core"
	"publisher/src/product/application"
	repositories "publisher/src/product/infrastructure"
	"publisher/src/product/infrastructure/controller"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine, rabbitMQ *core.RabbitMQ, db *core.Database) {
    rabbitRepo := &repositories.RabbitMQRepository{rabbitMQ}
    dbRepo := repositories.NewProductRepository(db.DB)

    productUseCase := application.NewProductUseCase(rabbitRepo)
    editUseCase := application.NewEditProductUseCase(rabbitRepo)
    deleteUseCase := application.NewDeleteProductUseCase(rabbitRepo)
    listUseCase := application.NewListProductUseCase(dbRepo)
    getByIDUseCase := application.NewGetByIDProductUseCase(dbRepo)

    productCtrl := controllers.NewProductController(productUseCase, dbRepo)
    editCtrl := controllers.NewEditProductController(editUseCase, dbRepo)
    deleteCtrl := controllers.NewDeleteProductController(deleteUseCase, dbRepo)
    listCtrl := controllers.NewListProductController(listUseCase)
    getByIDCtrl := controllers.NewGetByIDProductController(getByIDUseCase)

    productGroup := router.Group("/products")
    {
        productGroup.POST("", productCtrl.CreateProductHandler)
        productGroup.PUT("/:id", editCtrl.UpdateProductHandler)
        productGroup.DELETE("/:id", deleteCtrl.DeleteProductHandler)
        productGroup.GET("/", listCtrl.GetAllProducts)
        productGroup.GET("/:id", getByIDCtrl.GetProductByID)
    }
}