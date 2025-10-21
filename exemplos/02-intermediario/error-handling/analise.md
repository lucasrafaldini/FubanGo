# Análise de Tratamento de Erros Ruins

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas de tratamento de erros são consideradas ruins.

## 1. Ignorar Erros
```go
file, _ := os.Open("arquivo.txt")
defer file.Close()
data := make([]byte, 100)
_, _ = file.Read(data)
```
**Problemas:**
- Ignora falhas críticas;
- Mascara problemas reais (ex: arquivo não existe);
- Dificulta debug e manutenção;
- Pode causar pânico em runtime devido a valores nulos;
- Comportamento imprevisível da aplicação.

**Como melhorar:**
- Sempre verificar erros retornados;
- Tratar erros apropriadamente (com logging, retry, fallback);
- Propagar erros relevantes para o chamador;
- Adicionar contexto ao erro ao retorná-lo;
- Documentar casos de erro esperados na função.

## 2. Uso Inadequado de Panic
```go
func PanicInsteadOfError(value int) int {
    if value < 0 {
        panic("valor negativo não permitido")
    }
    return value * 2
}
```
**Problemas:**
- Interrompe execução abruptamente (não controlado);
- Difícil de recuperar do estado anterior;
- Não segue o contrato da função (como retornar erro);
- Mistura fluxo de erro com fluxo normal de controle;
- Pode crashar a aplicação inteira.

**Como melhorar:**
- Retornar error em vez de panic sempre que possível
- Usar panic apenas para erros irrecuperáveis (e ainda assim fazer logging;)
- Documentar condições que causam panic claramente;
- Considerar tipos de erro personalizados para diferentes falhas;
- Manter consistência no tratamento de erros em toda a base de código.

## 3. Erros Genéricos
```go
func ReturnGenericError() error {
    return fmt.Errorf("erro")
}
```
**Problemas:**
- Mensagem não informativa;
- Impossível tratar erros específicos;
- Difícil depurar a causa raiz;
- Não fornece contexto útil;
- Não permite decisões programáticas baseadas no tipo de erro.

**Como melhorar:**
- Criar tipos de erro específicos;
- Incluir detalhes relevantes na mensagem de erro (ex: parâmetros, estado);
- Usar erro em cadeia (wrap) com %w;
- Permitir comparação de tipos de erro com errors.Is/As;
- Adicionar métodos úteis aos erros personalizados.

## 4. Perda de Contexto
```go
func LoseErrorContext(path string) error {
    _, err := os.Stat(path)
    if err != nil {
        return fmt.Errorf("falha ao verificar arquivo")
    }
    return nil
}
```
**Problemas:**
- Perde informação do erro original (ex: tipo de erro, mensagem);
- Dificulta troubleshooting (ex: arquivo não encontrado vs permissão negada);
- Impossibilita tratamento específico do erro;
- Reduz utilidade das mensagens de erro;
- Quebra a cadeia de erros útil para debugging.

**Como melhorar:**
- Usar fmt.Errorf com %w para preservar o erro original;
- Preservar erro original ao adicionar contexto;
- Adicionar contexto sem perder informação do erro original;
- Criar hierarquia de erros quando apropriado;
- Implementar interfaces de erro úteis para facilitar tratamento.

## 5. Mistura de Erros e Logs
```go
func MixErrorAndLogging(value int) (int, error) {
    if value == 0 {
        fmt.Println("Erro: divisão por zero")
        return 0, fmt.Errorf("divisão por zero")
    }
    // ...
}
```
**Problemas:**
- Mistura responsabilidades (erro vs logging);
- Dificulta controle de logs (em produção vs desenvolvimento);
- Duplica informação (em logs e erros);
- Dificulta testes e manutenção;
- Viola separação de concerns na arquitetura.

**Como melhorar:**
- Separar logging de tratamento de erro;
- Usar níveis de log apropriados (Info, Warn, Error);
- Centralizar logging em vez de espalhar pelo código;
- Estruturar logs para fácil análise (ex: JSON);
- Manter consistência no formato de logs.

## 6. Falta de Agrupamento de Erros
```go
func LoadConfig() error {
    if err := checkPermissions(); err != nil {
        return fmt.Errorf("erro de permissão: %v", err)
    }
    // ...
}
```
**Problemas:**
- Erros não relacionados são tratados da mesma forma;
- Difícil identificar origem do erro;
- Tratamento inconsistente;
- Falta de hierarquia clara de erros;
- Dificuldade em testes unitários.

**Como melhorar:**
- Criar tipos de erro específicos para diferentes categorias;
- Implementar interfaces de erro para agrupar erros relacionados;
- Agrupar erros relacionados em pacotes ou módulos;
- Usar errors.Is/As para diferenciar tipos de erro;
- Criar hierarquia clara de erros para facilitar tratamento.

## 7. Recover Indiscriminado
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recuperado de:", r)
    }
}()
```
**Problemas:**
- Recupera de qualquer panic sem distinção;
- Mascara problemas sérios no código;
- Dificulta debug de falhas reais;
- Pode deixar estado inconsistente após o panic;
- Não diferencia tipos de panic (ex: erros de lógica vs falhas críticas).

**Como melhorar:**
- Recuperar apenas panics específicos quando apropriado;
- Documentar uso de recover claramente;
- Manter estado consistente após recuperação;
- Logar informação relevante ao recuperar de um panic;
- Considerar reinicialização limpa do sistema após panics críticos.

## Conclusão

Tratamento de erros ruim pode:
- Tornar o sistema instável;
- Dificultar manutenção;
- Complicar debugging;
- Mascarar problemas reais;
- Criar comportamentos imprevisíveis.

Boas práticas incluem:
- Tratamento explícito de erros;
- Erros tipados e contextualizados;
- Hierarquia clara de erros;
- Logging apropriado;
- Recuperação controlada de panics;
- Documentação clara.