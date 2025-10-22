package goroutines

import (
	"context"
	"sync"
	"testing"
	"time"
)

// Benchmark de criação de goroutines sem controle vs worker pool
func BenchmarkGoroutineCreation_Uncontrolled(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		go func() {
			time.Sleep(time.Microsecond)
			close(done)
		}()
		<-done
	}
}

func BenchmarkGoroutineCreation_WorkerPool(b *testing.B) {
	pool := NewWorkerPool(10)
	defer pool.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		done := make(chan struct{})
		_ = pool.Submit(func() {
			time.Sleep(time.Microsecond)
			close(done)
		})
		<-done
	}
}

// Benchmark de closure variable sharing
func BenchmarkClosureSharing_Bad(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		results := make([]int, 10)

		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				results[j] = j // race condition com j
			}()
		}
		wg.Wait()
	}
}

func BenchmarkClosureSharing_Good(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		results := make([]int, 10)

		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				results[index] = index
			}(j)
		}
		wg.Wait()
	}
}

// Benchmark de comunicação: mutex vs atomic
func BenchmarkCommunication_Mutex(b *testing.B) {
	var mu sync.Mutex
	var counter int

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	})
}

func BenchmarkCommunication_Atomic(b *testing.B) {
	counter := &SafeCounter{}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})
}

// Benchmark de recurso compartilhado
func BenchmarkSharedResource_Unsafe(b *testing.B) {
	resource := &BadSharedResource{
		data: make(map[string]string),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Apenas 1 update por iteração para evitar crash
		resource.data["key"] = "value"
	}
}

func BenchmarkSharedResource_Safe(b *testing.B) {
	resource := NewSafeResource()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		resource.Update("key", "value")
	}
}

// Benchmark de concorrência limitada
func BenchmarkConcurrency_Unlimited(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for _, item := range items {
			wg.Add(1)
			item := item
			go func() {
				defer wg.Done()
				_ = item * item
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrency_Limited(b *testing.B) {
	items := make([]int, 100)
	for i := range items {
		items[i] = i
	}
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ProcessItems(ctx, items)
	}
}

// Benchmark de processamento em lote
func BenchmarkBatch_NoPool(b *testing.B) {
	items := make([]int, 1000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for _, item := range items {
			wg.Add(1)
			item := item
			go func() {
				defer wg.Done()
				_ = item * item
			}()
		}
		wg.Wait()
	}
}

func BenchmarkBatch_WithPool(b *testing.B) {
	items := make([]int, 1000)
	processor := NewBatchProcessor(100, 10)
	defer processor.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processor.Process(items)
	}
}

// Benchmark de panic recovery
func BenchmarkPanic_NoRecover(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			defer func() {
				_ = recover()
			}()
			_ = 42
		}()
	}
}

func BenchmarkPanic_WithRecover(b *testing.B) {
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SafeGoroutine(ctx, func() error {
			return nil
		})
	}
}

// Benchmark de sincronização com WaitGroup
func BenchmarkSync_BadWaitGroup(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 10; j++ {
			go func() {
				wg.Add(1) // RUIM: dentro da goroutine
				defer wg.Done()
				time.Sleep(time.Microsecond)
			}()
		}
		time.Sleep(time.Millisecond) // espera "suficiente"
	}
}

func BenchmarkSync_GoodWaitGroup(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 10; j++ {
			wg.Add(1) // BOM: antes da goroutine
			go func() {
				defer wg.Done()
				time.Sleep(time.Microsecond)
			}()
		}
		wg.Wait()
	}
}
