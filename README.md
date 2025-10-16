# FubanGo 🚫✨

FubanGo é um repositório educacional que ensina Golang através de anti-padrões e exemplos negativos.

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
        /canais
        /contexto
    /04-casos-reais
        /design-de-api
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
