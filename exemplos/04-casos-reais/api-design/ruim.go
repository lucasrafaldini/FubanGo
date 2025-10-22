package apidesign

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Side-effects em endpoint GET
func BadGetEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// RUIM: GET está deletando dados
		deleteAllUsers()
		fmt.Fprintf(w, "All users deleted")
	}
}

func deleteAllUsers() {
	// simula deleção
}

// 2. Vazamento de dados sensíveis
func LeakSensitiveData(w http.ResponseWriter, r *http.Request) {
	// RUIM: retorna dados sensíveis sem filtro
	user := struct {
		ID       int
		Name     string
		Email    string
		Password string // RUIM: senha em texto claro na resposta
		SSN      string // RUIM: dado sensível exposto
		Token    string // RUIM: token de autenticação exposto
	}{
		ID:       1,
		Name:     "John",
		Email:    "john@example.com",
		Password: "secret123",
		SSN:      "123-45-6789",
		Token:    "jwt-token-here",
	}

	// Retorna HTML ao invés de JSON
	fmt.Fprintf(w, "<html><body>User: %+v</body></html>", user)
}

// 3. Falta de autenticação e autorização
func DeleteAllWithoutAuth(w http.ResponseWriter, r *http.Request) {
	// RUIM: endpoint destrutivo sem verificação de autenticação
	// Qualquer um pode deletar todos os dados
	fmt.Fprintf(w, "Deleting all data...")
	// deleta tudo sem verificar quem está fazendo a requisição
}

// 4. Mistura de responsabilidades - handler com lógica de negócio
func BadHandler(w http.ResponseWriter, r *http.Request) {
	// RUIM: handler faz parsing, validação, lógica de negócio e acesso a DB
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)

	// Validação no handler
	if input["name"] == nil {
		w.WriteHeader(400)
		return
	}

	// Lógica de negócio no handler
	name := input["name"].(string)
	processedName := processBusinessLogic(name)

	// Acesso direto a DB no handler
	db, _ := sql.Open("postgres", "conn-string")
	defer db.Close()
	db.Exec("INSERT INTO users(name) VALUES(?)", processedName)

	// Resposta direto no handler
	fmt.Fprintf(w, "User created")
}

func processBusinessLogic(name string) string {
	return name
}

// 5. Falta de versionamento
func NoVersioning(w http.ResponseWriter, r *http.Request) {
	// RUIM: endpoint sem versão na URL ou header
	// Mudanças na API quebrarão clientes existentes
	fmt.Fprintf(w, "Response without version")
}

// 6. Tipos genéricos sem contratos claros
func GenericTypes(req interface{}) interface{} {
	// RUIM: usa interface{} perdendo type safety
	// Não há contrato claro de entrada/saída
	return nil
}

// 7. Tratamento de erros inconsistente
func InconsistentErrors(w http.ResponseWriter, r *http.Request) {
	// Erro 1: retorna texto simples
	if r.Header.Get("Auth") == "" {
		fmt.Fprintf(w, "Error: not authorized")
		return
	}

	// Erro 2: retorna JSON
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
		return
	}

	// Erro 3: retorna status code diferente para mesmo tipo de erro
	if r.Header.Get("Token") == "" {
		w.WriteHeader(403) // deveria ser 401
		return
	}
}

// 8. Não validar entrada
func NoInputValidation(w http.ResponseWriter, r *http.Request) {
	var input map[string]interface{}
	json.NewDecoder(r.Body).Decode(&input)

	// RUIM: usa input diretamente sem validar
	// Pode causar panic se campos esperados não existirem
	name := input["name"].(string)
	age := input["age"].(int)

	fmt.Fprintf(w, "Name: %s, Age: %d", name, age)
}

// 9. Endpoints que fazem demais
func GodEndpoint(w http.ResponseWriter, r *http.Request) {
	// RUIM: um único endpoint faz múltiplas operações diferentes
	action := r.URL.Query().Get("action")

	if action == "create" {
		// cria usuário
	} else if action == "update" {
		// atualiza usuário
	} else if action == "delete" {
		// deleta usuário
	} else if action == "list" {
		// lista usuários
	} else if action == "export" {
		// exporta para CSV
	} else if action == "import" {
		// importa de CSV
	}
	// ... muitas outras ações
}

// 10. Não usar códigos HTTP apropriados
func WrongStatusCodes(w http.ResponseWriter, r *http.Request) {
	// RUIM: sempre retorna 200 mesmo em erro
	var input map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(200) // RUIM: deveria ser 400
		fmt.Fprintf(w, "Invalid JSON")
		return
	}

	// Recurso não encontrado
	if input["id"] == nil {
		w.WriteHeader(200) // RUIM: deveria ser 404
		fmt.Fprintf(w, "Not found")
		return
	}
}
