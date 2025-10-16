# Análise de Funções Mal Implementadas

Este documento analisa os problemas encontrados no arquivo `ruim.go` e explica por que certas práticas com funções são consideradas ruins.

## 1. Muitos Parâmetros e Retornos
```go
func ProcessUserData(firstName, lastName, email, phone, address, city, state, country, zipCode string, age int, isActive, isAdmin, isPremium bool) (string, string, error, bool, int)
```
**Problemas:**
- Número excessivo de parâmetros;
- Tipos de retorno confusos e sem nomes;
- Difícil de manter e estender;
- Violação do princípio de responsabilidade única;
- Difícil de testar.

**Como melhorar:**
- Usar structs para agrupar dados relacionados (ex: `User`, `Address`);
- Nomear os valores de retorno (ex: `fullName string, contactInfo string, err error`);
- Dividir em funções menores e mais específicas;
- Considerar o uso de builders ou construtores (ex: `UserBuilder`);
- Documentar claramente o propósito e uso.

## 2. Uso de Variáveis Globais
```go
var result string
```
**Problemas:**
- Estado global mutável (side effects);
- Difícil de rastrear modificações e bugs;
- Pode causar race conditions em código concorrente;
- Torna o código difícil de testar;
- Aumenta o acoplamento.

**Como melhorar:**
- Passar estado como parâmetro (ex: `func Process(data string) string`);
- Retornar novos valores ao invés de modificar globais (ex: `return newResult`);
- Usar closures quando necessário (para encapsular estado);
- Encapsular estado em tipos (como structs);

## 3. Código Repetitivo
```go
if isActive {
    result = result + " (Ativo)"
}
if isAdmin {
    result = result + " (Admin)"
}
if isPremium {
    result = result + " (Premium)"
}
```
**Problemas:**
- Denecessariamente verborrágico;
- Duplicação de lógica;
- Difícil de manter;
- Propenso a erros;
- Violação do princípio DRY.

**Como melhorar:**
- Usar slice/map para status (ex: `statuses := []string{}` e depois `strings.Join(statuses, ", ")`);
- Criar função auxiliar para formatação;
- Usar string builder para eficiência (em vez de concatenação direta);
- Considerar enum para status (se aplicável).

## 4. Função que Faz Muitas Coisas
```go
func DoEverything(data []string) {
    // Validação
    // Processamento
    // Formatação
    // Logging
    // Retorno de resultado
}
```
**Problemas:**
- Violação do princípio de responsabilidade única;
- Difícil de testar;
- Difícil de reutilizar;
- Mistura diferentes níveis de abstração (validação, processamento, logging);
- Difícil de manter.

**Como melhorar:**
- Dividir em funções específicas;
- Separar responsabilidades;
- Criar interfaces claras (ex: `Validator`, `Processor`, `Formatter`);
- Manter um nível consistente de abstração.

## 5. Recursão Mal Implementada
```go
func BadRecursion(n int) int
    if n <= 0 {
        return 0
    }
    return n + BadRecursion(n-1)
```
**Problemas:**
- Sem caso base explícito;
- Risco de stack overflow para valores grandes;
- Não otimizada (não usa tail recursion);
- Difícil de entender a intenção da função.

**Como melhorar:**
- Adicionar caso base explícito (comentado);
- Considerar versão iterativa (se aplicável);
- Documentar a lógica da recursão;
- Adicionar validações de entrada.

## 6. Tratamento de Erros Ignorado
```go
func IgnoreErrors(input string) string {
    num, _ := fmt.Scanf(input, "%d")
    return fmt.Sprintf("Número: %v", num)
}
```
**Problemas:**
- Ignora erros importantes;
- Pode causar comportamento inesperado (ex: `num` pode ser zero);
- Dificulta debug e manutenção;
- Código não confiável.

**Como melhorar:**
- Sempre tratar erros apropriadamente (ex: `if err != nil { return "", err }`);
- Retornar erros quando relevante (ex: `return "", fmt.Errorf("failed to parse input: %w", err)`);
- Documentar casos de erro e como lidar com eles;
- Considerar uso de tipos error personalizados (se necessário).

## 7. Função Anônima Complexa
```go
processAddress := func(addr, city, state, country, zip string) string
```
**Problemas:**
- Complexidade desnecessária;
- Difícil de reutilizar;
- Difícil de testar;
- Escopo limitado.

**Como melhorar:**
- Transformar em função nomeada;
- Dividir em funções menores;
- Usar builder pattern (ex: `AddressBuilder`);
- Melhorar legibilidade.

## Conclusão

Funções mal implementadas podem:
- Tornar o código difícil de manter;
- Causar bugs difíceis de encontrar;
- Dificultar testes;
- Reduzir a reusabilidade;
- Aumentar a complexidade desnecessariamente.

Boas práticas incluem:
- Funções pequenas e focadas;
- Nomes claros e descritivos;
- Tratamento apropriado de erros;
- Evitar estado global;
- Documentação clara;
- Testes adequados.