package interfaces

import (
	"testing"
)

// Benchmark de interface grande
func BenchmarkBigInterface(b *testing.B) {
	impl := &BigImplementation{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		impl.ReadData()
		impl.ProcessData()
		impl.WriteData([]byte("test"))
	}
}

// Benchmark de interfaces pequenas e focadas
func BenchmarkSmallInterfaces(b *testing.B) {
	handler := &FileHandler{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, _ := handler.Read()
		processed, _ := handler.Process(data)
		handler.Write(processed)
	}
}

// Benchmark de container com interface{}
func BenchmarkBadContainer(b *testing.B) {
	container := &BadContainer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container.Store("test")
		_ = container.Retrieve()
	}
}

// Benchmark de container com generics
func BenchmarkGoodContainer(b *testing.B) {
	container := NewContainer("test")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		container.Store("new test")
		_ = container.Retrieve()
	}
}

// Benchmark de type assertions excessivas
func BenchmarkBadTypeAssertions(b *testing.B) {
	asserter := &BadTypeAssert{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		asserter.ProcessAnything("test")
		asserter.ProcessAnything(42)
		asserter.ProcessAnything(3.14)
	}
}

// Benchmark de processador com tipo específico
func BenchmarkGoodTypeSpecific(b *testing.B) {
	processor := &StringProcessor{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = processor.Process("test")
	}
}
