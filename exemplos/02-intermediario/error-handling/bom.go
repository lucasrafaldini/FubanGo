package errorhandling

import (
	"errors"
	"fmt"
	"os"
)

// ConfigErrorType define tipos específicos de erros de configuração
type ConfigErrorType int

const (
	PermissionError ConfigErrorType = iota
	ReadError
	ParseError
)

// ConfigError é um tipo de erro personalizado com contexto
type ConfigError struct {
	Type    ConfigErrorType
	Message string
	Err     error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// NewConfigError cria um novo erro de configuração com contexto
func NewConfigError(errType ConfigErrorType, message string, err error) *ConfigError {
	return &ConfigError{
		Type:    errType,
		Message: message,
		Err:     err,
	}
}

// SafeFileRead lê arquivo com tratamento apropriado de erros
func SafeFileRead(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("falha ao abrir arquivo %s: %w", filename, err)
	}
	defer file.Close()

	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		return nil, fmt.Errorf("falha ao ler arquivo %s: %w", filename, err)
	}

	return data[:n], nil
}

// ValidatePositive valida valor com erro apropriado
func ValidatePositive(value int) (int, error) {
	if value < 0 {
		return 0, fmt.Errorf("valor %d é negativo", value)
	}
	return value * 2, nil
}

// LoadConfig carrega configuração com hierarquia de erros
func LoadConfig() error {
	if err := checkConfigPermissions(); err != nil {
		return NewConfigError(PermissionError, "erro de permissão", err)
	}

	if err := readConfigFile(); err != nil {
		return NewConfigError(ReadError, "erro de leitura", err)
	}

	if err := parseConfigContent(); err != nil {
		return NewConfigError(ParseError, "erro de parse", err)
	}

	return nil
}

func checkConfigPermissions() error {
	// Simulação de verificação de permissões
	return nil
}

func readConfigFile() error {
	// Simulação de leitura de arquivo
	return nil
}

func parseConfigContent() error {
	// Simulação de parse de conteúdo
	return nil
}

// SafeDivide demonstra tratamento de erro com logging estruturado
func SafeDivide(a, b int) (result int, err error) {
	// Defer para logging centralizado em caso de erro
	defer func() {
		if err != nil {
			LogError("divisão falhou", "a", a, "b", b, "erro", err)
		}
	}()

	if b == 0 {
		return 0, errors.New("divisão por zero não permitida")
	}

	return a / b, nil
}

// LogError centraliza logging de erros
func LogError(message string, keyvals ...interface{}) {
	// Em produção, usar um logger estruturado real
	fmt.Printf("ERROR: %s | ", message)
	for i := 0; i < len(keyvals); i += 2 {
		if i+1 < len(keyvals) {
			fmt.Printf("%v: %v ", keyvals[i], keyvals[i+1])
		}
	}
	fmt.Println()
}

// SafeRecover recupera apenas de panics específicos
func SafeRecover(handler func()) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				if x == "erro esperado" {
					LogError("recuperado de erro esperado", "panic", x)
					handler()
				} else {
					// Re-panic para erros não esperados
					panic(r)
				}
			default:
				// Re-panic para tipos não esperados
				panic(r)
			}
		}
	}()
}
