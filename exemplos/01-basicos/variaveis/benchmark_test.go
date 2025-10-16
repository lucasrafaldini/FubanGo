package variaveis

import (
	"testing"
)

func BenchmarkBadVariableDeclarations(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string = "texto"
		var num int = 42
		var slice []int = []int{}
		var m map[string]int = map[string]int{}

		_ = str
		_ = num
		_ = slice
		_ = m
	}
}

func BenchmarkGoodVariableDeclarations(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str := "texto"
		num := 42
		slice := make([]int, 0, 10)
		m := make(map[string]int)

		_ = str
		_ = num
		_ = slice
		_ = m
	}
}
