# Análise do Código Ruim de Variáveis

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas são consideradas ruins.

## 1. Nomes de Variáveis Não Descritivos
```go
var x string
var y int
var z bool
```
**Problemas:**
- Nomes como `x`, `y`, `z` não comunicam o propósito da variável;
- Dificulta a manutenção e leitura do código;
- Torna o código menos auto-documentado.

**Como melhorar:**
- Use nomes descritivos que indicam o propósito da variável (Ex: `userName`, `age`, `isActive`);
- Siga convenções de nomenclatura do Go (camelCase para variáveis);
- Considere o contexto onde a variável será usada para escolher um nome apropriado.

## 2. Não Aproveitando Inferência de Tipo
```go
var str string = "texto"
var num int = 42
var flag bool = false
```
**Problemas:**
- Sintaxe verbosa desnecessária;
- Redundância na declaração do tipo;
- Código menos conciso.

**Como melhorar:**
- Use `:=` para declaração curta com inferência de tipo;
- Use `var` apenas quando necessário (zero values ou escopo package-level).

## 3. Declarações Redundantes
```go
var slice []int = []int{}
var m map[string]int = map[string]int{}
```
**Problemas:**
- Sintaxe desnecessariamente verbosa;
- Repetição do tipo na inicialização da variável;
- Código mais difícil de ler.

**Como melhorar:**
- Use forma curta: `slice := []int{}`;
- Para maps: `m := make(map[string]int)`.

## 4. Variáveis Não Utilizadas
```go
temporary := "isso não será usado"
```
**Problemas:**
- Ocupa memória desnecessariamente;
- Causa ruído no código;
- Pode indicar lógica incompleta ou esquecida.

**Como melhorar:**
- Remova variáveis não utilizadas;
- Se temporária, use-a ou remova;
- Use _ para valores descartados.

## 5. Escopo Global Desnecessário
```go
GlobalVar := "não deveria ser global"
```
**Problemas:**
- Aumenta acoplamento;
- Dificulta testes;
- Pode causar efeitos colaterais inesperados.

**Como melhorar:**
- Mantenha variáveis no menor escopo possível;
- Use parâmetros de função ao invés de globais;
- Se global for necessário, documente bem o propósito.

## 6. Conversões Desnecessárias
```go
x = string("abc")
y = int(42)
z = bool(true)
```
**Problemas:**
- Conversões redundantes quando o tipo já é conhecido;
- Código mais verborrágico;
- Pode mascarar conversões realmente necessárias (ex: entre tipos diferentes).

**Como melhorar:**
- Remova conversões desnecessárias;
- Use conversões apenas quando tipos são diferentes.

## 7. Shadowing de Variáveis
```go
if true {
    str := "outro texto"
    println(str)
}
```
**Problemas:**
- Pode causar bugs difíceis de encontrar;
- Torna o código confuso;
- Dificulta o rastreamento do valor real da variável.

**Como melhorar:**
- Evite usar o mesmo nome de variável em escopos diferentes. Não é proibido, mas deve ser evitado;
- Use nomes únicos e descritivos;
- Mantenha escopos pequenos e claros.

## 8. Valores Mágicos
```go
if num > 42 {
    println("maior que a resposta...")
}
```
**Problemas:**
- Números mágicos dificultam a manutenção do código;
- Difícil entender o significado do valor (só duas pessoas sabem: o dev e Deus. Depois de um tempo, nem o dev lembra, só Deus sabe);
- Duplicação se o valor precisar mudar.

**Como melhorar:**
- Use constantes nomeadas para valores importantes;
- Documente o significado dos valores importantes;
- Agrupe constantes relacionadas.

## 9. Falta de Agrupamento Lógico
```go
userFirstName := "João"
userAge := 25
userLastName := "Silva"
userEmail := "joao@example.com"
```
**Problemas:**
- Variáveis relacionadas espalhadas pelo código;
- Difícil manter coesão;
- Propensão a erros na manutenção.

**Como melhorar:**
- Use structs para agrupar dados relacionados (Ex: `type User struct { FirstName string; LastName string; Age int; Email string }`);
- Mantenha ordem lógica nas declarações de variáveis;
- Considere criar tipos personalizados para conjuntos de dados.

