package variaveis

// Constantes nomeadas para valores significativos
const (
	MinimumAge    = 18
	MaxNameLength = 50
)

// User agrupa dados relacionados em uma struct
type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// Exemplo de boas práticas com variáveis
func GoodVariableExample() {
	// Uso de struct para agrupar dados relacionados
	user := User{
		FirstName: "João",
		LastName:  "Silva",
		Age:       25,
		Email:     "joao@example.com",
	}

	// Uso de inferência de tipo
	count := 0
	isActive := true

	// Inicialização eficiente de slices
	numbers := make([]int, 0, 10)

	// Inicialização de map com make
	cache := make(map[string]int)

	// Uso de constantes para valores significativos
	if user.Age < MinimumAge {
		println("Usuário menor de idade")
		return
	}

	// Escopo limitado para variáveis temporárias
	if name := user.FirstName; len(name) > MaxNameLength {
		println("Nome muito longo")
		return
	}

	// Uso apropriado das variáveis declaradas
	if isActive {
		cache["user_count"] = count
		numbers = append(numbers, user.Age)
	}
}
