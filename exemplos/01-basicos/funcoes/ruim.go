package funcoes

import (
	"fmt"
	"strings"
)

// Variável global desnecessária
var result string

// Função com muitos parâmetros e retornos confusos
func ProcessUserData(firstName, lastName, email, phone, address, city, state, country, zipCode string, age int, isActive, isAdmin, isPremium bool) (string, string, error, bool, int) {
	if firstName == "" || lastName == "" {
		return "", "", fmt.Errorf("invalid name"), false, -1
	}

	// Modificando variável global
	result = firstName + " " + lastName

	// Código repetitivo
	if isActive {
		result = result + " (Ativo)"
	}
	if isAdmin {
		result = result + " (Admin)"
	}
	if isPremium {
		result = result + " (Premium)"
	}

	// Função anônima desnecessariamente complexa
	processAddress := func(addr, city, state, country, zip string) string {
		var sb strings.Builder
		sb.WriteString(addr)
		sb.WriteString(", ")
		sb.WriteString(city)
		sb.WriteString(" - ")
		sb.WriteString(state)
		sb.WriteString(", ")
		sb.WriteString(country)
		sb.WriteString(" (")
		sb.WriteString(zip)
		sb.WriteString(")")
		return sb.String()
	}

	fullAddress := processAddress(address, city, state, country, zipCode)

	// Retornos múltiplos sem nomes
	return result, fullAddress, nil, true, age * 2
}

// Função que faz muitas coisas diferentes
func DoEverything(data []string) {
	// Processamento de dados
	for i := 0; i < len(data); i++ {
		data[i] = strings.ToUpper(data[i])
	}

	// Validação
	for _, item := range data {
		if item == "" {
			panic("item vazio encontrado")
		}
	}

	// Logging
	for _, item := range data {
		fmt.Printf("Processado: %s\n", item)
	}

	// Cálculos
	total := 0
	for _, item := range data {
		total += len(item)
	}

	// Atualização de estado global
	result = fmt.Sprintf("Total processado: %d", total)
}

// Função recursiva mal implementada
func BadRecursion(n int) int {
	// Sem caso base explícito
	if n > 0 {
		return n + BadRecursion(n-1)
	}
	return 0
}

// Função que ignora erros
func IgnoreErrors(input string) string {
	// Ignora erros de conversão
	num, _ := fmt.Sscanf(input, "%d")
	return fmt.Sprintf("Número: %v", num)
}
