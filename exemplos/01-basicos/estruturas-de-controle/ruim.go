package controlstructures

import (
	"fmt"
)

// Exemplo de estruturas de controle implementadas da pior maneira possível
func BadControlStructures() {
	// If aninhados excessivamente
	x := 10
	if x > 0 {
		if x < 20 {
			if x%2 == 0 {
				if x > 5 {
					fmt.Println("x é par, maior que 5 e menor que 20")
				}
			}
		}
	}

	// Switch sem default e casos redundantes
	switch x {
	case 10:
		fmt.Println("é 10")
	case 5 + 5: // caso redundante e confuso (exemplo ruim)
		fmt.Println("este caso é redundante e confuso")
	case 3 + 7: // outro caso redundante
		fmt.Println("outro caso redundante")
	}

	// For com continue/break desnecessários
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		} else {
			if i > 5 {
				break
			} else {
				fmt.Println(i)
			}
		}
	}

	// Loop infinito com break baseado em condição
	counter := 0
	for {
		counter++
		if counter >= 10 {
			break
		}
	}

	// Range sobre slice com índice não utilizado
	numbers := []int{1, 2, 3, 4, 5}
	for i, _ := range numbers {
		numbers[i] = numbers[i] * 2
	}

	// Condições complexas e difíceis de entender
	a, b, c := 1, 2, 3
	if ((a+b)*c)/(a*b) > 1 && (a+b+c)%2 == 0 || (a*b*c)%2 == 1 {
		fmt.Println("condição complexa satisfeita")
	}
}
