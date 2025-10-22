# FubanGo 🚫✨

### **fubango**  
**_substantivo masculino_** \[informal, gíria\]  
**feminino:** *fubanga*  

1. Pessoa e/ou coisa considerada **desprovida de graça**, **desajeitada** ou de **aparência descuidada**. 

    **Ser humano de estética duvidosa**, dono de um visual que parece ter sido montado às pressas durante um apagão.
   > *“O cara era todo fubango, mas tinha um bom coração.”* 

   > *“O bicho apareceu todo fubango na festa, com a camisa florida laranja e papete do Gustavo Kuerten azul.”*  
 

2. Indivíduo tido como **cafona** ou de **gosto duvidoso** no modo de se vestir ou se comportar.

    **Pessoa com gosto peculiar**, que acredita piamente que brilho combina com oncinha e perfume forte é sinal de respeito.

   > *“Aquele terno amarelo ficou muito fubango.”*

   > *“A fubanga chegou iluminando o bar inteiro com aquele salto de acrílico.”* 


**Etimologia:** origem incerta; possivelmente criação popular de tom zombeteiro, associada à sonoridade expressiva. Talvez inventada num boteco entre um gole e outro, quando alguém precisava de uma palavra pra descrever o indescritível.  

**Sinônimos:** brega, cafona, feioso, desengonçado, bagaceiro (pop.).

**Antônimos:** estiloso, elegante, bonito.


*FubanGo* é um repositório educacional que ensina Golang através de anti-padrões e exemplos negativos.
---

## Propósito

Este projeto tem como objetivo ensinar Golang de uma maneira única: mostrando primeiro como **não** fazer as coisas. Para cada conceito, apresentamos:

1. Um problema a ser resolvido;
2. Uma implementação propositalmente ruim;
3. Análise detalhada de por que a implementação é problemática;
4. A forma correta de implementação;
5. Comparação entre as duas abordagens.

## Estrutura do Projeto

```
/exemplos
    /01-basicos
        /variaveis
        /estruturas-de-controle
        /funcoes
    /02-intermediario
        /error-handling
        /concorrencia
        /interfaces
    /03-avancado
        /goroutines
        /channels
        /context
    /04-casos-reais
        /api-design
        /database
        /testes
```

Cada diretório contém:
- `ruim.go` - Implementação propositalmente ruim
- `analise.md` - Análise detalhada dos problemas
- `bom.go` - Implementação seguindo as melhores práticas
- `benchmark_test.go` - Testes de performance (quando aplicável)

## Como Usar

1. Leia o problema proposto
2. Analise o código ruim em `ruim.go`
3. Leia a análise detalhada em `analise.md`
4. Compare com a solução correta em `bom.go`
5. Execute os benchmarks para ver a diferença de performance

## Índice de Anti-Padrões Documentados

### 📁 01-Básicos

