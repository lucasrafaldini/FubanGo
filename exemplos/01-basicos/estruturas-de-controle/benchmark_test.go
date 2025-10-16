package controlstructures

import (
	"testing"
)

// Benchmark para estruturas de controle ruins
func BenchmarkBadControlStructures(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := 10
		// If's aninhados
		if x > 0 {
			if x < 20 {
				if x%2 == 0 {
					if x > 5 {
						_ = x
					}
				}
			}
		}

		// Loop com continue/break
		for j := 0; j < 10; j++ {
			if j%2 == 0 {
				continue
			} else {
				if j > 5 {
					break
				}
			}
		}
	}
}

// Benchmark para estruturas de controle boas
func BenchmarkGoodControlStructures(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := 10
		// Condição simplificada
		if x > 0 && x < 20 && x%2 == 0 && x > 5 {
			_ = x
		}

		// Loop simplificado
		for j := 0; j < 10 && j <= 5; j++ {
			if j%2 != 0 {
				_ = j
			}
		}
	}
}
