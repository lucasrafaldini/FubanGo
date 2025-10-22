package contextx
package contextx

import (
	"context"
	"testing"
	"time"
)

// Benchmark de uso de timer vs context
func BenchmarkTimer_AfterFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t := time.AfterFunc(10*time.Millisecond, func() {})
		t.Stop() // pelo menos para no benchmark
	}
}

func BenchmarkTimer_WithContext(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		cancel()
		_ = ctx
	}
}

// Benchmark de operações bloqueantes
func BenchmarkBlocking_NoContext(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1)
		ch <- 1
		<-ch
	}
}

func BenchmarkBlocking_WithContext(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1)
		ch <- 1
		_, _ = NonBlockingOperation(ctx, ch)
	}
}

// Benchmark de timeout com e sem defer cancel
func BenchmarkTimeout_NoDefer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		_ = ctx
		// sem cancel - vazamento
	}
}

func BenchmarkTimeout_WithDefer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		cancel()
		_ = ctx
	}
}

// Benchmark de type safety
func BenchmarkType_Interface(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var iface interface{} = ctx
		_ = iface
	}
}

func BenchmarkType_Context(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ContextWithCorrectType(ctx)
	}
}

// Benchmark de propagação de context
func BenchmarkPropagation_Copy(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var iface interface{} = ctx
		_ = CopyContext(iface)
	}
}

func BenchmarkPropagation_Proper(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = PropagateContext(ctx, "test")
	}
}

// Benchmark de select com context
func BenchmarkSelect_NoContext(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1)
		ch <- 42
		select {
		case v := <-ch:
			_ = v
		case <-time.After(time.Millisecond):
		}
	}
}

func BenchmarkSelect_WithContext(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1)
		ch <- 42
		select {
		case v := <-ch:
			_ = v
		case <-ctx.Done():
		}
	}
}

// Benchmark de criação de context derivado
func BenchmarkContext_Background(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		_ = ctx
	}
}

func BenchmarkContext_WithCancel(b *testing.B) {
	parent := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithCancel(parent)
		cancel()
		_ = ctx
	}
}

func BenchmarkContext_WithTimeout(b *testing.B) {
	parent := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(parent, time.Second)
		cancel()
		_ = ctx
	}
}

// Benchmark de trabalho com cancelamento
func BenchmarkWork_NoCancellation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Sleep(time.Microsecond)
	}
}

func BenchmarkWork_WithCancellation(b *testing.B) {
	ctx := context.Background()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = DoWork(ctx)
	}
}
