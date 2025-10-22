# Análise de Uso de Database Ruim

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas de uso de banco de dados são consideradas ruins.

## 1. Abrir e Fechar Conexão Por Requisição
```go
func BadQuery(dbURL string) {
    db, _ := sql.Open("postgres", dbURL)
    defer db.Close()
    
    query := "SELECT * FROM users WHERE name = '" + "'" + "'"
    db.Query(query)
}
```
**Problemas:**
- Overhead de criar conexão para cada operação;
- Desperdiça recursos do sistema e banco de dados;
- Performance drasticamente reduzida;
- Limite de conexões pode ser atingido rapidamente;
- Latência alta por estabelecimento de conexão.

**Como melhorar:**
- Criar `sql.DB` uma vez na inicialização da aplicação (ex: `var db *sql.DB` global ou injetado);
- Reutilizar pool de conexões gerenciado automaticamente por `sql.DB`;
- Configurar tamanho do pool adequadamente (ex: `db.SetMaxOpenConns(25)`);
- Nunca fechar `sql.DB` exceto no shutdown da aplicação;
- Usar context com timeout para queries individuais.

## 2. SQL Injection por Concatenação de Strings
```go
query := "SELECT * FROM users WHERE name = '" + "'" + "'"
db.Query(query)
```
**Problemas:**
- Vulnerável a SQL injection (ataque de segurança crítico);
- Entrada maliciosa pode executar comandos arbitrários no banco;
- Possível vazamento ou destruição de dados;
- Viola princípios básicos de segurança;
- Pode permitir escalação de privilégios.

**Como melhorar:**
- Usar queries parametrizadas (ex: `db.Query("SELECT * FROM users WHERE name = $1", userName)`);
- Usar prepared statements para queries repetidas (ex: `stmt, _ := db.Prepare("SELECT * FROM users WHERE id = $1")`);
- Nunca concatenar input do usuário diretamente em queries;
- Usar ORMs que sanitizam entrada automaticamente (ex: GORM, sqlx);
- Validar e sanitizar entrada antes de usar em queries.

## 3. Ignorar Erros de Operações de Banco
```go
db, _ := sql.Open("postgres", dbURL)
defer db.Close()

query := "SELECT * FROM users WHERE name = '" + "'" + "'"
db.Query(query)
```
**Problemas:**
- Erros de conexão ignorados (banco pode estar offline);
- Erros de query ignorados (sintaxe SQL incorreta, permissões);
- Comportamento imprevisível quando operação falha;
- Dificulta debugging e troubleshooting;
- Dados podem estar corrompidos sem detecção.

**Como melhorar:**
- Sempre verificar erros retornados (ex: `if err != nil { return err }`);
- Logar erros com contexto adequado para debugging;
- Retornar erros para camadas superiores tratarem;
- Usar defer para garantir cleanup mesmo em erro (ex: `defer rows.Close()`);
- Implementar retry logic para erros transientes quando apropriado.

## 4. Transação Sem Rollback em Caso de Erro
```go
func BadTransaction(db *sql.DB) {
    tx, _ := db.Begin()
    tx.Exec("INSERT INTO users(name) VALUES('x')")
    // esquece tx.Rollback() em caso de erro
    tx.Commit()
}
```
**Problemas:**
- Transação commitada mesmo se operação falhar;
- Estado inconsistente no banco de dados;
- Violação de garantias ACID;
- Difícil rastrear quando dados ficam corrompidos;
- Pode causar locks prolongados no banco.

**Como melhorar:**
- Usar defer para rollback automático (ex: `defer tx.Rollback()` - é seguro mesmo após commit);
- Verificar erros de cada operação na transação;
- Commit apenas se todas operações tiverem sucesso;
- Usar padrão de transação segura (ex: função helper que garante rollback);
- Considerar usar bibliotecas que facilitam transações (ex: `sqlx.Tx`).

## 5. Falta de Context com Timeout
```go
db.Query(query)
tx.Exec("INSERT INTO users(name) VALUES('x')")
```
**Problemas:**
- Queries podem bloquear indefinidamente;
- Sem controle de cancelamento de operações longas;
- Recursos ficam presos em queries lentas;
- Impossível implementar timeout adequado;
- Dificulta shutdown graceful da aplicação.

**Como melhorar:**
- Usar métodos com context (ex: `db.QueryContext(ctx, query)`);
- Criar context com timeout apropriado (ex: `ctx, cancel := context.WithTimeout(ctx, 5*time.Second); defer cancel()`);
- Propagar context de requisições HTTP para queries de banco;
- Implementar deadlines diferentes para operações diferentes (read vs write);
- Monitorar queries longas e otimizar.

## 6. Não Usar Prepared Statements
```go
db.Query("SELECT * FROM users WHERE name = '" + userName + "'")
```
**Problemas:**
- Query precisa ser parseada toda vez;
- Performance reduzida para queries repetidas;
- Maior carga no banco de dados;
- Vulnerável a SQL injection;
- Desperdiça recursos de CPU e memória.

**Como melhorar:**
- Preparar statements que serão reutilizados (ex: `stmt, err := db.Prepare("SELECT * FROM users WHERE id = $1")`);
- Cache de prepared statements para queries frequentes;
- Usar bibliotecas que gerenciam prepared statements automaticamente;
- Fechar statements quando não mais necessários (ex: `defer stmt.Close()`);
- Balancear uso de prepared statements vs queries ad-hoc.

## 7. Ausência de Migrations e Versionamento de Schema
```go
// Código não demonstra, mas comum em projetos sem migrations
```
**Problemas:**
- Schema de banco não versionado junto com código;
- Difícil sincronizar estrutura entre ambientes;
- Rollback de mudanças de schema é manual e propenso a erros;
- Documentação de mudanças de schema inexistente;
- Deploy arriscado por inconsistências de schema.

**Como melhorar:**
- Usar ferramenta de migrations (ex: golang-migrate, goose);
- Versionar migrations junto com código da aplicação;
- Aplicar migrations automaticamente no deploy;
- Testar migrations em ambiente de staging antes de produção;
- Manter migrations idempotentes e reversíveis quando possível.

## Conclusão

Uso inadequado de banco de dados pode:
- Causar vulnerabilidades críticas de segurança (SQL injection);
- Degradar severamente a performance da aplicação;
- Gerar inconsistências e corrupção de dados;
- Dificultar manutenção e troubleshooting;
- Impactar negativamente a experiência do usuário.

Boas práticas incluem:
- Reutilizar pool de conexões (`sql.DB`);
- Sempre usar queries parametrizadas;
- Verificar e tratar todos os erros;
- Implementar transações com rollback adequado;
- Usar context com timeout em todas as operações;
- Preparar statements para queries repetidas;
- Versionar schema com migrations;
- Monitorar performance e otimizar queries lentas.
