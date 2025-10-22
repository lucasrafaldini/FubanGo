package goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Goroutine sem controle de término
func LaunchUncontrolledGoroutines() {
	for i := 0; i < 1000; i++ {
		go func() {
			// Goroutine que roda indefinidamente
			for {
				time.Sleep(time.Second)
				fmt.Println("ainda rodando...")
			}
		}()
	}
}

// Compartilhamento de variáveis da closure
func ClosureVariableSharing() {
	for i := 0; i < 10; i++ {
		go func() {
			// Todas as goroutines veem o mesmo 'i'
			fmt.Println(i)
		}()
	}
}

// Número excessivo de goroutines
func TooManyGoroutines() {
	// Criando goroutines sem limite
	for i := 0; i < 1000000; i++ {
		go func() {
			// Simulando trabalho
			time.Sleep(time.Second)
		}()
	}
}

// Comunicação através de variáveis compartilhadas
var sharedCounter int
var mutex sync.Mutex

func BadCommunication() {
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			sharedCounter++
			mutex.Unlock()
		}()
	}
}

// Goroutines vazando em loops
func GoroutineLeakInLoop() {
	ch := make(chan int)

	for i := 0; i < 100; i++ {
		go func() {
			// Canal nunca é lido
			ch <- i
		}()
	}
}

// Panic em goroutine sem recuperação
func PanicInGoroutine() {
	go func() {
		// Panic não recuperado quebra o programa
		panic("erro não tratado")
	}()
}

// CPU-bound em muitas goroutines
func CPUBoundInGoroutines() {
	// Criando mais goroutines que núcleos de CPU
	for i := 0; i < runtime.NumCPU()*100; i++ {
		go func() {
			// Trabalho CPU-intensivo
			for j := 0; j < 1000000; j++ {
				_ = j * j
			}
		}()
	}
}

// Sincronização incorreta
func BadSynchronization() {
	var wg sync.WaitGroup
	results := make([]int, 100)

	for i := 0; i < 100; i++ {
		// WaitGroup.Add deve ser chamado antes da goroutine
		go func(i int) {
			wg.Add(1) // ERRADO: pode perder contagem
			defer wg.Done()
			results[i] = i * i
		}(i)
	}

	wg.Wait() // Pode terminar antes das goroutines começarem
}

// Bloqueio mútuo com canais
func DeadlockWithChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		// Tentando enviar para ch1 e receber de ch2
		ch1 <- 1
		<-ch2
	}()

	go func() {
		// Tentando enviar para ch2 e receber de ch1
		ch2 <- 1
		<-ch1
	}()
}

// Ordem de execução não garantida
func UnpredictableOrder() {
	for i := 0; i < 10; i++ {
		go func(n int) {
			fmt.Printf("ordem: %d\n", n)
		}(i)
	}
	// Sem sincronização, ordem é imprevisível
}

// Timeout mal implementado
func BadTimeout() {
	go func() {
		// Trabalho longo sem possibilidade de cancelamento
		time.Sleep(time.Hour)
	}()

	// Timeout não afeta a goroutine
	time.Sleep(time.Second * 5)
	fmt.Println("timeout")
}

// Recurso compartilhado sem proteção
type BadSharedResource struct {
	data map[string]string
}

func (b *BadSharedResource) UpdateConcurrently() {
	for i := 0; i < 100; i++ {
		go func(n int) {
			// Race condition no map
			b.data[fmt.Sprintf("key%d", n)] = "value"
		}(i)
	}
}
