package errorhandling

import (
	"testing"
)

// Benchmark da implementação ruim de leitura de arquivo
func BenchmarkBadFileRead(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IgnoreAllErrors()
	}
}

// Benchmark da implementação boa de leitura de arquivo
func BenchmarkGoodFileRead(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SafeFileRead("arquivo.txt")
	}
}

// Benchmark de tratamento ruim de erro com panic
func BenchmarkBadErrorHandling(b *testing.B) {
	defer func() {
		recover() // Recupera de todos os panics para não quebrar o benchmark
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PanicInsteadOfError(-1)
	}
}

// Benchmark de tratamento bom de erro
func BenchmarkGoodErrorHandling(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ValidatePositive(-1)
	}
}

// Benchmark de recuperação ruim de panic
func BenchmarkBadPanicRecovery(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RecoverEverything()
	}
}

// Benchmark de recuperação boa de panic
func BenchmarkGoodPanicRecovery(b *testing.B) {
	handler := func() {}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SafeRecover(handler)
	}
}
