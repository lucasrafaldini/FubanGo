package apidesign

import (
	"encoding/json"
	"net/http"
)

// DTOs e separação de responsabilidades
type UserDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

// Handler fino que delega lógica
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// Aqui deveríamos extrair ID, validar, autenticar, etc.
	user := UserDTO{ID: 1, Name: "João"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Função de serviço separada
func GetUserService(id int) (UserDTO, error) {
	// Simula busca por usuário
	return UserDTO{ID: id, Name: "João"}, nil
}
