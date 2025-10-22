# Análise de Uso de Context Ruim

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas com `context` são consideradas ruins.

## 1. Uso de Timer Sem Cancelamento
```go
func BadContextUsage() {
    t := time.AfterFunc(10*time.Second, func() {})
    _ = t // timer criado e não cancelado
}
```
**Problemas:**
- Timer criado mas nunca cancelado (vazamento de recursos);
- Função anônima pode executar mesmo após função retornar;
- Não usa context para controle de ciclo de vida;
- Acúmulo de timers em memória ao longo do tempo;
- Desperdício de recursos do runtime.

**Como melhorar:**
- Usar `context.WithTimeout` para controle de tempo (ex: `ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second); defer cancel()`);
- Sempre chamar `t.Stop()` para cancelar timer quando não for mais necessário;
- Armazenar referência ao timer e cancelá-lo explicitamente (ex: `defer t.Stop()`);
- Usar context para propagar cancelamento em cascata;
- Documentar ciclo de vida de recursos temporais.

## 2. Operação Bloqueante Sem Context
```go
func BlockingOperation() {
    ch := make(chan int)
    <-ch // bloqueia para sempre
}
```
**Problemas:**
- Bloqueia indefinidamente sem possibilidade de cancelamento;
- Não aceita context como parâmetro;
- Impossibilita timeout ou interrupção da operação;
- Goroutine fica presa para sempre (vazamento);
- Dificulta shutdown graceful da aplicação.

**Como melhorar:**
- Receber context como primeiro parâmetro (ex: `func BlockingOperation(ctx context.Context)`);
- Usar `select` com `ctx.Done()` para respeitar cancelamento (ex: `select { case v := <-ch: return v; case <-ctx.Done(): return ctx.Err() }`);
- Implementar timeout apropriado com `context.WithTimeout`;
- Documentar comportamento em caso de cancelamento;
- Testar cenários de timeout e cancelamento com testes unitários.

## 3. Ignorar Função Cancel de Context
```go
func TimeoutIgnored() {
    _ = time.AfterFunc(time.Second, func() {})
    // timer não armazenado nem cancelado
}
```
**Problemas:**
- Timer criado mas referência descartada imediatamente;
- Impossível cancelar o timer posteriormente;
- Vazamento de recursos (timer fica ativo até expirar);
- Não usa context.WithTimeout que fornece controle adequado;
- Degradação de performance em alta carga.

**Como melhorar:**
- Usar `context.WithTimeout` que retorna função cancel (ex: `ctx, cancel := context.WithTimeout(ctx, time.Second); defer cancel()`);
- Sempre chamar `defer cancel()` após criar contexto com timeout;
- Armazenar referência ao timer se usar `time.AfterFunc` (ex: `timer := time.AfterFunc(...); defer timer.Stop()`);
- Garantir limpeza de recursos em todos os caminhos de execução;
- Usar linters para detectar vazamentos (ex: `go vet`, `staticcheck`).

## 4. Context Como Interface Genérica
```go
func ContextAsValueOnly(ctx interface{}) {
    // Transformar context em interface{} perde a semântica
    _ = ctx
}
```
**Problemas:**
- Perde type safety ao usar `interface{}` em vez de `context.Context`;
- Impossibilita uso de métodos de context (ex: `Done()`, `Err()`, `Deadline()`);
- API não clara e propensa a erros de tipo;
- Viola convenções de Go para uso de context;
- Dificulta manutenção e compreensão do código.

**Como melhorar:**
- Sempre declarar parâmetro como `context.Context` (ex: `func ContextAsValueOnly(ctx context.Context)`);
- Nunca converter context para `interface{}` desnecessariamente;
- Usar métodos de context apropriadamente (ex: `select { case <-ctx.Done(): return ctx.Err() }`);
- Seguir convenção de context como primeiro parâmetro;
- Documentar uso de context na função.

## 5. Passar Context por Cópia Incorreta
```go
func CopyContext(ctx interface{}) interface{} {
    // Contexts devem ser passados como context.Context
    return ctx
}
```
**Problemas:**
- Usa `interface{}` em vez de `context.Context`;
- Perde semântica e capacidades de cancelamento;
- Impossibilita propagação correta de deadlines e valores;
- Código não idiomático em Go;
- Dificulta detecção de erros em tempo de compilação.

**Como melhorar:**
- Declarar tipo explícito `context.Context` (ex: `func CopyContext(ctx context.Context) context.Context`);
- Propagar context sem conversões desnecessárias;
- Criar contextos derivados quando necessário (ex: `context.WithValue(ctx, key, value)`);
- Manter assinatura de função clara e type-safe;
- Seguir guidelines oficiais de uso de context.

## Conclusão

Uso inadequado de context pode:
- Causar vazamentos de recursos (timers, goroutines);
- Impossibilitar cancelamento adequado de operações;
- Perder type safety e semântica do context;
- Dificultar manutenção e debugging;
- Violar convenções e boas práticas de Go.

Boas práticas incluem:
- Usar `context.Context` como tipo explícito;
- Sempre cancelar timers e contextos criados;
- Aceitar context como primeiro parâmetro em funções;
- Respeitar cancelamento em operações bloqueantes;
- Evitar conversões desnecessárias para `interface{}`;
- Documentar uso e ciclo de vida de recursos.