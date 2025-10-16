//go:build ignore

package variaveis

// Este arquivo contém o exemplo "ruim" de variáveis. Ele é intencionalmente
// excluído da compilação normal usando a build tag `ignore`, para que o
// repositório possa ser usado para ensino sem que exemplos malformados que
// declaram variáveis não utilizadas quebrem `go test`.

// Exemplo de declaração e uso de variáveis da pior maneira possível
func BadVariableExample() {
	// Declaração de variáveis com nomes não descritivos
	var x string
	var y int
	var z bool

	// Não aproveitando inferência de tipo
	var str string = "texto"
	var num int = 42
	var flag bool = false

	// Declarações redundantes
	var slice []int = []int{}
	var m map[string]int = map[string]int{}

	// Variáveis não utilizadas
	temporary := "isso não será usado"

	// Escopo desnecessariamente global
	GlobalVar := "não deveria ser global"

	// Conversões desnecessárias
	x = string("abc")
	y = int(42)
	z = bool(true)

	// Shadowing de variáveis
	if true {
		str := "outro texto"
		println(str) // shadowing a variável externa
	}

	// Uso de valores mágicos sem constantes
	if num > 42 {
		println("maior que a resposta para a vida, o universo e tudo mais")
	}

	// Falta de agrupamento lógico de variáveis relacionadas
	userFirstName := "João"
	userAge := 25
	userLastName := "Silva"
	userEmail := "joao@example.com"
}