#### [Variáveis](exemplos/01-basicos/variaveis)
1. [Nomes de Variáveis Não Descritivos](exemplos/01-basicos/variaveis/analise.md#L5) - `ruim.go:7`
2. [Não Aproveitando Inferência de Tipo](exemplos/01-basicos/variaveis/analise.md#L21) - `ruim.go:11`
3. [Declarações Redundantes](exemplos/01-basicos/variaveis/analise.md#L36) - `ruim.go:17`
4. [Variáveis Não Utilizadas](exemplos/01-basicos/variaveis/analise.md#L50) - `ruim.go:22`
5. [Escopo Global Desnecessário](exemplos/01-basicos/variaveis/analise.md#L64) - `ruim.go:26`
6. [Conversões Desnecessárias](exemplos/01-basicos/variaveis/analise.md#L78) - `ruim.go:30`
7. [Shadowing de Variáveis](exemplos/01-basicos/variaveis/analise.md#L93) - `ruim.go:35`
8. [Valores Mágicos](exemplos/01-basicos/variaveis/analise.md#L110) - `ruim.go:42`
9. [Falta de Agrupamento Lógico](exemplos/01-basicos/variaveis/analise.md#L126) - `ruim.go:47`

#### [Estruturas de Controle](exemplos/01-basicos/estruturas-de-controle)
1. [If's Aninhados Excessivamente](exemplos/01-basicos/estruturas-de-controle/analise.md#L5) - `ruim.go:7`
2. [Switch Mal Estruturado](exemplos/01-basicos/estruturas-de-controle/analise.md#L29) - `ruim.go:19`
3. [For com Continue/Break Desnecessários](exemplos/01-basicos/estruturas-de-controle/analise.md#L52) - `ruim.go:33`
4. [Loop Infinito com Break](exemplos/01-basicos/estruturas-de-controle/analise.md#L78) - `ruim.go:47`
5. [Range com Índice Não Utilizado](exemplos/01-basicos/estruturas-de-controle/analise.md#L99) - `ruim.go:61`
6. [Condições Complexas](exemplos/01-basicos/estruturas-de-controle/analise.md#L116) - `ruim.go:67`

#### [Funções](exemplos/01-basicos/funcoes)
1. [Muitos Parâmetros e Retornos](exemplos/01-basicos/funcoes/analise.md#L5) - `ruim.go:9`
2. [Uso de Variáveis Globais](exemplos/01-basicos/funcoes/analise.md#L23) - `ruim.go:7`
3. [Código Repetitivo](exemplos/01-basicos/funcoes/analise.md#L40) - `ruim.go:23`
4. [Função que Faz Muitas Coisas](exemplos/01-basicos/funcoes/analise.md#L65) - `ruim.go:34`
5. [Recursão Mal Implementada](exemplos/01-basicos/funcoes/analise.md#L88) - `ruim.go:53`
6. [Tratamento de Erros Ignorado](exemplos/01-basicos/funcoes/analise.md#L108) - `ruim.go:60`
7. [Função Anônima Complexa](exemplos/01-basicos/funcoes/analise.md#L127) - `ruim.go:66`

### 📁 02-Intermediário

#### [Error Handling](exemplos/02-intermediario/error-handling)
1. [Ignorar Erros](exemplos/02-intermediario/error-handling/analise.md#L5) - `ruim.go:10`
2. [Uso Inadequado de Panic](exemplos/02-intermediario/error-handling/analise.md#L26) - `ruim.go:18`
3. [Erros Genéricos](exemplos/02-intermediario/error-handling/analise.md#L49) - `ruim.go:25`
4. [Perda de Contexto](exemplos/02-intermediario/error-handling/analise.md#L69) - `ruim.go:34`
5. [Mistura de Erros e Logs](exemplos/02-intermediario/error-handling/analise.md#L93) - `ruim.go:43`
6. [Falta de Agrupamento de Erros](exemplos/02-intermediario/error-handling/analise.md#L117) - `ruim.go:51`
7. [Recover Indiscriminado](exemplos/02-intermediario/error-handling/analise.md#L140) - `ruim.go:64`

#### [Concorrência](exemplos/02-intermediario/concorrencia)
1. [Race Conditions](exemplos/02-intermediario/concorrencia/analise.md#L5) - `ruim.go:9`
2. [Deadlocks](exemplos/02-intermediario/concorrencia/analise.md#L34) - `ruim.go:21`
3. [Goroutine Leaks](exemplos/02-intermediario/concorrencia/analise.md#L63) - `ruim.go:44`
4. [Uso Incorreto de Canais](exemplos/02-intermediario/concorrencia/analise.md#L87) - `ruim.go:53`
5. [Compartilhamento sem Sincronização](exemplos/02-intermediario/concorrencia/analise.md#L110) - `ruim.go:63`
6. [Select Mal Implementado](exemplos/02-intermediario/concorrencia/analise.md#L133) - `ruim.go:75`
7. [WaitGroup Mal Usado](exemplos/02-intermediario/concorrencia/analise.md#L156) - `ruim.go:91`
8. [Mutex por Valor](exemplos/02-intermediario/concorrencia/analise.md#L181) - `ruim.go:104`

#### [Interfaces](exemplos/02-intermediario/interfaces)
1. [Interface Grande e Não Coesa](exemplos/02-intermediario/interfaces/analise.md#L5) - `ruim.go:7`
2. [Exposição de Detalhes de Implementação](exemplos/02-intermediario/interfaces/analise.md#L28) - `ruim.go:16`
3. [Dependência de Tipos Concretos](exemplos/02-intermediario/interfaces/analise.md#L50) - `ruim.go:24`
4. [Violação do ISP](exemplos/02-intermediario/interfaces/analise.md#L71) - `ruim.go:33`
5. [Uso Excessivo de interface{}](exemplos/02-intermediario/interfaces/analise.md#L95) - `ruim.go:48`
6. [Erro Personalizado Incorreto](exemplos/02-intermediario/interfaces/analise.md#L115) - `ruim.go:54`
7. [Container Genérico Ruim](exemplos/02-intermediario/interfaces/analise.md#L138) - `ruim.go:65`
8. [Embedding Excessivo](exemplos/02-intermediario/interfaces/analise.md#L157) - `ruim.go:75`

### 📁 03-Avançado

#### [Goroutines](exemplos/03-avancado/goroutines)
1. [Goroutines Sem Controle de Término](exemplos/03-avancado/goroutines/analise.md#L5) - `ruim.go:11`
2. [Compartilhamento de Variáveis da Closure](exemplos/03-avancado/goroutines/analise.md#L32) - `ruim.go:21`
3. [Número Excessivo de Goroutines](exemplos/03-avancado/goroutines/analise.md#L56) - `ruim.go:32`
4. [Comunicação Através de Variáveis Compartilhadas](exemplos/03-avancado/goroutines/analise.md#L80) - `ruim.go:41`
5. [Vazamento de Goroutines em Loops](exemplos/03-avancado/goroutines/analise.md#L109) - `ruim.go:58`
6. [Panic em Goroutine Sem Recuperação](exemplos/03-avancado/goroutines/analise.md#L134) - `ruim.go:69`
7. [CPU-Bound com Muitas Goroutines](exemplos/03-avancado/goroutines/analise.md#L156) - `ruim.go:82`
8. [Sincronização Incorreta com WaitGroup](exemplos/03-avancado/goroutines/analise.md#L182) - `ruim.go:96`
9. [Deadlock com Canais](exemplos/03-avancado/goroutines/analise.md#L213) - `ruim.go:110`
10. [Ordem de Execução Não Garantida](exemplos/03-avancado/goroutines/analise.md#L244) - `ruim.go:120`
11. [Timeout Mal Implementado](exemplos/03-avancado/goroutines/analise.md#L268) - `ruim.go:132`
12. [Recurso Compartilhado Sem Proteção](exemplos/03-avancado/goroutines/analise.md#L293) - `ruim.go:146`

#### [Channels](exemplos/03-avancado/channels)
1. [Canal Sem Buffer Quando Necessário](exemplos/03-avancado/channels/analise.md#L5) - `ruim.go:9`
2. [Fechamento Múltiplo](exemplos/03-avancado/channels/analise.md#L30) - `ruim.go:21`
3. [Envio para Canal Fechado](exemplos/03-avancado/channels/analise.md#L56) - `ruim.go:35`
4. [Select Bloqueante](exemplos/03-avancado/channels/analise.md#L78) - `ruim.go:48`
5. [Canal Compartilhado](exemplos/03-avancado/channels/analise.md#L103) - `ruim.go:58`
6. [Direção Não Especificada](exemplos/03-avancado/channels/analise.md#L121) - `ruim.go:72`
7. [Loop Infinito](exemplos/03-avancado/channels/analise.md#L139) - `ruim.go:78`
8. [Range Sem Fechamento](exemplos/03-avancado/channels/analise.md#L160) - `ruim.go:90`

#### [Context](exemplos/03-avancado/context)
1. [Uso de Timer Sem Cancelamento](exemplos/03-avancado/context/analise.md#L5) - `ruim.go:9`
2. [Operação Bloqueante Sem Context](exemplos/03-avancado/context/analise.md#L26) - `ruim.go:19`
3. [Ignorar Função Cancel de Context](exemplos/03-avancado/context/analise.md#L47) - `ruim.go:27`
4. [Context Como Interface Genérica](exemplos/03-avancado/context/analise.md#L68) - `ruim.go:36`
5. [Passar Context por Cópia Incorreta](exemplos/03-avancado/context/analise.md#L89) - `ruim.go:42`

### 📁 04-Casos Reais

#### [API Design](exemplos/04-casos-reais/api-design)
1. [Side-Effects em Endpoints GET](exemplos/04-casos-reais/api-design/analise.md#L5) - `ruim.go:11`
2. [Vazamento de Dados Sensíveis](exemplos/04-casos-reais/api-design/analise.md#L26) - `ruim.go:20`
3. [Falta de Autenticação e Autorização](exemplos/04-casos-reais/api-design/analise.md#L45) - `ruim.go:43`
4. [Mistura de Responsabilidades](exemplos/04-casos-reais/api-design/analise.md#L63) - `ruim.go:51`
5. [Falta de Versionamento](exemplos/04-casos-reais/api-design/analise.md#L84) - `ruim.go:72`
6. [Ausência de Contratos Claros](exemplos/04-casos-reais/api-design/analise.md#L104) - `ruim.go:79`
7. [Falta de Tratamento de Erros Consistente](exemplos/04-casos-reais/api-design/analise.md#L125) - `ruim.go:85`

#### [Database](exemplos/04-casos-reais/database)
1. [Abrir e Fechar Conexão Por Requisição](exemplos/04-casos-reais/database/analise.md#L5) - `ruim.go:11`
2. [SQL Injection por Concatenação de Strings](exemplos/04-casos-reais/database/analise.md#L29) - `ruim.go:21`
3. [Ignorar Erros de Operações de Banco](exemplos/04-casos-reais/database/analise.md#L48) - `ruim.go:31`
4. [Transação Sem Rollback em Caso de Erro](exemplos/04-casos-reais/database/analise.md#L70) - `ruim.go:47`
5. [Falta de Context com Timeout](exemplos/04-casos-reais/database/analise.md#L93) - `ruim.go:64`
6. [Não Usar Prepared Statements](exemplos/04-casos-reais/database/analise.md#L112) - `ruim.go:73`
7. [Ausência de Migrations e Versionamento de Schema](exemplos/04-casos-reais/database/analise.md#L130) - `ruim.go:80`

#### [Testes](exemplos/04-casos-reais/testes)
1. [Testes Dependentes de Ordem e Estado Global](exemplos/04-casos-reais/testes/analise.md#L5) - `ruim.go:11`
2. [Testes Lentos Sem Mocks](exemplos/04-casos-reais/testes/analise.md#L25) - `ruim.go:25`
3. [Uso de Sleep em Testes](exemplos/04-casos-reais/testes/analise.md#L45) - `ruim.go:43`
4. [Falta de Assertions e Validações](exemplos/04-casos-reais/testes/analise.md#L67) - `ruim.go:59`
5. [Testes Não Isolados (Sem Setup/Teardown)](exemplos/04-casos-reais/testes/analise.md#L87) - `ruim.go:75`
6. [Ausência de Table-Driven Tests](exemplos/04-casos-reais/testes/analise.md#L107) - `ruim.go:98`
7. [Não Rodar com Race Detector](exemplos/04-casos-reais/testes/analise.md#L125) - `ruim.go:117`
8. [Cobertura de Testes Não Medida](exemplos/04-casos-reais/testes/analise.md#L147) - `ruim.go:134`
9. [Mocks Mal Implementados](exemplos/04-casos-reais/testes/analise.md#L165) - `ruim.go:151`

---

**Total: 87 anti-padrões documentados** 🚫

## Contribuindo

Contribuições são bem-vindas! Se você tem um exemplo de código ruim que pode ser educativo, sinta-se à vontade para abrir um PR.

## Aviso

⚠️ O código em arquivos `ruim.go` é propositalmente ruim e NÃO deve ser usado em produção!

## Rodando os Benchmarks (parte executável)

Este repositório é principalmente um conjunto de exemplos para leitura. A única parte "executável" são os benchmarks em cada exemplo. Para rodá-los localmente siga as instruções abaixo.

Pré-requisitos:
- Go 1.18+ instalado (recomendo 1.20+)

Passos rápidos:

1. Inicialize o módulo (já incluí um `go.mod` mínimo no repositório). Se você quiser usar outro módulo, ajuste conforme necessário:

```bash
# se quiser recriar o módulo com outro path
go mod init github.com/SEU_USUARIO/fubango
```

2. Baixe dependências (se houver):

```bash
go mod tidy
```

3. Execute todos os benchmarks do projeto (recomendado):

```bash
go test -bench . ./exemplos/... -benchmem
```

4. Executar benchmarks de um pacote específico (ex.: variables):

```bash
go test -bench . ./exemplos/01-basicos/variaveis -benchmem
```

Dicas:
- Para rodar com detector de race (útil ao comparar implementações ruins/boas):

```bash
go test -race -bench . ./exemplos/... -benchmem
```

- Para coletar resultados em formato compatível com benchstat você pode gravar saídas em arquivos e comparar:

```bash
go test -bench . ./exemplos/01-basicos/variaveis -benchmem > ruim.txt
go test -bench . ./exemplos/01-basicos/variaveis -benchmem -run BenchmarkGood -benchmem > bom.txt
benchstat ruim.txt bom.txt
```

Observação: alguns benchmarks ilustrativos podem depender de pacotes externos (por exemplo `golang.org/x/sync/errgroup`) — rode `go mod tidy` para buscar as dependências necessárias antes de executar os benchs.

## Roadmap

Para ver o plano completo de evolução do projeto, incluindo próximas fases, metas e cronograma detalhado, consulte o **[ROADMAP.md](ROADMAP.md)**.

**Status Atual:** Fase 1 - 100% Completa ✅ | [Ver detalhes →](ROADMAP.md)

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE) - veja o arquivo LICENSE para mais detalhes.

Copyright © 2025 Lucas Rafaldini

---

**Feito com ❤️ para a comunidade Go**  
*"Aprendendo com os erros, evoluindo com os acertos"*
