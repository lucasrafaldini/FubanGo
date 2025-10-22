# Análise de API Design Ruim

Este documento analisa os problemas no arquivo `ruim.go` e explica por que certas práticas de design de API são consideradas ruins.

## 1. Side-Effects em Endpoints GET
```go
func BadAPI() {
    fmt.Println("GET /users -> retorna HTML com dados sensíveis")
    fmt.Println("POST /deleteAll -> deleta tudo sem autenticação")
}
```
**Problemas:**
- Endpoints GET causam modificações de dados (side-effects);
- Viola semântica HTTP (GET deve ser idempotente e seguro);
- Cacheable por padrão pode causar mudanças não intencionais;
- Dificulta troubleshooting e debugging;
- Viola princípio de segurança da web.

**Como melhorar:**
- Usar GET apenas para leitura (ex: `GET /users` retorna lista de usuários);
- Usar POST/PUT/DELETE para operações que modificam estado (ex: `DELETE /users/:id`);
- Implementar idempotência em operações apropriadas;
- Seguir convenções RESTful ou GraphQL;
- Documentar comportamento de cada endpoint claramente.

## 2. Vazamento de Dados Sensíveis
```go
// Retorna HTML com dados sensíveis
fmt.Println("GET /users -> retorna HTML com dados sensíveis")
```
**Problemas:**
- Expõe dados sensíveis sem controle de acesso;
- Retorna mais dados do que necessário (over-fetching);
- HTML não é formato adequado para APIs modernas;
- Dificulta consumo por clientes diversos;
- Viola princípios de privacidade e segurança.

**Como melhorar:**
- Retornar apenas campos necessários (ex: DTO/View Models);
- Usar JSON como formato padrão de resposta;
- Implementar filtragem de campos por permissão de usuário;
- Mascarar ou omitir dados sensíveis (ex: senhas, tokens);
- Seguir princípio de menor privilégio no acesso a dados.

## 3. Falta de Autenticação e Autorização
```go
fmt.Println("POST /deleteAll -> deleta tudo sem autenticação")
```
**Problemas:**
- Endpoints críticos sem proteção de autenticação;
- Qualquer cliente pode executar operações destrutivas;
- Sem controle de quem faz o quê (auditoria impossível);
- Vulnerável a ataques maliciosos;
- Viola requisitos básicos de segurança.

**Como melhorar:**
- Implementar autenticação (ex: JWT, OAuth2, API keys);
- Adicionar autorização baseada em roles/permissões (RBAC);
- Validar identidade do usuário em endpoints sensíveis;
- Implementar rate limiting e throttling;
- Registrar logs de auditoria para operações críticas.

## 4. Mistura de Responsabilidades
```go
func BadHandler(req interface{}) interface{} {
    // Processa, escreve no banco e retorna resposta diretamente
    return nil
}
```
**Problemas:**
- Handler contém lógica de negócio e acesso a dados;
- Difícil de testar isoladamente;
- Viola princípio de responsabilidade única;
- Dificulta reutilização de lógica em outros contextos;
- Aumenta acoplamento entre camadas.

**Como melhorar:**
- Separar handler de lógica de negócio (ex: padrão MVC ou Clean Architecture);
- Usar camada de service para lógica de negócio;
- Usar repositories para acesso a dados;
- Handler apenas valida entrada, chama service e formata resposta;
- Implementar injeção de dependências para testabilidade.

## 5. Falta de Versionamento
```go
func BadAPI() {
    // Endpoints sem versionamento
}
```
**Problemas:**
- Mudanças na API quebram clientes existentes;
- Impossível manter compatibilidade retroativa;
- Dificulta evolução da API ao longo do tempo;
- Clientes não sabem qual versão estão consumindo;
- Rollback de mudanças problemáticas é difícil.

**Como melhorar:**
- Versionar API desde o início (ex: `/v1/users`, `/v2/users`);
- Usar versionamento em header ou URL;
- Manter múltiplas versões simultaneamente durante transição;
- Deprecar versões antigas de forma controlada;
- Documentar mudanças entre versões (changelog).

## 6. Ausência de Contratos Claros
```go
func BadHandler(req interface{}) interface{} {
    // Tipos genéricos sem estrutura definida
    return nil
}
```
**Problemas:**
- Uso de `interface{}` perde type safety;
- Contratos de entrada/saída não claros;
- Difícil para clientes saberem como usar a API;
- Propenso a erros em runtime;
- Impossível gerar documentação automática.

**Como melhorar:**
- Definir structs tipadas para request e response (ex: `type CreateUserRequest struct { Name string }`);
- Usar validação de schema (ex: JSON Schema);
- Documentar API com OpenAPI/Swagger;
- Gerar cliente automaticamente a partir de especificação;
- Validar entrada antes de processar.

## 7. Falta de Tratamento de Erros Consistente
```go
func BadHandler(req interface{}) interface{} {
    // Não retorna erros estruturados
    return nil
}
```
**Problemas:**
- Erros não retornados de forma estruturada;
- Cliente não sabe como interpretar falhas;
- Códigos HTTP inconsistentes ou incorretos;
- Mensagens de erro não amigáveis;
- Dificulta debugging do lado cliente.

**Como melhorar:**
- Usar estrutura consistente para erros (ex: `{ "error": "message", "code": "ERR_001" }`);
- Retornar códigos HTTP apropriados (400 para bad request, 401 para unauthorized, etc.);
- Incluir detalhes úteis sem expor informações sensíveis;
- Documentar possíveis erros para cada endpoint;
- Usar middleware para tratamento centralizado de erros.

## Conclusão

API mal projetada pode:
- Expor vulnerabilidades de segurança;
- Dificultar manutenção e evolução;
- Causar frustração em desenvolvedores consumidores;
- Gerar bugs difíceis de rastrear;
- Impactar negativamente a experiência do usuário.

Boas práticas incluem:
- Seguir semântica HTTP e princípios RESTful;
- Implementar autenticação e autorização adequadas;
- Separar responsabilidades em camadas;
- Versionar API desde o início;
- Documentar contratos claramente (OpenAPI);
- Tratar erros de forma consistente e informativa;
- Validar entrada e sanitizar saída;
- Proteger dados sensíveis.
