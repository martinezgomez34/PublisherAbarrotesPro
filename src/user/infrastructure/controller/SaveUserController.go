package controllers

import (
	"encoding/json"
	"net/http"
	"publisher/src/user/application"
	"publisher/src/user/domain"
	"github.com/google/uuid"
)

type UserController struct {
	UseCase *application.UserUseCase
}

func NewUserController(useCase *application.UserUseCase) *UserController {
	return &UserController{UseCase: useCase}
}


func (uc *UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Error al leer el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	user := domain.User{
		ID:       uuid.New().String(), 
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password, 
	}

	message := domain.Message{
		Type: domain.MessageTypeRegister,
		User: user,
	}

	if err := uc.UseCase.CreateUser(message); err != nil {
		http.Error(w, "Error al registrar el usuario", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuario registrado exitosamente"})
}
