package channels

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Pipeline representa um pipeline de processamento
type Pipeline struct {
	input  chan int
	output chan int
	done   chan struct{}
}

// NewPipeline cria um novo pipeline com buffer apropriado
func NewPipeline(bufferSize int) *Pipeline {
	return &Pipeline{
		input:  make(chan int, bufferSize),
		output: make(chan int, bufferSize),
		done:   make(chan struct{}),
	}
}

// Process processa dados com controle de cancelamento
func (p *Pipeline) Process(ctx context.Context) {
	go func() {
		defer close(p.output)
		for {
			select {
			case value, ok := <-p.input:
				if !ok {
					return
				}
				select {
				case p.output <- value * 2:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Send envia dados com timeout
func (p *Pipeline) Send(ctx context.Context, value int) error {
	select {
	case p.input <- value:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Second):
		return fmt.Errorf("timeout ao enviar")
	}
}

// Close fecha o pipeline de forma segura
func (p *Pipeline) Close() {
	close(p.input)
	<-p.done
}

// SafeChannel encapsula um canal com controle de acesso
type SafeChannel struct {
	ch     chan int
	closed bool
	mu     sync.RWMutex
}

// NewSafeChannel cria um novo canal seguro
func NewSafeChannel(buffer int) *SafeChannel {
	return &SafeChannel{
		ch: make(chan int, buffer),
	}
}

// Send envia dados de forma segura
func (sc *SafeChannel) Send(value int) error {
	sc.mu.RLock()
	defer sc.mu.RUnlock()

	if sc.closed {
		return fmt.Errorf("canal fechado")
	}

	sc.ch <- value
	return nil
}

// Close fecha o canal de forma segura
func (sc *SafeChannel) Close() {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if !sc.closed {
		sc.closed = true
		close(sc.ch)
	}
}

// FanOut implementa o padrão fan-out com controle
type FanOut struct {
	input     <-chan int
	workers   int
	processor func(int) error
	errChan   chan error
}

// NewFanOut cria uma nova instância de FanOut
func NewFanOut(input <-chan int, workers int, processor func(int) error) *FanOut {
	return &FanOut{
		input:     input,
		workers:   workers,
		processor: processor,
		errChan:   make(chan error, workers),
	}
}

// Run executa o processamento com N workers
func (f *FanOut) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	for i := 0; i < f.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case value, ok := <-f.input:
					if !ok {
						return
					}
					if err := f.processor(value); err != nil {
						select {
						case f.errChan <- err:
						default:
							// Buffer cheio, loga erro
						}
					}
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(f.errChan)
	}()

	// Coleta erros
	for err := range f.errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

// BufferedPipe implementa um pipe com buffer dinâmico
type BufferedPipe struct {
	input    chan int
	output   chan int
	buffer   []int
	capacity int
	mu       sync.Mutex
}

// NewBufferedPipe cria um novo pipe com buffer
func NewBufferedPipe(capacity int) *BufferedPipe {
	bp := &BufferedPipe{
		input:    make(chan int),
		output:   make(chan int),
		capacity: capacity,
	}
	go bp.process()
	return bp
}

func (bp *BufferedPipe) process() {
	for {
		if len(bp.buffer) == 0 {
			// Buffer vazio, espera por input
			value, ok := <-bp.input
			if !ok {
				close(bp.output)
				return
			}
			bp.buffer = append(bp.buffer, value)
			continue
		}

		select {
		case value, ok := <-bp.input:
			if !ok {
				// Drena buffer e fecha
				for _, v := range bp.buffer {
					bp.output <- v
				}
				close(bp.output)
				return
			}
			if len(bp.buffer) < bp.capacity {
				bp.buffer = append(bp.buffer, value)
			}
		case bp.output <- bp.buffer[0]:
			bp.buffer = bp.buffer[1:]
		}
	}
}

// ProperRangeWithClose demonstra uso correto de range com close
func ProperRangeWithClose() {
	ch := make(chan int, 10)

	// Producer fecha o canal quando termina
	go func() {
		defer close(ch) // IMPORTANTE: fecha o canal
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	// Consumer usa range que termina quando canal fecha
	for v := range ch {
		_ = v
	}
}

// WellSizedBuffer demonstra dimensionamento adequado de buffer
func WellSizedBuffer(itemCount int) {
	// Buffer dimensionado baseado na carga esperada
	bufferSize := itemCount / 10 // 10% da carga
	if bufferSize < 10 {
		bufferSize = 10 // mínimo razoável
	}
	if bufferSize > 1000 {
		bufferSize = 1000 // máximo para evitar uso excessivo de memória
	}

	ch := make(chan int, bufferSize)

	go func() {
		defer close(ch)
		for i := 0; i < itemCount; i++ {
			ch <- i
		}
	}()

	for v := range ch {
		_ = v
	}
}

// DirectedChannels demonstra uso correto de direção de canais
func DirectedChannels() {
	ch := make(chan int, 5)

	// Goroutine que só envia
	go sendOnly(ch)

	// Goroutine que só recebe
	receiveOnly(ch)
}

func sendOnly(ch chan<- int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func receiveOnly(ch <-chan int) {
	for v := range ch {
		_ = v
	}
}

// ControlledLoop demonstra loop com controle de parada via context
func ControlledLoop(ctx context.Context) {
	ch := make(chan int, 10)

	// Producer com controle
	go func() {
		defer close(ch)
		for i := 0; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()

	// Consumer com controle
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return
			}
			_ = v
		case <-ctx.Done():
			return
		}
	}
}
