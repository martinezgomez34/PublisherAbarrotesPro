package main

import (
	"log"
	"publisher/src/core"
	"time"

	"github.com/joho/godotenv"
	productRoute "publisher/src/product/infrastructure/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	productQueue := core.NewRabbitMQ("product_messages")
	defer productQueue.Close()

	userQueue := core.NewRabbitMQ("user_messages")
	defer userQueue.Close()

	db := core.NewDatabase()
	defer db.Close()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	productRoute.SetupProductRoutes(router, productQueue, db)

	log.Println("Servidor corriendo en :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}