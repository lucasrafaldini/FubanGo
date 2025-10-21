package concorrencia

import (
	"context"
	"testing"
	"time"
)

// Benchmark de contador com race condition
func BenchmarkBadCounter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BadConcurrentCounter()
	}
}

// Benchmark de contador thread-safe
func BenchmarkGoodCounter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SafeConcurrentCounter()
	}
}

// Benchmark de goroutine leak
func BenchmarkBadGoroutineLeak(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BadGoroutineLeak()
	}
}

// Benchmark de goroutine com context
func BenchmarkGoodGoroutine(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SafeGoroutine(ctx)
	}
}

// Benchmark de acesso não sincronizado a map
func BenchmarkBadSharedState(b *testing.B) {
	state := BadSharedState{
		data: make(map[string]int),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.UpdateData()
	}
}

// Benchmark de acesso sincronizado a map
func BenchmarkGoodSharedState(b *testing.B) {
	state := NewSafeSharedState()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.Update("key", i)
		state.Get("key")
		state.Delete("key")
	}
}

// Benchmark de select sem timeout
func BenchmarkBadSelect(b *testing.B) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			ch1 <- 1
			time.Sleep(time.Millisecond)
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BadSelect()
	}
}

// Benchmark de select com timeout
func BenchmarkGoodSelect(b *testing.B) {
	ctx := context.Background()
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			ch1 <- 1
			time.Sleep(time.Millisecond)
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SafeSelect(ctx, ch1, ch2)
	}
}
