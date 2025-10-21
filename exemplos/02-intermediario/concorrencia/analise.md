# Análise de Concorrência Mal Implementada

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas de concorrência são consideradas ruins.

## 1. Race Conditions
```go
var counter int
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
```
**Problemas:**
- Acesso concorrente não sincronizado a `counter`;
- Resultado imprevisível devido a condições de corrida;
- Difícil de debugar;
- Comportamento indefinido (às vezes funciona, às vezes não);
- Violação de memória.

**Como melhorar:**
- Usar sync.Mutex para proteção (lock/unlock);
- Considerar atomic operations (atomic.AddInt32);
- Evitar estado compartilhado (stateless design);
- Usar canais para comunicação (channel-based design);
- Executar `go run -race` para detectar condições de corrida.

## 2. Deadlocks
```go
func BadDeadlock() {
    go func() {
        mutex1.Lock()
        mutex2.Lock()
        // ...
    }()
    go func() {
        mutex2.Lock()
        mutex1.Lock()
        // ...
    }()

}
```
**Problemas:**
- Bloqueio mútuo permanente entre goroutines;
- Recursos presos indefinidamente (mutexes não liberados);
- Difícil de recuperar sem reiniciar o programa;
- Difícil de detectar em testes.

**Como melhorar:**
- Ordenar locks consistentemente (sempre lock mutex1 antes de mutex2);
- Usar tryLock quando possível (não bloqueia se não conseguir o lock);
- Minimizar tempo de lock (manter seção crítica curta);
- Considerar sync.RWMutex para leitura/escrita;
- Preferir canais quando apropriado.

## 3. Goroutine Leaks
```go
func BadGoroutineLeak() {
    ch := make(chan int)
    go func() {
        ch <- 42 // Bloqueia para sempre
    }()
    // Nenhum receptor para o canal, goroutine fica presa
}
```
**Problemas:**
- Goroutines presas para sempre (não conseguem enviar/receber);
- Vazamento de memória devido a goroutines não finalizadas;
- Recursos não liberados (deadlocks);
- Degradação de performance ao longo do tempo;
- Difícil de diagnosticar em produção.

**Como melhorar:**
- Sempre fornecer mecanismo de saída para goroutines;
- Usar context para cancelamento controlado (context.WithCancel);
- Implementar timeouts em operações bloqueantes (select com timeout);
- Usar buffered channels quando apropriado;
- Monitorar número de goroutines em execução.

## 4. Uso Incorreto de Canais
```go
func BadChannelUsage() {
    ch := make(chan int, 1)
    close(ch)
    // close(ch) // Panic!
    // ch <- 1   // Panic!
}
```
**Problemas:**
- Fechamento múltiplo causa panic;
- Envio para canal fechado causa panic;
- Buffer subdimensionado pode bloquear (deadlock);
- Bloqueios desnecessários em canais não bufferizados;
- Perda de mensagens se não houver receptor.

**Como melhorar:**
- Definir ownership claro (quem fecha o canal);
- Dimensionar buffers adequadamente (considerar carga);
- Usar select com default para evitar bloqueios;
- Implementar timeouts em operações de canal;
- Documentar padrões de uso de canais.

## 5. Compartilhamento sem Sincronização
```go
type BadSharedState struct {
    data map[string]int
}
func (b *BadSharedState) Update(key string, value int) {
    b.data[key] = value // Acesso concorrente sem proteção
}
```
**Problemas:**
- Race conditions em maps não sincronizados (mapas não são thread-safe);
- Corrupção de dados (mapas inconsistentes);
- Comportamento indefinido em acesso concorrente;
- Crashes aleatórios do programa;
- Difícil debug e testes.

**Como melhorar:**
- Usar sync.Map (mapa thread-safe);
- Implementar mutex para proteger acesso ao mapa;
- Considerar imutabilidade (deep copy) para leitura;
- Usar channels para comunicação de estado;
- Copiar dados quando necessário para evitar compartilhamento.

## 6. Select Mal Implementado
```go
func BadSelect() {
    select {
    case <-ch1:
    case <-ch2:
    } // Pode bloquear para sempre
}
```
**Problemas:**
- Bloqueio indefinido se nenhum canal estiver pronto;
- Sem timeout pode levar a deadlocks;
- Sem case default pode causar espera indefinida;
- Recursos presos sem liberação;
- Difícil cancelamento de operações.

**Como melhorar:**
- Adicionar timeout usando time.After;
- Incluir case default quando apropriado;
- Usar context para cancelamento controlado (context.WithTimeout);
- Implementar fallbacks em caso de bloqueio (retry logic);
- Documentar comportamento esperado do select.

## 7. WaitGroup Mal Usado
```go
func BadWaitGroup() {
    var wg sync.WaitGroup
    go func() {
        wg.Add(1) // Muito tarde
        // ...
        wg.Done()
    }()
}
```
**Problemas:**
- Contagem incorreta do WaitGroup;
- Race conditions na manipulação do WaitGroup;
- Deadlocks (WaitGroup nunca chega a zero);
- Comportamento indefinido em sincronização;
- Difícil debug e testes.

**Como melhorar:**
- Add antes de goroutine ser iniciada;
- Passar WaitGroup por ponteiro para evitar cópias;
- Documentar responsabilidades de cada goroutine;
- Manter escopo claro do WaitGroup;
- Usar defer para Done imediatamente após Add.

## 8. Mutex por Valor
```go
type BadMutexStruct struct {
    sync.Mutex
    count int
}
func (b *BadMutexStruct) Increment() {
    b.Lock()
    defer b.Unlock()
    b.count++
}

func BadMutexUsage() {
    b1 := BadMutexStruct{}
    b2 := b1 // Cópia do mutex
    go func() {
        b1.Increment()
    }()
    go func() {
        b2.Increment()
    }()
}
```
**Problemas:**
- Cópia de mutex leva a múltiplas instâncias independentes;
- Comportamento indefinido (locks não sincronizados);
- Proteção ineficaz do estado compartilhado;
- Deadlocks sutis ou corrupção de dados;
- Difícil de detectar em testes.

**Como melhorar:**
- Usar mutex por ponteiro (evitar cópias);
- Embedar mutex como ponteiro (*sync.Mutex);
- Documentar restrições de uso do struct;
- Usar ferramentas de análise estática para detectar cópias;
- Considerar sync.Map para estado compartilhado.

## Conclusão

Concorrência mal implementada pode causar:
- Comportamento imprevisível;
- Deadlocks;
- Vazamentos de memória;
- Race conditions;
- Degradação de performance.

Boas práticas incluem:
- Sincronização apropriada (mutexes, canais);
- Ownership claro de recursos (quem gerencia o quê);
- Timeouts e cancelamento (controlados via context);
- Testes com -race (ex: go test -race);
- Documentação clara das expectativas de concorrência;
- Monitoramento de recursos em produção.