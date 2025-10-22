package rwtesting

import (
	"testing"
	"time"
)

// Nota: Benchmarks de testes são meta - estamos benchmarking funções de teste
// O foco aqui é demonstrar overhead de diferentes abordagens

// Benchmark: Testes dependentes de ordem (com estado global)
func BenchmarkTestDependsOnOrder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		globalCounter = 0
		globalCounter++
	}
}

// Benchmark: Testes isolados (sem estado global)
func BenchmarkTestIsolated(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		localCounter := 0
		localCounter++
	}
}

// Benchmark: Teste lento com sleep (sem mocks)
func BenchmarkSlowTestWithSleep(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Millisecond) // Simula chamada lenta
	}
}

// Benchmark: Teste rápido com mock
func BenchmarkFastTestWithMock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Mock retorna imediatamente
		_ = "mocked result"
	}
}

// Benchmark: Sleep ao invés de sincronização
func BenchmarkBadSleepSync(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		done := make(chan bool, 1)
		go func() {
			done <- true
		}()
		time.Sleep(10 * time.Millisecond) // Overhead desnecessário
	}
}

// Benchmark: Sincronização adequada com canal
func BenchmarkGoodChannelSync(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		done := make(chan bool, 1)
		go func() {
			done <- true
		}()
		<-done // Espera exatamente o necessário
	}
}

// Benchmark: Sem assertions adequadas
func BenchmarkNoAssertions(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := sum(2, 2)
		_ = result // Não valida
	}
}

// Benchmark: Com assertions
func BenchmarkWithAssertions(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := sum(2, 2)
		if result != 4 {
			b.Fatal("assertion failed")
		}
	}
}

// Benchmark: Setup duplicado em cada teste
func BenchmarkDuplicatedSetup(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Setup duplicado
		db := openDatabase()
		insertTestData(db)
		// teste...
	}
}

// Benchmark: Setup centralizado
func BenchmarkCentralizedSetup(b *testing.B) {
	// Setup uma vez antes do loop
	db := openDatabase()
	insertTestData(db)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// apenas teste
		_ = db
	}
}

// Benchmark: Múltiplos testes separados (sem table-driven)
func BenchmarkSeparateTests(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if sum(1, 1) != 2 {
			b.Fatal("1+1 failed")
		}
		if sum(2, 2) != 4 {
			b.Fatal("2+2 failed")
		}
		if sum(5, 3) != 8 {
			b.Fatal("5+3 failed")
		}
	}
}

// Benchmark: Table-driven test
func BenchmarkTableDriven(b *testing.B) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{5, 3, 8},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if sum(tt.a, tt.b) != tt.expected {
				b.Fatal("test failed")
			}
		}
	}
}

// Benchmark: Acesso concorrente sem proteção (race condition)
func BenchmarkConcurrentAccessUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := 0
		done := make(chan bool)

		for j := 0; j < 10; j++ {
			go func() {
				counter++ // race condition
				done <- true
			}()
		}

		for j := 0; j < 10; j++ {
			<-done
		}
	}
}

// Benchmark: Acesso concorrente com proteção
func BenchmarkConcurrentAccessSafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter := 0
		done := make(chan bool)
		updates := make(chan int, 10)

		// Goroutine para atualizar counter de forma segura
		go func() {
			for range updates {
				counter++
			}
			close(done)
		}()

		for j := 0; j < 10; j++ {
			go func() {
				updates <- 1
			}()
		}

		close(updates)
		<-done
	}
}

// Benchmark: Apenas happy path
func BenchmarkOnlyHappyPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := divide(10, 2)
		if result != 5 {
			b.Fatal("division failed")
		}
	}
}

// Benchmark: Testa também condições de erro
func BenchmarkWithErrorCases(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Happy path
		result := divide(10, 2)
		if result != 5 {
			b.Fatal("division failed")
		}

		// Error case
		result = divide(10, 0)
		if result != 0 {
			b.Fatal("division by zero should return 0")
		}
	}
}

// Benchmark: Mock mal implementado (sempre retorna o mesmo)
func BenchmarkBadMock(b *testing.B) {
	mock := &BadMock{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result1 := mock.Method2(1)
		result2 := mock.Method2(999)
		_ = result1
		_ = result2
	}
}

// Benchmark: Mock bem implementado (valida inputs)
func BenchmarkGoodMock(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Mock que valida argumentos e retorna valores específicos
		inputs := []int{1, 999}
		for _, input := range inputs {
			if input == 1 {
				_ = "result for 1"
			} else if input == 999 {
				_ = "result for 999"
			}
		}
	}
}

// Benchmark: Estado compartilhado entre testes
func BenchmarkSharedState(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sharedResource["key"] = "value1"
		_ = sharedResource["key"]

		sharedResource["key"] = "value2"
		_ = sharedResource["key"]
	}
}

// Benchmark: Estado isolado
func BenchmarkIsolatedState(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Cada iteração usa seu próprio estado
		localResource := map[string]string{}

		localResource["key"] = "value1"
		_ = localResource["key"]

		localResource["key"] = "value2"
		_ = localResource["key"]
	}
}
