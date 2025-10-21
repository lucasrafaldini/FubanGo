package concorrencia

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// SafeCounter implementa um contador thread-safe
type SafeCounter struct {
	value atomic.Int64
}

func (c *SafeCounter) Increment() {
	c.value.Add(1)
}

func (c *SafeCounter) Value() int64 {
	return c.value.Load()
}

// SafeConcurrentCounter usa atomic operations
func SafeConcurrentCounter() int64 {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	return counter.Value()
}

// SafeResource implementa locks ordenados
type SafeResource struct {
	mu    sync.Mutex
	value int
}

func (r *SafeResource) Update(other *SafeResource) bool {
	// Previne deadlock ordenando locks
	first, second := r, other
	if uintptr(unsafe.Pointer(other)) < uintptr(unsafe.Pointer(r)) {
		first, second = other, r
	}

	first.mu.Lock()
	defer first.mu.Unlock()

	second.mu.Lock()
	defer second.mu.Unlock()

	// Operação segura
	r.value += other.value
	return true
}

// SafeGoroutine implementa cancelamento via context
func SafeGoroutine(ctx context.Context) error {
	ch := make(chan int, 1) // buffer previne leak

	go func() {
		defer close(ch)
		// Simula trabalho
		time.Sleep(time.Second)
		select {
		case ch <- 42:
		case <-ctx.Done():
			return
		}
	}()

	select {
	case result := <-ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(2 * time.Second):
		return context.DeadlineExceeded
	}

}

// SafeSharedState implementa acesso thread-safe a map
type SafeSharedState struct {
	sync.RWMutex
	data map[string]int
}

func NewSafeSharedState() *SafeSharedState {
	return &SafeSharedState{
		data: make(map[string]int),
	}
}

func (s *SafeSharedState) Update(key string, value int) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *SafeSharedState) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
}

func (s *SafeSharedState) Get(key string) (int, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// SafeSelect implementa timeout e cancelamento
func SafeSelect(ctx context.Context, ch1, ch2 <-chan int) (int, error) {
	select {
	case val := <-ch1:
		return val, nil
	case val := <-ch2:
		return val, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	case <-time.After(time.Second):
		return 0, context.DeadlineExceeded
	default:
		return 0, nil // Non-blocking
	}
}

// Worker representa uma tarefa com gerenciamento seguro
type Worker struct {
	wg   *sync.WaitGroup
	done chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		wg:   &sync.WaitGroup{},
		done: make(chan struct{}),
	}
}

func (w *Worker) Start(tasks []func()) {
	for _, task := range tasks {
		w.wg.Add(1)
		go func(t func()) {
			defer w.wg.Done()
			select {
			case <-w.done:
				return
			default:
				t()
			}
		}(task)
	}
}

func (w *Worker) Stop() {
	close(w.done)
	w.wg.Wait()
}

// ThreadSafeStruct implementa mutex como ponteiro
type ThreadSafeStruct struct {
	mu    *sync.Mutex // Ponteiro para prevenir cópia
	count int
}

func NewThreadSafeStruct() *ThreadSafeStruct {
	return &ThreadSafeStruct{
		mu: &sync.Mutex{},
	}
}

func (t *ThreadSafeStruct) Increment() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.count++
}
