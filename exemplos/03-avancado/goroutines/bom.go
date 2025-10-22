package goroutines

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

// WorkerPool implementa um pool de workers controlado
type WorkerPool struct {
	workers  int
	tasks    chan func()
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	stopping atomic.Bool
}

func NewWorkerPool(workers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	pool := &WorkerPool{
		workers: workers,
		tasks:   make(chan func(), workers*2), // buffer para evitar bloqueio
		ctx:     ctx,
		cancel:  cancel,
	}
	pool.Start()
	return pool
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case task, ok := <-p.tasks:
					if !ok {
						return
					}
					task()
				case <-p.ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *WorkerPool) Submit(task func()) error {
	if p.stopping.Load() {
		return fmt.Errorf("pool está parando")
	}
	select {
	case p.tasks <- task:
		return nil
	case <-p.ctx.Done():
		return p.ctx.Err()
	}
}

func (p *WorkerPool) Stop() {
	p.stopping.Store(true)
	p.cancel()
	close(p.tasks)
	p.wg.Wait()
}

// SafeCounter implementa contador thread-safe
type SafeCounter struct {
	value atomic.Int64
}

func (c *SafeCounter) Increment() {
	c.value.Add(1)
}

func (c *SafeCounter) Value() int64 {
	return c.value.Load()
}

// SafeResource implementa recurso compartilhado seguro
type SafeResource struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewSafeResource() *SafeResource {
	return &SafeResource{
		data: make(map[string]string),
	}
}

func (s *SafeResource) Update(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// ProcessItems processa items com limite de concorrência
func ProcessItems(ctx context.Context, items []int) error {
	g, ctx := errgroup.WithContext(ctx)

	// Limita número de goroutines ativas
	sem := make(chan struct{}, runtime.NumCPU())

	for _, item := range items {
		item := item // copia para closure
		g.Go(func() error {
			select {
			case sem <- struct{}{}: // adquire semáforo
				defer func() { <-sem }() // libera semáforo
			case <-ctx.Done():
				return ctx.Err()
			}

			// Processa item com timeout
			timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			return processItem(timeoutCtx, item)
		})
	}

	return g.Wait()
}

func processItem(ctx context.Context, item int) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Simulação de processamento
		time.Sleep(time.Millisecond * 100)
		return nil
	}
}

// SafeGoroutine executa função com recuperação de panic
func SafeGoroutine(ctx context.Context, f func() error) error {
	var err error
	done := make(chan struct{})

	go func() {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic recuperado: %v", r)
			}
			close(done)
		}()
		err = f()
	}()

	select {
	case <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// BatchProcessor processa items em lotes
type BatchProcessor struct {
	batchSize int
	pool      *WorkerPool
}

func NewBatchProcessor(batchSize, workers int) *BatchProcessor {
	return &BatchProcessor{
		batchSize: batchSize,
		pool:      NewWorkerPool(workers),
	}
}

func (b *BatchProcessor) Process(items []int) error {
	for i := 0; i < len(items); i += b.batchSize {
		end := i + b.batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[i:end]

		if err := b.pool.Submit(func() {
			for _, item := range batch {
				_ = item // processa item
			}
		}); err != nil {
			return err
		}
	}
	return nil
}

func (b *BatchProcessor) Stop() {
	b.pool.Stop()
}

// AvoidDeadlock demonstra como evitar deadlock usando canais com buffer
func AvoidDeadlock() {
	// Usa buffer para quebrar ciclo de dependência
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	go func() {
		ch1 <- 1
		<-ch2
	}()

	go func() {
		ch2 <- 1
		<-ch1
	}()
}

// OrderedExecution garante ordem de execução com sincronização
func OrderedExecution(n int) {
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Printf("ordem: %d\n", index)
		}(i)
	}

	// Aguarda todas goroutines completarem
	wg.Wait()
}

// CancellableTimeout demonstra timeout correto com cancelamento
func CancellableTimeout(ctx context.Context) error {
	// Cria contexto com timeout
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	done := make(chan error, 1)

	go func() {
		// Trabalho que respeita cancelamento
		select {
		case <-time.After(time.Hour):
			done <- nil
		case <-ctx.Done():
			done <- ctx.Err()
		}
	}()

	// Aguarda conclusão ou timeout
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
