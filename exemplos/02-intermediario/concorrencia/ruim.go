package concorrencia

import (
	"fmt"
	"sync"
	"time"
)

// Variável global compartilhada sem proteção
var counter int

// Race condition em variável global
func BadConcurrentCounter() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++ // Race condition
			wg.Done()
		}()
	}
	wg.Wait()
}

// Deadlock clássico com dois mutexes
var (
	mutex1 = &sync.Mutex{}
	mutex2 = &sync.Mutex{}
)

func BadDeadlock() {
	// Goroutine 1
	go func() {
		mutex1.Lock()
		time.Sleep(time.Millisecond) // Aumenta chance de deadlock
		mutex2.Lock()

		// Nunca alcançado devido ao deadlock
		mutex2.Unlock()
		mutex1.Unlock()
	}()

	// Goroutine 2
	go func() {
		mutex2.Lock()
		time.Sleep(time.Millisecond) // Aumenta chance de deadlock
		mutex1.Lock()

		// Nunca alcançado devido ao deadlock
		mutex1.Unlock()
		mutex2.Unlock()
	}()
}

// Leak de goroutines
func BadGoroutineLeak() {
	// Canal sem buffer que ninguém lê
	ch := make(chan int)

	// Esta goroutine ficará presa para sempre
	go func() {
		ch <- 42 // Bloqueia para sempre
	}()
}

// Uso incorreto de canais
func BadChannelUsage() {
	ch := make(chan int, 1)

	// Fechando canal múltiplas vezes
	close(ch)
	// close(ch) // Causaria panic

	// Tentando enviar para canal fechado
	// ch <- 1 // Causaria panic

	// Buffer muito pequeno causando bloqueio
	smallBuf := make(chan int, 1)
	go func() {
		for i := 0; i < 1000; i++ {
			smallBuf <- i // Pode bloquear
		}
	}()
}

// Compartilhamento de memória sem sincronização
type BadSharedState struct {
	data map[string]int
}

func (s *BadSharedState) UpdateData() {
	// Acesso não sincronizado ao map
	go func() {
		s.data["key"] = 42
	}()
	go func() {
		delete(s.data, "key")
	}()
}

// Select mal implementado
func BadSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Select sem default ou timeout
	select {
	case <-ch1:
		fmt.Println("ch1 recebido")
	case <-ch2:
		fmt.Println("ch2 recebido")
	} // Pode bloquear para sempre
}

// Uso incorreto de WaitGroup
func BadWaitGroup() {
	var wg sync.WaitGroup

	// Esquecendo de chamar Add antes de goroutine
	go func() {
		wg.Add(1) // Muito tarde, pode perder contagem
		// trabalho...
		wg.Done()
	}()

	// WaitGroup copiado por valor em goroutine
	myWg := wg // cópia por valor!
	go func() {
		defer myWg.Done() // Opera em uma cópia!
		// trabalho...
	}()

	wg.Wait()
}

// Mutex copiado por valor
type BadMutexStruct struct {
	sync.Mutex
	count int
}

func BadMutexCopy() {
	m := BadMutexStruct{}

	// Copiando mutex por valor
	m2 := m // Cria uma cópia do mutex!

	m.Lock()
	m2.Lock() // Deadlock potencial, pois é um mutex diferente
}
