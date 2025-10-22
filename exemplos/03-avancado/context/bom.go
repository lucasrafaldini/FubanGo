package contextx

import (
	"context"
	"fmt"
	"time"
)

// GoodContextUsage corrige BadContextUsage usando context.WithTimeout
// ao invés de time.AfterFunc
func GoodContextUsage() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // importante: sempre cancelar para liberar recursos

	// Usa context para controlar operação
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("trabalho concluído")
	case <-ctx.Done():
		fmt.Println("cancelado:", ctx.Err())
	}
}

// NonBlockingOperation corrige BlockingOperation aceitando context
// e usando select com ctx.Done()
func NonBlockingOperation(ctx context.Context, ch <-chan int) (int, error) {
	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

// TimeoutRespected corrige TimeoutIgnored usando context.WithTimeout
// e defer cancel() para evitar vazamento
func TimeoutRespected() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // IMPORTANTE: libera recursos do timer

	return DoWork(ctx)
}

// Exemplo correto: função que respeita contexto e deadlines
func DoWork(ctx context.Context) error {
	// Simula trabalho que pode ser cancelado
	select {
	case <-time.After(2 * time.Second):
		// trabalho concluído
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ContextWithCorrectType corrige ContextAsValueOnly usando context.Context
// ao invés de interface{}
func ContextWithCorrectType(ctx context.Context) error {
	// Pode usar métodos de context
	select {
	case <-time.After(time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PropagateContext corrige CopyContext propagando context corretamente
// com tipo explícito
func PropagateContext(ctx context.Context, data string) (string, error) {
	// Propaga context para operações downstream
	result, err := FetchWithContext(ctx, data)
	if err != nil {
		return "", fmt.Errorf("fetch failed: %w", err)
	}
	return string(result), nil
}

// Exemplo de criação de contexto com cancel e uso adequado
func Parent() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // importante para liberar recursos

	if err := DoWork(ctx); err != nil {
		fmt.Println("DoWork falhou:", err)
	}
}

// Função que aceita context como primeiro parâmetro e passa adiante
func FetchWithContext(ctx context.Context, url string) ([]byte, error) {
	// Aqui apenas um exemplo: respeitar ctx em operações de I/O
	select {
	case <-time.After(100 * time.Millisecond):
		return []byte("ok"), nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// WaitForValue demonstra select com timeout e context
func WaitForValue(ctx context.Context, ch <-chan int) (int, error) {
	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
