package channels

import (
	"context"
	"sync"
	"testing"
	"time"
)

// Benchmark de canal com e sem buffer
func BenchmarkChannel_Unbuffered(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch)
		}()
		for range ch {
		}
	}
}

func BenchmarkChannel_Buffered(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 100)
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch)
		}()
		for range ch {
		}
	}
}

// Benchmark de Pipeline
func BenchmarkPipeline_Simple(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 10)
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch)
		}()
		for range ch {
		}
	}
}

func BenchmarkPipeline_WithContext(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pipeline := NewPipeline(100)
		ctx := context.Background()
		pipeline.Process(ctx)

		// Consumer em background
		go func() {
			for range pipeline.output {
			}
		}()

		// Envia dados
		for j := 0; j < 100; j++ {
			_ = pipeline.Send(ctx, j)
		}

		// Fecha e aguarda pipeline finalizar
		close(pipeline.input)
	}
}

// Benchmark de envio seguro vs unsafe
func BenchmarkSend_Unsafe(b *testing.B) {
	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch {
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func BenchmarkSend_Safe(b *testing.B) {
	ch := NewSafeChannel(1000)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for range ch.ch {
		}
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ch.Send(i)
	}
	ch.Close()
	wg.Wait()
}

// Benchmark de select com e sem timeout
func BenchmarkSelect_Blocking(b *testing.B) {
	ch := make(chan int, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
		select {
		case v := <-ch:
			_ = v
		}
	}
}

func BenchmarkSelect_WithTimeout(b *testing.B) {
	ch := make(chan int, 1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- i
		select {
		case v := <-ch:
			_ = v
		case <-time.After(time.Millisecond):
		}
	}
}

// Benchmark de fan-out pattern
func BenchmarkFanOut_Uncontrolled(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make(chan int, 100)
		var wg sync.WaitGroup

		for w := 0; w < 10; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range input {
				}
			}()
		}

		for j := 0; j < 100; j++ {
			input <- j
		}
		close(input)
		wg.Wait()
	}
}

func BenchmarkFanOut_Controlled(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := make(chan int, 100)
		fanout := NewFanOut(input, 10, func(v int) error {
			_ = v
			return nil
		})

		go func() {
			for j := 0; j < 100; j++ {
				input <- j
			}
			close(input)
		}()

		ctx := context.Background()
		_ = fanout.Run(ctx)
	}
}

// Benchmark de range com e sem close
func BenchmarkRange_WithoutClose(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 100)
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			// não fecha - em benchmark podemos usar timeout
		}()

		count := 0
		for range ch {
			count++
			if count >= 100 {
				break
			}
		}
	}
}

func BenchmarkRange_WithClose(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 100)
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch) // fecha corretamente
		}()

		for range ch {
		}
	}
}

// Benchmark de direção de canais
func BenchmarkDirection_Bidirectional(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 10)
		go func(c chan int) {
			for j := 0; j < 10; j++ {
				c <- j
			}
			close(c)
		}(ch)

		for range ch {
		}
	}
}

func BenchmarkDirection_SendOnly(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 10)
		go func(c chan<- int) {
			for j := 0; j < 10; j++ {
				c <- j
			}
			close(c)
		}(ch)

		for range ch {
		}
	}
}

// Benchmark de buffer sizing
func BenchmarkBuffer_TooSmall(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 1) // buffer muito pequeno
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch)
		}()
		for range ch {
		}
	}
}

func BenchmarkBuffer_WellSized(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 100) // buffer adequado
		go func() {
			for j := 0; j < 100; j++ {
				ch <- j
			}
			close(ch)
		}()
		for range ch {
		}
	}
}

// Benchmark de loop controlado
func BenchmarkLoop_Infinite(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch := make(chan int, 10)
		done := make(chan struct{})

		go func() {
			count := 0
			for {
				select {
				case ch <- count:
					count++
					if count >= 100 {
						close(done)
						return
					}
				}
			}
		}()

		<-done
	}
}

func BenchmarkLoop_Controlled(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan int, 10)

		go func() {
			for j := 0; j < 100; j++ {
				select {
				case ch <- j:
				case <-ctx.Done():
					return
				}
			}
			cancel()
		}()

		<-ctx.Done()
	}
}
