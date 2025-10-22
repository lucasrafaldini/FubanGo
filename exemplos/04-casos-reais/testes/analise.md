# Análise de Testes Ruins

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas de testes são consideradas ruins.

## 1. Testes Dependentes de Ordem e Estado Global
```go
func BadTests(t *testing.T) {
    // Dependência de ordem e estado compartilhado
}
```
**Problemas:**
- Testes não são independentes e isolados;
- Ordem de execução afeta resultado (não determinístico);
- Falhas em um teste afetam outros testes;
- Dificulta execução paralela de testes;
- Impossível rodar testes individuais de forma confiável.

**Como melhorar:**
- Cada teste deve configurar seu próprio estado (ex: setup em cada `t.Run()`);
- Usar `t.Cleanup()` ou defer para limpar estado após teste;
- Evitar variáveis globais ou resetá-las em cada teste;
- Rodar testes com `-shuffle=on` para detectar dependências;
- Usar subtests (`t.Run`) para melhor organização e isolamento.

## 2. Testes Lentos Sem Mocks
```go
func SlowTest(t *testing.T) {
    // Simula processo demorado e não valida resultado
}
```
**Problemas:**
- Testes demoram muito tempo para executar;
- Dependem de recursos externos (rede, banco, filesystem);
- Feedback lento durante desenvolvimento;
- CI/CD fica lento e caro;
- Desenvolvedores evitam rodar testes frequentemente.

**Como melhorar:**
- Usar mocks para dependências externas (ex: interfaces mockadas);
- Separar testes unitários rápidos de testes de integração;
- Usar tags para classificar testes (ex: `//go:build integration`);
- Rodar apenas testes unitários durante desenvolvimento;
- Reservar testes de integração para CI ou execução periódica.

## 3. Uso de Sleep em Testes
```go
func BadSleepTest(t *testing.T) {
    svc := &Service{}
    go svc.Do()
    // time.Sleep(time.Second * 5) // espera sem garantia
}
```
**Problemas:**
- Sleep arbitrário não garante que operação completou;
- Testes ficam lentos desnecessariamente;
- Flaky tests (falham intermitentemente);
- Não há sincronização real entre goroutines;
- Difícil ajustar tempo de sleep (muito curto falha, muito longo é lento).

**Como melhorar:**
- Usar canais para sincronização (ex: `done := make(chan struct{}); <-done`);
- Implementar hooks ou callbacks para notificar conclusão;
- Usar `sync.WaitGroup` para aguardar goroutines;
- Implementar polling com timeout curto se necessário;
- Testar comportamento assíncrono de forma determinística.

## 4. Falta de Assertions e Validações
```go
func SlowTest(t *testing.T) {
    // Simula processo demorado e não valida resultado
}
```
**Problemas:**
- Teste não verifica se resultado está correto;
- Passa mesmo quando funcionalidade está quebrada;
- Falsa sensação de segurança;
- Não documenta comportamento esperado;
- Desperdiça tempo de execução sem valor.

**Como melhorar:**
- Usar assertions claras (ex: `if got != want { t.Errorf("got %v, want %v", got, want) }`);
- Verificar todos os aspectos relevantes do resultado;
- Usar bibliotecas de assertion quando apropriado (ex: testify);
- Documentar expectativas no teste;
- Usar table-driven tests para cobrir múltiplos casos.

## 5. Testes Não Isolados (Sem Setup/Teardown)
```go
func BadTests(t *testing.T) {
    // Não limpa estado após teste
}
```
**Problemas:**
- Estado de um teste vaza para próximo teste;
- Arquivos, conexões ou recursos não são limpos;
- Testes falham quando rodados múltiplas vezes;
- Dificulta identificar causa de falhas;
- Pode causar problemas em CI/CD.

**Como melhorar:**
- Usar `t.Cleanup()` para garantir limpeza (ex: `t.Cleanup(func() { os.Remove(tempFile) })`);
- Usar defer para cleanup imediato após criação de recursos;
- Criar funções helper de setup e teardown;
- Usar subtests para escopo de cleanup mais granular;
- Garantir limpeza mesmo em caso de falha do teste.

## 6. Ausência de Table-Driven Tests
```go
// Múltiplas funções de teste similares ao invés de table-driven
```
**Problemas:**
- Duplicação de código de teste;
- Difícil adicionar novos casos de teste;
- Inconsistência entre testes similares;
- Mais código para manter;
- Difícil visualizar cobertura de casos.

**Como melhorar:**
- Usar pattern de table-driven tests (ex: `tests := []struct{ name, input, want }{...}`);
- Iterar sobre casos de teste com `t.Run(tt.name, func(t *testing.T) {...})`;
- Facilita adição de novos casos sem duplicação;
- Melhora legibilidade e manutenibilidade;
- Torna explícito quais casos estão sendo testados.

## 7. Não Rodar com Race Detector
```go
func BadSleepTest(t *testing.T) {
    svc := &Service{}
    go svc.Do()
    // Possível race condition não detectada
}
```
**Problemas:**
- Race conditions não detectadas em código concorrente;
- Bugs sutis aparecem apenas em produção;
- Comportamento não determinístico;
- Difícil reproduzir e debugar;
- Pode causar corrupção de dados.

**Como melhorar:**
- Rodar testes com `-race` flag regularmente (ex: `go test -race ./...`);
- Integrar race detector no CI/CD;
- Usar ferramentas de análise estática (ex: `go vet`, `staticcheck`);
- Tratar warnings de race condition imediatamente;
- Testar código concorrente extensivamente.

## 8. Cobertura de Testes Não Medida
```go
// Testes existem mas não se sabe quais linhas estão cobertas
```
**Problemas:**
- Não há visibilidade de quais partes do código estão testadas;
- Caminhos críticos podem não ter cobertura;
- Difícil identificar gaps de testes;
- Refatorações podem quebrar funcionalidade não testada;
- Qualidade geral do código pode degradar.

**Como melhorar:**
- Rodar testes com `-cover` flag (ex: `go test -cover ./...`);
- Gerar relatórios de cobertura (ex: `go test -coverprofile=coverage.out`);
- Visualizar cobertura com `go tool cover -html=coverage.out`;
- Estabelecer meta de cobertura mínima (ex: 80%);
- Integrar métricas de cobertura no CI/CD.

## 9. Mocks Mal Implementados
```go
// Uso de mocks que não validam comportamento
```
**Problemas:**
- Mocks não verificam se foram chamados corretamente;
- Testes passam mas comportamento real pode estar errado;
- Interface entre componentes não é validada;
- Mudanças em contrato não são detectadas;
- Testes se tornam inúteis.

**Como melhorar:**
- Usar bibliotecas de mock que validam chamadas (ex: gomock, testify/mock);
- Verificar número de chamadas e parâmetros recebidos;
- Implementar mocks que falham se usados incorretamente;
- Usar interfaces pequenas e focadas para facilitar mocking;
- Complementar com testes de integração quando possível.

## Conclusão

Testes ruins podem:
- Dar falsa sensação de segurança;
- Atrasar ciclo de desenvolvimento;
- Dificultar refatoração e manutenção;
- Não detectar bugs reais;
- Criar frustração na equipe de desenvolvimento.

Boas práticas incluem:
- Testes unitários isolados e determinísticos;
- Uso de mocks para dependências externas;
- Evitar sleep; usar sincronização adequada;
- Assertions claras e completas;
- Setup/teardown com `t.Cleanup()`;
- Table-driven tests para casos múltiplos;
- Rodar com `-race` e `-cover` regularmente;
- Testes rápidos e confiáveis;
- Separar testes unitários de integração.
