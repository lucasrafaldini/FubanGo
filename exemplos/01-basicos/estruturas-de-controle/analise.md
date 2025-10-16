# Análise de Estruturas de Controle Ruins

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas com estruturas de controle são consideradas ruins.

## 1. If's Aninhados Excessivamente
```go
if x > 0 {
    if x < 20 {
        if x%2 == 0 {
            if x > 5 {
                fmt.Println("x é par, maior que 5 e menor que 20")
            }
        }
    }
}
```
**Problemas:**
- Difícil de ler e manter;
- Aumenta a complexidade ciclomática (isso significa mais caminhos possíveis no código);
- Dificulta o entendimento do fluxo lógico;
- Torna o código mais propenso a erros.

**Como melhorar:**
- Combinar condições usando operadores lógicos && (AND) e || (OR);
- Early return para condições negativas (reduzindo aninhamento);
- Extrair para funções menores com nomes descritivos;
- Usar cláusulas guard para simplificar a lógica.

## 2. Switch Mal Estruturado
```go
switch x {
case 10:
    fmt.Println("é 10")
case 5 + 5:
    fmt.Println("também é 10")
case 20 - 10:
    fmt.Println("ainda é 10")
}
```
**Problemas:**
- Casos redundantes;
- Falta de case default;
- Expressões desnecessárias nos cases (I say tomeito, you say tomáto; I say poteito, you say potáto);
- Código não robusto para mudanças (se x mudar, o switch pode não funcionar como esperado).

**Como melhorar:**
- Eliminar redundância nos cases;
- SEMPRE incluir um case default;
- Usar expressões simples e claras (não complexas e desnecessariamente obscuras);
- Agrupar casos relacionados.

## 3. For com Continue/Break Desnecessários
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    } else {
        if i > 5 {
            break
        } else {
            fmt.Println(i)
        }
    }
}
```
**Problemas:**
- Lógica confusa e difícil de seguir;
- Uso excessivo de continue/break (parece que aprendeu ontem e quer usar em tudo);
- Estrutura else desnecessária;
- Código verboso e redundante.

**Como melhorar:**
- Inverter condições para simplificar (early return);
- Eliminar continue/break quando possível;
- Remover else desnecessários;
- Extrair lógica complexa para funções.

## 4. Loop Infinito com Break
```go
counter := 0
for {
    counter++
    if counter >= 10 {
        break
    }
}
```
**Problemas:**
- Loop infinito com condição de saída;
- Difícil de entender a intenção;
- Pode causar problemas de performance;
- Código não idiomático.

**Como melhorar:**
- Usar for com condição explícita (mais claro e idiomático);
- Definir claramente critério de parada no loop;
- Evitar loops infinitos quando possível (o que é quase sempre).

## 5. Range com Índice Não Utilizado
```go
for i, _ := range numbers {
    numbers[i] = numbers[i] * 2
}
```
**Problemas:**
- Sintaxe denecessariamente verbosa;
- Variável blank desnecessária;
- Não utiliza as características do range (valor retornado);
- Código menos legível.

**Como melhorar:**
- Omitir variáveis não utilizadas;
- Usar sintaxe mais concisa;
- Aproveitar valor retornado pelo range.

## 6. Condições Complexas
```go
if ((a+b)*c)/(a*b) > 1 && (a+b+c)%2 == 0 || (a*b*c)%2 == 1 {
    fmt.Println("Condição complexa atendida")
}
```
**Problemas:**
- Difícil de entender e manter;
- Alta probabilidade de erros (era pra ser uma linha, mas para corrigir vai ter que quebrar em várias);
- Difícil de testar;
- Baixa legibilidade.

**Como melhorar:**
- Quebrar em múltiplas condições;
- Criar funções auxiliares;
- Usar variáveis intermediárias;
- Documentar a lógica complexa.

## Conclusão

Estruturas de controle mal implementadas podem tornar o código:
- Difícil de entender;
- Propenso a erros;
- Difícil de manter;
- Menos performático;
- Menos testável.

Seguir as boas práticas ajuda a criar código mais:
- Legível;
- De fácil manutenção;
- Confiável;
- Eficiente;
- Facilmente testável.