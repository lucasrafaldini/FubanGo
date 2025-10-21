package errorhandling

import (
	"fmt"
	"os"
)

// Erro genérico reutilizado
var ErrGeneric = fmt.Errorf("algo deu errado")

// Função que ignora erros completamente
func IgnoreAllErrors() string {
	file, _ := os.Open("arquivo.txt")
	defer file.Close()

	data := make([]byte, 100)
	_, _ = file.Read(data)

	return string(data)
}

// Função que usa panic ao invés de retornar erros
func PanicInsteadOfError(value int) int {
	if value < 0 {
		panic("valor negativo não permitido")
	}
	return value * 2
}

// Função que retorna apenas mensagem de erro sem contexto
func ReturnGenericError() error {
	return fmt.Errorf("erro")
}

// Função que perde informação do erro original
func LoseErrorContext(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("falha ao verificar arquivo")
	}
	return nil
}

// Função que mistura retorno de erro com logs
func MixErrorAndLogging(value int) (int, error) {
	if value == 0 {
		fmt.Println("Erro: divisão por zero")
		return 0, fmt.Errorf("divisão por zero")
	}

	result := 100 / value
	fmt.Printf("Resultado: %d\n", result)
	return result, nil
}

// Função que não agrupa erros relacionados
type BadConfigError struct {
	Msg string
}

func (e BadConfigError) Error() string {
	return e.Msg
}

func BadLoadConfig() error {
	// Erros não agrupados e sem hierarquia
	if err := checkPermissions(); err != nil {
		return fmt.Errorf("erro de permissão: %v", err)
	}
	if err := readConfig(); err != nil {
		return fmt.Errorf("erro de leitura: %v", err)
	}
	if err := parseConfig(); err != nil {
		return fmt.Errorf("erro de parse: %v", err)
	}
	return nil
}

func checkPermissions() error {
	return BadConfigError{"sem permissão"}
}

func readConfig() error {
	return BadConfigError{"arquivo não encontrado"}
}

func parseConfig() error {
	return BadConfigError{"formato inválido"}
}

// Função que tenta recuperar de todos os panics indiscriminadamente
func RecoverEverything() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de:", r)
		}
	}()

	// Qualquer panic será recuperado, mesmo os que não deveriam
	panic("erro crítico que deveria derrubar a aplicação")
}
