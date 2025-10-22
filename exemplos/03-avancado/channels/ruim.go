package channels

import (
	"fmt"
	"time"
)

// Canal sem buffer quando deveria ter
func UnbufferedBlockingChannel() {
	ch := make(chan int) // canal sem buffer

	// Sender bloqueia desnecessariamente
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i // bloqueia até alguém ler
		}
	}()

	// Receiver processando lentamente
	for i := 0; i < 1000; i++ {
		value := <-ch
		time.Sleep(time.Millisecond) // processamento lento
		_ = value
	}
}

// Fechando canal múltiplas vezes
func MultipleChannelClose() {
	ch := make(chan int)

	go func() {
		close(ch)
	}()

	go func() {
		close(ch) // panic: close of closed channel
	}()
}

// Enviando para canal fechado
func SendToClosedChannel() {
	ch := make(chan int)
	close(ch)

	ch <- 1 // panic: send on closed channel
}

// Select sem default ou timeout
func BlockingSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Pode bloquear indefinidamente
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	}
}

// Canal compartilhado sem controle de acesso
var globalChan = make(chan int)

func SharedChannelMisuse() {
	// Múltiplos escritores sem coordenação
	go func() {
		globalChan <- 1
	}()

	go func() {
		globalChan <- 2
	}()

	// Múltiplos leitores sem coordenação
	go func() {
		<-globalChan
	}()

	go func() {
		<-globalChan
	}()
}

// Direção do canal não especificada
func UndirectedChannel(ch chan int) {
	// Não fica claro se o canal é para leitura ou escrita
	ch <- 1
	<-ch
}

// Loop infinito em canal
func InfiniteChannelLoop() {
	ch := make(chan int)

	// Producer que nunca para
	go func() {
		for {
			ch <- 1
		}
	}()

	// Consumer que nunca para
	for {
		<-ch
	}
}

// Range em canal nunca fechado
func NeverClosingRange() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		// Esqueceu de fechar o canal
	}()

	// Range bloqueia para sempre
	for v := range ch {
		_ = v
	}
}

// Buffer mal dimensionado
func BadBufferSize() {
	// Buffer muito pequeno
	ch := make(chan int, 1)

	// Muitos dados para pouco buffer
	for i := 0; i < 1000000; i++ {
		ch <- i // bloqueia frequentemente
	}
}

// Ignorando erros em select
func IgnoringErrors() {
	ch := make(chan int)
	errCh := make(chan error)

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-errCh:
		// Erro ignorado
	}
}

// Vazamento de goroutine com channel
func ChannelLeakingGoroutine() {
	done := make(chan bool)

	go func() {
		// Trabalho que pode demorar
		time.Sleep(time.Hour)
		done <- true
	}()

	// Timeout muito curto, goroutine continua rodando
	select {
	case <-done:
		fmt.Println("concluído")
	case <-time.After(time.Second):
		return // goroutine vaza
	}
}

// Padrão de fan-out mal implementado
func BadFanOut() {
	input := make(chan int)

	// Número fixo e possivelmente inadequado de workers
	for i := 0; i < 100; i++ {
		go func() {
			for v := range input {
				// Processamento sem controle de erro ou cancelamento
				_ = v
			}
		}()
	}
}
