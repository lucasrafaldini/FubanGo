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

## 🗺️ Roadmap

### Fase 1: Completar Exemplos Go (Em Andamento)
- [x] **Exemplos Básicos**: Variáveis, estruturas de controle, funções
- [x] **Exemplos Intermediários**: Error handling, concorrência, interfaces  
- [x] **Exemplos Avançados**: Goroutines, canais, context
- [x] **Casos Reais**: API design, database, testes
- [ ] **Documentação**: Completar análises em português para todos os exemplos

### Fase 2: Divulgação e Colaboração
- [ ] **Comunidade**: Publicar em fóruns Go (Reddit, Discord, Slack)
- [ ] **Redes Sociais**: Divulgar no Twitter, LinkedIn, dev.to, etc.
- [ ] **Conferências**: Submeter palestras em eventos Go
- [ ] **Colaboradores**: Recrutar contribuidores da comunidade
- [ ] **Issues**: Criar issues para novos exemplos e melhorias
- [ ] **Hacktoberfest**: Participar em eventos de código aberto
- [ ] **Internacionalização**: Traduzir documentação para:
  - [ ] **Inglês**: Expandir para desenvolvedores internacionais
  - [ ] **Francês**: Mercado francófono e Canadá
  - [ ] **Alemão**: Comunidade tech alemã e Europa Central
  - [ ] **Mandarim**: Gigante mercado chinês de desenvolvedores
  - [ ] **Espanhol**: Comunidade latino-americana

### Fase 3: Expansão Multi-linguagem
- [ ] **JavaScript/TypeScript**: Exemplos de anti-padrões em JS/TS
- [ ] **Python**: Código ruim famoso de bibliotecas Python antigas
- [ ] **Java**: Exemplos de código legacy problemático
- [ ] **C/C++**: Snippets de projetos famosos com problemas históricos
- [ ] **Rust**: Exemplos de código "unsafe" mal utilizado
- [ ] **PHP**: Código ruim de versões antigas do PHP

### Fase 4: Projetos Históricos Famosos
- [ ] **Linux Kernel**: Trechos problemáticos de versões antigas
- [ ] **Apache HTTP Server**: Código legacy com problemas de segurança
- [ ] **MySQL**: Implementações antigas com vazamentos de memória
- [ ] **WordPress**: Código PHP problemático de versões antigas
- [ ] **jQuery**: Exemplos de uso inadequado da biblioteca
- [ ] **Bootstrap**: CSS/JS problemático de versões iniciais


---

**Status Atual**: Fase 1 - 80% completa  
**Próximo Milestone**: Finalizar padronização dos exemplos Go  
**Meta de 2024**: Completar Fases 1 e 2, iniciar Fase 3
