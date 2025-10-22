# Análise de Goroutines Mal Implementadas

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas com goroutines são consideradas ruins.

## 1. Goroutines Sem Controle de Término
```go
func LaunchUncontrolledGoroutines() {
    for i := 0; i < 1000; i++ {
        go func() {
            for {
                time.Sleep(time.Second)
                fmt.Println("ainda rodando...")
            }
        }()
    }
}
```
**Problemas:**
- Goroutines executam indefinidamente sem possibilidade de parada;
- Vazamento de recursos (1000 goroutines ficam ativas para sempre);
- Difícil de monitorar e controlar ciclo de vida;
- Impossível de cancelar ou fazer shutdown graceful;
- Consumo descontrolado de memória e CPU.

**Como melhorar:**
- Usar context para cancelamento (ex: `ctx, cancel := context.WithCancel(ctx); defer cancel()`);
- Implementar sinais de parada com canais (ex: `done := make(chan struct{})`);
- Monitorar número de goroutines ativas (ex: `runtime.NumGoroutine()`);
- Limitar número máximo de goroutines concorrentes;
- Implementar timeout para operações de longa duração.

## 2. Compartilhamento de Variáveis da Closure
```go
func ClosureVariableSharing() {
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i) // Todas veem o mesmo 'i'
        }()
    }
}
```
**Problemas:**
- Race condition na variável `i` do loop;
- Valores inesperados e imprevisíveis (geralmente imprime 10 várias vezes);
- Comportamento não determinístico;
- Difícil de debugar e reproduzir problemas;
- Resultado incorreto da execução.

**Como melhorar:**
- Passar variável como parâmetro (ex: `go func(n int) { fmt.Println(n) }(i)`);
- Criar cópia local da variável (ex: `i := i; go func() { fmt.Println(i) }()`);
- Evitar captura de variáveis mutáveis em closures;
- Documentar comportamento de closure claramente;
- Usar `go test -race` para detectar race conditions.

## 3. Número Excessivo de Goroutines
```go
func TooManyGoroutines() {
    for i := 0; i < 1000000; i++ {
        go func() {
            time.Sleep(time.Second)
        }()
    }
}
```
**Problemas:**
- Esgotamento de recursos do sistema (milhão de goroutines);
- Degradação severa de performance;
- Possível OOM (Out of Memory) crash;
- Overhead de scheduling do runtime;
- Sistema fica instável e lento.

**Como melhorar:**
- Usar worker pools para limitar concorrência (ex: usar semáforo ou canal com buffer);
- Implementar padrão de worker pool fixo (ex: `runtime.NumCPU()` workers);
- Monitorar uso de recursos (ex: memória, CPU);
- Implementar backpressure para controlar taxa de criação;
- Dimensionar adequadamente baseado em benchmarks.

## 4. Comunicação Através de Variáveis Compartilhadas
```go
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
```
**Problemas:**
- Estado compartilhado global aumenta acoplamento;
- Complexidade de sincronização com mutexes;
- Difícil de escalar e manter;
- Propenso a deadlocks em código mais complexo;
- Performance limitada por contenção de locks.

**Como melhorar:**
- Usar canais para comunicação (ex: `ch := make(chan int); go func() { ch <- value }()`);
- Seguir princípio CSP (Communicating Sequential Processes);
- Evitar estado compartilhado sempre que possível;
- Usar `sync/atomic` para operações atômicas simples (ex: `atomic.AddInt64(&counter, 1)`);
- Documentar padrão de acesso e ownership de dados.

## 5. Vazamento de Goroutines em Loops
```go
func GoroutineLeakInLoop() {
    ch := make(chan int)
    for i := 0; i < 100; i++ {
        go func() {
            ch <- i // Canal nunca é lido
        }()
    }
}
```
**Problemas:**
- Goroutines ficam bloqueadas esperando envio para canal;
- Vazamento de memória (100 goroutines presas);
- Recursos não liberados nunca;
- Sistema fica instável ao longo do tempo;
- Difícil de detectar em produção.

**Como melhorar:**
- Usar buffer apropriado no canal (ex: `ch := make(chan int, 100)`);
- Implementar timeout com select (ex: `select { case ch <- i: case <-time.After(timeout): }`);
- Garantir que canal seja lido (ex: consumidor em goroutine separada);
- Usar context para cancelamento (ex: `case <-ctx.Done():`);
- Monitorar número de goroutines com `runtime.NumGoroutine()`.

## 6. Panic em Goroutine Sem Recuperação
```go
func PanicInGoroutine() {
    go func() {
        panic("erro não tratado")
    }()
}
```
**Problemas:**
- Panic em goroutine causa crash do programa inteiro;
- Estado inconsistente após panic;
- Difícil debug sem stack trace adequado;
- Recursos não liberados (defer não executado em outras goroutines);
- Comportamento imprevisível da aplicação.

**Como melhorar:**
- Usar recover em cada goroutine (ex: `defer func() { if r := recover(); r != nil { log.Error(r) } }()`);
- Logar erros com contexto adequado;
- Implementar fallback ou retry logic;
- Monitorar panics com observabilidade (ex: métricas, alertas);
- Manter estado consistente mesmo após recuperação.

## 7. CPU-Bound com Muitas Goroutines
```go
func CPUBoundInGoroutines() {
    for i := 0; i < runtime.NumCPU()*100; i++ {
        go func() {
            for j := 0; j < 1000000; j++ {
                _ = j * j
            }
        }()
    }
}
```
**Problemas:**
- Overhead excessivo de scheduling (100x mais goroutines que cores);
- Uso ineficiente de CPU (thrashing);
- Degradação de performance por context switching;
- Consumo excessivo de recursos;
- Goroutines competindo por tempo de CPU.

