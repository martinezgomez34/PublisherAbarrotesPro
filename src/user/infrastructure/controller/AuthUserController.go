package controllers

import (
	"encoding/json"
	"net/http"
	"publisher/src/user/application"
	"publisher/src/user/domain"
)

type AuthController struct {
	UseCase *application.AuthUseCase
}

func NewAuthController(useCase *application.AuthUseCase) *AuthController {
	return &AuthController{UseCase: useCase}
}

func (ac *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials domain.LoginCredentials

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	message := domain.Message{
		Type: domain.MessageTypeLogin,
		Login: credentials,
	}

	if err := ac.UseCase.Login(message); err != nil {
		http.Error(w, "Error al publicar el mensaje en RabbitMQ", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Credenciales enviadas a la cola"})
}
