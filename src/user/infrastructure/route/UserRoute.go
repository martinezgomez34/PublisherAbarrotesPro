package route

import (
	"publisher/src/core"
	"publisher/src/user/application"
	repositories "publisher/src/user/infrastructure"
	controllers "publisher/src/user/infrastructure/controller"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(r *mux.Router, rabbitMQ *core.RabbitMQ) {
	repo := &repositories.RabbitMQRepository{rabbitMQ}
	authUseCase := application.NewAuthUseCase(repo)
	userUseCase := application.NewUserUseCase(repo)

	authController := controllers.NewAuthController(authUseCase)
	userController := controllers.NewUserController(userUseCase)

	r.HandleFunc("/login", authController.LoginHandler).Methods("POST")
	r.HandleFunc("/register", userController.CreateUserHandler).Methods("POST")
}