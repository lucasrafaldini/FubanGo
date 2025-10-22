package apidesign

import (
	"net/http/httptest"
	"strings"
	"testing"
)

// Benchmark: Side-effects em GET
func BenchmarkBadGetEndpoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/delete", nil)
		w := httptest.NewRecorder()
		BadGetEndpoint(w, req)
	}
}

func BenchmarkGoodGetEndpoint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/users/1", nil)
		w := httptest.NewRecorder()
		GetUserHandler(w, req)
	}
}

// Benchmark: Vazamento de dados
func BenchmarkLeakSensitiveData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		LeakSensitiveData(w, req)
	}
}

func BenchmarkSafeUserData(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/user", nil)
		w := httptest.NewRecorder()
		GetUserHandler(w, req)
	}
}

// Benchmark: Handler com lógica de negócio
func BenchmarkBadHandler_MixedResponsibilities(b *testing.B) {
	body := `{"name": "John"}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/user", strings.NewReader(body))
		w := httptest.NewRecorder()
		BadHandler(w, req)
	}
}

func BenchmarkGoodHandler_Separated(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetUserService(1)
	}
}

// Benchmark: Validação de entrada
func BenchmarkNoInputValidation(b *testing.B) {
	body := `{"name": "John", "age": 30}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/user", strings.NewReader(body))
		w := httptest.NewRecorder()
		NoInputValidation(w, req)
	}
}

// Benchmark: Erros inconsistentes
func BenchmarkInconsistentErrors_NoAuth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/resource", nil)
		w := httptest.NewRecorder()
		InconsistentErrors(w, req)
	}
}

func BenchmarkInconsistentErrors_WrongMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/resource", nil)
		req.Header.Set("Auth", "token")
		w := httptest.NewRecorder()
		InconsistentErrors(w, req)
	}
}

// Benchmark: Códigos HTTP errados
func BenchmarkWrongStatusCodes_InvalidJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("POST", "/user", strings.NewReader("invalid json"))
		w := httptest.NewRecorder()
		WrongStatusCodes(w, req)
	}
}

// Benchmark: God Endpoint
func BenchmarkGodEndpoint_Create(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api?action=create", nil)
		w := httptest.NewRecorder()
		GodEndpoint(w, req)
	}
}

func BenchmarkGodEndpoint_Update(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api?action=update", nil)
		w := httptest.NewRecorder()
		GodEndpoint(w, req)
	}
}

func BenchmarkGodEndpoint_Delete(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/api?action=delete", nil)
		w := httptest.NewRecorder()
		GodEndpoint(w, req)
	}
}
