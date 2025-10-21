# Análise de Interfaces Mal Implementadas

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas com interfaces são consideradas ruins.

## 1. Interface Grande e Não Coesa
```go
type BigInterface interface {
    ReadData() []byte
    WriteData([]byte)
    ProcessData()
    // ... muitos outros métodos
}
```
**Problemas:**
- Viola o Princípio da Interface Segregada (ISP);
- Difícil de implementar completamente;
- Baixa coesão (muitos comportamentos diferentes);
- Acoplamento alto (difícil de modificar uma parte sem afetar outras);
- Difícil de manter e evoluir ao longo do tempo.

**Como melhorar:**
- Dividir em interfaces menores;
- Focar em comportamentos específicos;
- Seguir o princípio da responsabilidade única (SRP);
- Criar interfaces por necessidade de uso;
- Compor interfaces quando necessário (em vez de criar grandes interfaces).

## 2. Exposição de Detalhes de Implementação
```go
type BadDatabase interface {
    ConnectToMySQL()
    ExecuteSQLQuery(string)
    CloseMySQL()
}
```
**Problemas:**
- Acoplamento com tecnologia específica (MySQL);
- Difícil trocar implementação (ex: mudar para PostgreSQL);
- Viola princípio de abstração (dependência de detalhes);
- Inflexível para testes;
- Difícil de mockear em testes unitários.

**Como melhorar:**
- Abstrair detalhes de implementação (ex: Connect, ExecuteQuery, Close);
- Usar nomes genéricos (ex: DatabaseConnector);
- Focar em comportamento em vez de tecnologia;
- Permitir diferentes implementações (ex: MySQL, PostgreSQL, SQLite);
- Facilitar testes e mocks.

## 3. Dependência de Tipos Concretos
```go
type BadProcessor interface {
    Process(*BigImplementation)
    Handle(*CustomError)
}
```
**Problemas:**
- Acoplamento forte a tipos concretos;
- Difícil de testar (com tipos específicos);
- Inflexível para mudanças futuras;
- Viola inversão de dependência (DIP);
- Difícil de estender com novas implementações.

**Como melhorar:**
- Depender de interfaces em vez de tipos concretos (ex: Process(DataReader));
- Usar tipos abstratos (em vez de específicos);
- Injetar dependências via interfaces (ex: Handle(error));
- Permitir múltiplas implementações (ex: diferentes tipos de erros);
- Facilitar mocks e testes unitários.

## 4. Violação do ISP
```go
type BadWorker interface {
    DoWork()
    SendEmail()
    GenerateReport()
    UpdateDatabase()
    NotifyAdmin()
}
```
**Problemas:**
- Força implementações desnecessárias (ex: SimpleWorker não precisa de todos os métodos);
- Viola princípio de segregação de interfaces;
- Aumenta acoplamento entre componentes;
- Dificulta manutenção e evolução;
- Código não utilizado em muitas implementações (métodos desnecessários).

**Como melhorar:**
- Criar interfaces específicas para cada responsabilidade (ex: Worker, EmailSender, ReportGenerator);
- Separar responsabilidades claramente;
- Compor quando necessário (em vez de criar grandes interfaces);
- Implementar apenas o necessário em cada tipo;
- Seguir SOLID para design de interfaces (Single Responsibility e Interface Segregation Principles).

## 5. Uso Excessivo de interface{}
```go
type BadAcceptor interface {
    Accept(interface{})
}
```
**Problemas:**
- Perde type safety (não há verificação de tipos em tempo de compilação);
- Necessita type assertions frequentes;
- Propenso a erros em runtime (devido a tipos incorretos);
- Difícil de manter e entender o que é esperado;
- Performance inferior (devido a boxing/unboxing).

**Como melhorar:**
- Usar generics (ex: Accept[T any](T));
- Criar interfaces específicas para tipos esperados;
- Definir tipos concretos quando possível;
- Evitar type assertions desnecessárias;
- Manter type safety e clareza no código.

## 6. Erro Personalizado Incorreto
```go
type CustomError struct {
    message string
}
func (e CustomError) String() string {
    return e.message
}
```
**Problemas:**
- Não implementa interface error (não pode ser usado como erro padrão);
- Incompatível com standard library do Go;
- Difícil de usar com funções padrão de tratamento de erros;
- Inconsistente com Go idioms (ex: errors.New, fmt.Errorf);
- Perde funcionalidade de erros (não suporta wrapping).

**Como melhorar:**
- Implementar interface error (func (e CustomError) Error() string);
- Seguir padrões do Go para erros;
- Adicionar contexto útil ao erro;
- Permitir wrapping de erros (usando fmt.Errorf ou errors.Wrap);
- Facilitar handling e propagação de erros.

## 7. Container Genérico Ruim
```go
type BadContainer struct {
    data interface{}
}
```
**Problemas:**
- Perde type safety (não há verificação de tipos);
- Necessita type assertions frequentes;
- Propenso a erros em runtime;
- Difícil de usar e entender o que é armazenado;
- Performance inferior (devido a boxing/unboxing).

**Como melhorar:**
- Usar generics (ex: type Container[T any] struct { data T });
- Criar tipos específicos quando possível;
- Manter type safety;
- Documentar restrições de tipos;

## 8. Embedding Excessivo
```go
type BadReadWriter interface {
    Reader
    Writer
    // ... muitos outros métodos
}
```
**Problemas:**
- Interface muito grande;
- Baixa coesão (combina muitos comportamentos);
- Difícil de implementar completamente;
- Viola ISP (força implementação de métodos desnecessários);
- Acoplamento desnecessário entre componentes.

**Como melhorar:**
- Criar interfaces menores e específicas (ex: Reader, Writer);
- Compor quando necessário (em vez de criar grandes interfaces);
- Manter coesão e clareza;
- Seguir necessidade real de uso;
- Documentar propósito de cada interface.

## Conclusão

Interfaces mal projetadas podem:
- Aumentar complexidade;
- Dificultar manutenção;
- Reduzir flexibilidade;
- Complicar testes;
- Prejudicar performance;

Boas práticas incluem:
- Interfaces pequenas e coesas;
- Foco em comportamento;
- Composição quando apropriado;
- Design por contrato;
- Seguir SOLID;
- Facilitar testes;
- Manter type safety.