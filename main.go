package main

import (
    "log"
    "net/http"
    "github.com/joho/godotenv"
    "publisher/src/core"
    userRoute "publisher/src/user/infrastructure/route" 
    productRoute "publisher/src/product/infrastructure/route" 
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error cargando el archivo .env")
    }
   
    productQueue := core.NewRabbitMQ("product_queue")
	defer productQueue.Close()

	userQueue := core.NewRabbitMQ("user_queue")
	defer userQueue.Close()

   
    r := mux.NewRouter()
    userRoute.SetupUserRoutes(r, userQueue) 
    productRoute.SetupProductRoutes(r, productQueue) 

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:4200"}), 
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), 
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    http.Handle("/", corsHandler(r))

    log.Println("Servidor corriendo en :8081")
    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal("Error al iniciar el servidor: ", err)
    }
}