**Como melhorar:**
- Limitar goroutines ao número de cores (ex: `runtime.NumCPU()` workers);
- Usar worker pool com fila de trabalho;
- Balancear carga entre workers (ex: distribuir trabalho uniformemente);
- Monitorar uso de CPU e ajustar dinamicamente;
- Otimizar algoritmos antes de paralelizar.

## 8. Sincronização Incorreta com WaitGroup
```go
func BadSynchronization() {
    var wg sync.WaitGroup
    results := make([]int, 100)
    
    for i := 0; i < 100; i++ {
        go func(i int) {
            wg.Add(1) // ERRADO: pode perder contagem
            defer wg.Done()
            results[i] = i * i
        }(i)
    }
    
    wg.Wait()
}
```
**Problemas:**
- Race condition no WaitGroup (Add dentro da goroutine);
- Contagem incorreta pode fazer Wait() retornar antes;
- Resultados podem ser perdidos ou incompletos;
- Comportamento indefinido e não determinístico;
- Difícil de debugar e reproduzir.

**Como melhorar:**
- Chamar Add antes de lançar goroutine (ex: `wg.Add(1); go func() { defer wg.Done() }()`);
- Usar `errgroup` para melhor controle de erros (ex: `g, ctx := errgroup.WithContext(ctx)`);
- Garantir sincronização adequada de acesso a slices/maps;
- Validar resultados após conclusão;
- Testar com `go test -race` para detectar race conditions.

## 9. Deadlock com Canais
```go
func DeadlockWithChannels() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go func() {
        ch1 <- 1
        <-ch2
    }()
    
    go func() {
        ch2 <- 1
        <-ch1
    }()
}
```
**Problemas:**
- Deadlock circular entre goroutines;
- Ambas goroutines bloqueiam esperando a outra;
- Programa trava completamente;
- Recursos ficam presos indefinidamente;
- Detector de deadlock do Go mata o programa.

**Como melhorar:**
- Usar canais com buffer para quebrar ciclo (ex: `make(chan int, 1)`);
- Usar select com default para evitar bloqueio (ex: `select { case ch <- v: default: }`);
- Implementar timeout com context;
- Documentar ordem de operações em canais;
- Testar cenários de concorrência extensivamente.

## 10. Ordem de Execução Não Garantida
```go
func UnpredictableOrder() {
    for i := 0; i < 10; i++ {
        go func(n int) {
            fmt.Printf("ordem: %d\n", n)
        }(i)
    }
}
```
**Problemas:**
- Ordem de execução é não determinística;
- Sem sincronização para garantir ordem;
- Saída pode variar entre execuções;
- Testes podem passar/falhar aleatoriamente;
- Difícil garantir corretude em cenários ordenados.

**Como melhorar:**
- Usar WaitGroup para sincronização (ex: `wg.Wait()` antes de retornar);
- Usar canais para ordenação se necessário (ex: pipeline pattern);
- Documentar que ordem não é garantida se for aceitável;
- Implementar barreiras de sincronização quando ordem importa;
- Projetar código para ser order-independent quando possível.

## 11. Timeout Mal Implementado
```go
func BadTimeout() {
    go func() {
        time.Sleep(time.Hour)
    }()
    
    time.Sleep(time.Second * 5)
    fmt.Println("timeout")
}
```
**Problemas:**
- Goroutine continua executando após timeout;
- Não há cancelamento real da operação;
- Vazamento de goroutine (roda por 1 hora);
- Desperdício de recursos após timeout;
- Comportamento inconsistente com expectativa.

**Como melhorar:**
- Usar context com timeout (ex: `ctx, cancel := context.WithTimeout(ctx, 5*time.Second); defer cancel()`);
- Implementar select com ctx.Done() na goroutine;
- Propagar cancelamento para operações longas;
- Garantir limpeza de recursos ao cancelar;
- Documentar comportamento de timeout claramente.

## 12. Recurso Compartilhado Sem Proteção
```go
type BadSharedResource struct {
    data map[string]string
}

func (b *BadSharedResource) UpdateConcurrently() {
    for i := 0; i < 100; i++ {
        go func(n int) {
            b.data[fmt.Sprintf("key%d", n)] = "value"
        }(i)
    }
}
```
**Problemas:**
- Race condition em acesso ao map;
- Map não é thread-safe em Go;
- Corrupção de dados possível;
- Panic em runtime (concurrent map writes);
- Crash da aplicação.

**Como melhorar:**
- Usar `sync.Mutex` para proteger acesso (ex: `mu.Lock(); defer mu.Unlock()`);
- Usar `sync.Map` para maps concorrentes;
- Usar canais para serializar acessos;
- Criar cópias imutáveis quando possível;
- Testar com `go test -race` para detectar problemas.

## Conclusão

Goroutines mal implementadas podem:
- Causar vazamentos de recursos (memória, goroutines);
- Criar race conditions e corrupção de dados;
- Degradar performance severamente;
- Causar deadlocks e travamentos;
- Tornar sistema instável e imprevisível;
- Dificultar manutenção e debugging.

Boas práticas incluem:
- Controle adequado de lifecycle com context;
- Padrões de sincronização corretos (WaitGroup, canais);
- Limitação de recursos com worker pools;
- Tratamento de erros e panics;
- Comunicação via canais ao invés de memória compartilhada;
- Monitoramento de goroutines e recursos;
- Testes com `-race` flag;
- Documentação clara de comportamento concorrente.