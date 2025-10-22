package contextx

import (
	"time"
)

// Uso incorreto de context: não propagar/ignorar cancelamento
func BadContextUsage() {
	// Exemplo ruim: tratar timers e cancelamento de forma incorreta
	// Usando time.AfterFunc como substituto indevido de context
	t := time.AfterFunc(10*time.Second, func() {})
	_ = t // timer criado e não cancelado
}

// Função que bloqueia esperando por canal sem considerar contexto
func BlockingOperation() {
	ch := make(chan int)
	<-ch // bloqueia para sempre
}

// Função que cria context com timeout mas ignora o cancel (vazamento de timer)
func TimeoutIgnored() {
	// Exemplo ilustrativo (não usa context package aqui de propósito)
	_ = time.AfterFunc(time.Second, func() {})
	// timer não armazenado nem cancelado
}

// Função que usa context apenas como valor de configuração (ruim)
func ContextAsValueOnly(ctx interface{}) {
	// Transformar context em interface{} perde a semântica e impede uso de ctx.Done()
	_ = ctx
}

// Função que passa context por cópia desnecessária usando interface{}
func CopyContext(ctx interface{}) interface{} {
	// Contexts devem ser passados como context.Context e respeitados
	return ctx
}
