package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

// Mock DB para benchmarks que não precisam de conexão real
func mockDB() *sql.DB {
	// Retorna nil para simular cenários sem conexão real
	// Em produção, usaríamos testcontainers ou similar
	return nil
}

// Benchmark: Conexão por request vs pool
func BenchmarkBadQuery_NoPool(b *testing.B) {
	// Simula overhead de criar conexão toda vez
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// BadQuery tenta abrir nova conexão a cada chamada
		BadQuery("user:pass@tcp(localhost)/db")
	}
}

func BenchmarkGoodQuery_WithPool(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GoodQuery(db, "test")
	}
}

// Benchmark: SQL Injection (string concat vs prepared)
func BenchmarkSQLInjectionVulnerable(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SQLInjectionVulnerable(db, "John")
	}
}

func BenchmarkPreparedStatement(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GoodQuery(db, "John")
	}
}

// Benchmark: Ignorar erros vs tratar erros
func BenchmarkIgnoreErrors(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IgnoreErrors("user:pass@tcp(localhost)/db")
	}
}

// Benchmark: Transação sem rollback vs com rollback
func BenchmarkBadTransaction(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BadTransaction(db)
	}
}

func BenchmarkSafeTransaction(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SafeTransaction(db)
	}
}

// Benchmark: Sem context timeout vs com timeout
func BenchmarkNoContextTimeout(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoContextTimeout(db)
	}
}

func BenchmarkWithContextTimeout(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		stmt, err := db.PrepareContext(ctx, "SELECT * FROM users WHERE id = $1")
		if err == nil {
			stmt.Close()
		}
		cancel()
	}
}

// Benchmark: Prepared statements
func BenchmarkNoPreparedStatements(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	userIDs := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NoPreparedStatements(db, userIDs)
	}
}

// Benchmark: N+1 queries vs batch query
func BenchmarkMultipleQueries_N1Problem(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	userIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MultipleQueries(db, userIDs)
	}
}

// Benchmark: Resource leak vs proper cleanup
func BenchmarkLeakResources(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LeakResources(db)
	}
}

// Benchmark: Bad pool config vs good config
func BenchmarkBadConnectionPool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		db := BadConnectionPool("user:pass@tcp(localhost)/db")
		if db != nil {
			db.Close()
		}
	}
}

// Benchmark: SELECT * vs SELECT specific columns
func BenchmarkSelectStar(b *testing.B) {
	db := mockDB()
	if db == nil {
		b.Skip("Skipping: requires real database connection")
	}
	defer db.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectStar(db)
	}
}
