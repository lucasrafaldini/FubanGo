package rwtesting

import (
	"fmt"
	"testing"
	"time"
)

var globalCounter int

// 1. Testes dependentes de ordem e estado global
func TestDependsOnOrder1(t *testing.T) {
	// RUIM: depende que este rode primeiro
	globalCounter = 0
	globalCounter++
	if globalCounter != 1 {
		t.Fatal("Expected 1")
	}
}

func TestDependsOnOrder2(t *testing.T) {
	// RUIM: depende do estado deixado pelo teste anterior
	globalCounter++
	if globalCounter != 2 {
		t.Fatal("Expected 2")
	}
}

// 2. Testes lentos sem uso de mocks
func TestSlowWithoutMocks(t *testing.T) {
	// RUIM: faz chamada real à API externa
	// response := http.Get("https://api.real-service.com/data")
	// Simula demora
	time.Sleep(5 * time.Second)

	// RUIM: conecta em banco de dados real
	// db := sql.Open("postgres", "real-connection-string")
	time.Sleep(2 * time.Second)

	// RUIM: lê arquivo do disco
	// ioutil.ReadFile("/path/to/large/file.csv")
	time.Sleep(1 * time.Second)

	// Teste demora 8+ segundos quando poderia ser instantâneo com mocks
}

// 3. Uso de sleep ao invés de sincronização adequada
func TestWithBadSleep(t *testing.T) {
	done := make(chan bool)

	go func() {
		time.Sleep(100 * time.Millisecond)
		done <- true
	}()

	// RUIM: sleep arbitrário ao invés de esperar pelo canal
	time.Sleep(200 * time.Millisecond)

	// Pode falhar se goroutine demorar mais que 200ms
	// Pode ser mais lento que necessário se goroutine terminar em 50ms
}

// 4. Falta de assertions adequadas
func TestNoAssertions(t *testing.T) {
	result := sum(2, 2)

	// RUIM: não verifica nada
	_ = result

	// RUIM: apenas printa ao invés de falhar
	if result != 4 {
		println("Ops, sum está errado")
	}

	// RUIM: não usa t.Error ou t.Fatal
}

func sum(a, b int) int {
	return a + b
}

// 5. Sem setup e teardown adequados
func TestWithoutSetup(t *testing.T) {
	// RUIM: duplica código de setup em cada teste
	db := openDatabase()
	insertTestData(db)

	// teste...

	// RUIM: esquece de fazer cleanup
	// db nunca é fechado, dados de teste permanecem
}

func TestWithoutSetup2(t *testing.T) {
	// RUIM: duplica o mesmo setup novamente
	db := openDatabase()
	insertTestData(db)

	// teste...

	// RUIM: esquece cleanup novamente
}

func openDatabase() interface{} {
	return nil
}

func insertTestData(db interface{}) {}

// 6. Sem usar table-driven tests
func TestAdd1Plus1(t *testing.T) {
	if sum(1, 1) != 2 {
		t.Fatal("1+1 should be 2")
	}
}

func TestAdd2Plus2(t *testing.T) {
	if sum(2, 2) != 4 {
		t.Fatal("2+2 should be 4")
	}
}

func TestAdd5Plus3(t *testing.T) {
	if sum(5, 3) != 8 {
		t.Fatal("5+3 should be 8")
	}
}

// RUIM: código duplicado, difícil adicionar novos casos

// 7. Não executar com race detector
func TestConcurrentAccess(t *testing.T) {
	// RUIM: teste com acesso concurrent mas nunca roda com -race
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			counter++ // race condition!
		}()
	}

	time.Sleep(100 * time.Millisecond)

	// Passa nos testes mas tem race condition
	if counter < 0 {
		t.Fatal("Counter is negative")
	}
}

// 8. Não medir cobertura de testes
func TestOnlyHappyPath(t *testing.T) {
	// RUIM: só testa caminho feliz, nunca roda coverage
	result := divide(10, 2)
	if result != 5 {
		t.Fatal("Expected 5")
	}

	// Nunca testa divisão por zero
	// Nunca testa números negativos
	// Cobertura seria baixa mas não é medida
}

func divide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// 9. Mocks mal implementados
type BadMock struct {
	CallCount int
}

func (m *BadMock) Method1() string {
	m.CallCount++
	// RUIM: retorna sempre a mesma coisa independente de input
	return "mocked"
}

func (m *BadMock) Method2(id int) interface{} {
	// RUIM: ignora o parâmetro id
	return "always same"
}

func TestWithBadMock(t *testing.T) {
	mock := &BadMock{}

	// RUIM: mock não valida os argumentos recebidos
	result1 := mock.Method2(1)
	result2 := mock.Method2(999)

	// Ambos retornam a mesma coisa, não testando lógica real
	if result1 != result2 {
		t.Fatal("Mock returns different values")
	}

	// RUIM: não verifica se mock foi chamado corretamente
}

// 10. Não testar condições de erro
func TestIgnoreErrors(t *testing.T) {
	// RUIM: só testa quando tudo dá certo
	result, _ := riskyOperation(true)

	if result != "success" {
		t.Fatal("Expected success")
	}

	// Nunca testa riskyOperation(false) que retorna erro
}

func riskyOperation(succeed bool) (string, error) {
	if succeed {
		return "success", nil
	}
	return "", fmt.Errorf("failed")
}

// 11. Testes não isolados - compartilham estado
var sharedResource = map[string]string{}

func TestModifyShared1(t *testing.T) {
	// RUIM: modifica recurso compartilhado
	sharedResource["key"] = "value1"

	if sharedResource["key"] != "value1" {
		t.Fatal("Expected value1")
	}
}

func TestModifyShared2(t *testing.T) {
	// RUIM: pode falhar dependendo da ordem dos testes
	sharedResource["key"] = "value2"

	if sharedResource["key"] != "value2" {
		t.Fatal("Expected value2")
	}
}

// 12. Testes que testam implementação ao invés de comportamento
func TestImplementationDetails(t *testing.T) {
	// RUIM: testa detalhes de implementação interna
	obj := &SomeStruct{internalCounter: 0}

	obj.DoSomething()

	// RUIM: verifica estado interno ao invés de comportamento externo
	if obj.internalCounter != 1 {
		t.Fatal("Internal counter should be 1")
	}

	// Deveria testar o comportamento público, não detalhes internos
}

type SomeStruct struct {
	internalCounter int
}

func (s *SomeStruct) DoSomething() {
	s.internalCounter++
}
