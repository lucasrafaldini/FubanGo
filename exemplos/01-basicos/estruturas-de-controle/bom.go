package controlstructures

import (
	"fmt"
)

// Exemplo de boas práticas com estruturas de controle
func GoodControlStructures() {
	x := 10

	// Early return e condições combinadas
	if x <= 0 || x >= 20 || x%2 != 0 || x <= 5 {
		return
	}
	fmt.Println("x é par, maior que 5 e menor que 20")

	// Switch com default e casos bem definidos
	switch {
	case x < 5:
		fmt.Println("menor que 5")
	case x < 10:
		fmt.Println("entre 5 e 9")
	case x == 10:
		fmt.Println("igual a 10")
	default:
		fmt.Println("maior que 10")
	}

	// For limpo e direto
	for i := 0; i < 10; i++ {
		if i%2 != 0 && i <= 5 {
			fmt.Println(i)
		}
	}

	// For com condição clara
	for counter := 0; counter < 10; counter++ {
		// processamento
	}

	// Range idiomático
	numbers := []int{1, 2, 3, 4, 5}
	for i := range numbers {
		numbers[i] *= 2
	}

	// Condições complexas quebradas em partes lógicas
	a, b, c := 1, 2, 3

	isRatioValid := (float64(a+b)*float64(c))/(float64(a*b)) > 1
	isSumEven := (a+b+c)%2 == 0
	isProductOdd := (a*b*c)%2 == 1

	if isRatioValid && (isSumEven || isProductOdd) {
		fmt.Println("condição satisfeita")
	}
}
