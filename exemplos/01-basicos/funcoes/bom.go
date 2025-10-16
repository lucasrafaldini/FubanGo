package funcoes

import (
	"fmt"
	"strings"
)

// UserData agrupa dados relacionados
type UserData struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Address   Address
	Flags     UserFlags
}

// Address agrupa dados de endereço
type Address struct {
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}

// UserFlags agrupa flags booleanas
type UserFlags struct {
	IsActive  bool
	IsAdmin   bool
	IsPremium bool
}

// ProcessResult encapsula múltiplos retornos
type ProcessResult struct {
	FullName    string
	FullAddress string
	Age         int
	IsValid     bool
	Error       error
}

// ProcessUser processa dados do usuário de forma organizada
func ProcessUser(data UserData, age int) ProcessResult {
	if err := validateUserData(data); err != nil {
		return ProcessResult{Error: err}
	}

	return ProcessResult{
		FullName:    formatFullName(data),
		FullAddress: formatAddress(data.Address),
		Age:         age,
		IsValid:     true,
	}
}

// validateUserData valida os dados do usuário
func validateUserData(data UserData) error {
	if data.FirstName == "" || data.LastName == "" {
		return fmt.Errorf("nome inválido")
	}
	return nil
}

// formatFullName formata o nome completo com status
func formatFullName(data UserData) string {
	var parts []string
	parts = append(parts, data.FirstName+" "+data.LastName)

	// Map de status para evitar repetição
	status := map[string]bool{
		"Ativo":   data.Flags.IsActive,
		"Admin":   data.Flags.IsAdmin,
		"Premium": data.Flags.IsPremium,
	}

	for label, isSet := range status {
		if isSet {
			parts = append(parts, "("+label+")")
		}
	}

	return strings.Join(parts, " ")
}

// formatAddress formata o endereço completo
func formatAddress(addr Address) string {
	return fmt.Sprintf("%s, %s - %s, %s (%s)",
		addr.Street,
		addr.City,
		addr.State,
		addr.Country,
		addr.ZipCode,
	)
}

// ProcessItems processa items com responsabilidades separadas
func ProcessItems(items []string) (ProcessedItems, error) {
	processed := ProcessedItems{
		Items: make([]string, len(items)),
		Stats: ItemStats{},
	}

	for i, item := range items {
		if item == "" {
			return processed, fmt.Errorf("item vazio encontrado no índice %d", i)
		}
		processed.Items[i] = strings.ToUpper(item)
		processed.Stats.Total += len(item)
	}

	return processed, nil
}

// ProcessedItems encapsula resultado do processamento
type ProcessedItems struct {
	Items []string
	Stats ItemStats
}

// ItemStats mantém estatísticas do processamento
type ItemStats struct {
	Total int
}

// SumRecursive implementa recursão de forma segura
func SumRecursive(n int) (int, error) {
	// Validação de entrada
	if n < 0 {
		return 0, fmt.Errorf("número deve ser não-negativo")
	}

	// Caso base explícito
	if n == 0 {
		return 0, nil
	}

	// Chamada recursiva
	sum, err := SumRecursive(n - 1)
	if err != nil {
		return 0, err
	}

	return n + sum, nil
}

// ParseNumber trata erros apropriadamente
func ParseNumber(input string) (int, error) {
	var num int
	_, err := fmt.Sscanf(input, "%d", &num)
	if err != nil {
		return 0, fmt.Errorf("erro ao converter número: %w", err)
	}
	return num, nil
}
