package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// 1. Abre e fecha conexão a cada requisição
func BadQuery(dbURL string) {
	db, _ := sql.Open("postgres", dbURL)
	defer db.Close()

	// Concatenação de strings para query (SQL injection)
	query := "SELECT * FROM users WHERE name = '" + "admin" + "'"
	db.Query(query)
}

// 2. Concatenação de strings - SQL Injection
func SQLInjectionVulnerable(db *sql.DB, userName string) {
	// RUIM: concatenação direta permite SQL injection
	query := "SELECT * FROM users WHERE name = '" + userName + "'"
	rows, _ := db.Query(query)
	defer rows.Close()

	// Se userName = "admin' OR '1'='1" retorna todos os usuários
}

// 3. Ignorar erros de operações
func IgnoreErrors(dbURL string) {
	db, _ := sql.Open("postgres", dbURL) // erro ignorado
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM users") // erro ignorado
	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name) // erro ignorado
		fmt.Println(name)
	}
	// rows.Err() não verificado
}

// 4. Transação sem rollback em erro
func BadTransaction(db *sql.DB) {
	tx, _ := db.Begin()

	// Primeira operação
	_, err := tx.Exec("INSERT INTO users(name) VALUES('user1')")
	if err != nil {
		// RUIM: não faz rollback, apenas retorna
		return
	}

	// Segunda operação pode falhar
	tx.Exec("INSERT INTO orders(user_id) VALUES(999)")

	// RUIM: commit mesmo se segunda operação falhou
	tx.Commit()
}

// 5. Falta de context com timeout
func NoContextTimeout(db *sql.DB) {
	// Query sem context pode bloquear indefinidamente
	rows, _ := db.Query("SELECT * FROM large_table WHERE complex_condition = true")
	defer rows.Close()

	for rows.Next() {
		// processa dados
	}
}

// 6. Não usar prepared statements
func NoPreparedStatements(db *sql.DB, userIDs []int) {
	for _, id := range userIDs {
		// RUIM: query é parseada toda vez
		query := fmt.Sprintf("SELECT * FROM users WHERE id = %d", id)
		db.Query(query)
	}
}

// 7. Múltiplas queries quando poderia ser uma
func MultipleQueries(db *sql.DB, userIDs []int) {
	for _, id := range userIDs {
		// RUIM: N+1 queries ao invés de uma única query
		db.Query("SELECT * FROM users WHERE id = ?", id)
	}
}

// 8. Não fechar recursos
func LeakResources(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM users")
	// RUIM: esquece de fechar rows (vazamento de conexão)

	for rows.Next() {
		var name string
		rows.Scan(&name)
	}
	// rows nunca fechado
}

// 9. Pool de conexões mal configurado
func BadConnectionPool(dbURL string) *sql.DB {
	db, _ := sql.Open("postgres", dbURL)

	// RUIM: não configura limites do pool
	// Pode esgotar conexões do banco ou usar recursos excessivos
	return db
}

// 10. Usar SELECT * ao invés de campos específicos
func SelectStar(db *sql.DB) {
	// RUIM: retorna todas as colunas mesmo precisando apenas de algumas
	rows, _ := db.Query("SELECT * FROM users")
	defer rows.Close()

	for rows.Next() {
		var name string
		// Precisa apenas do nome mas carrega todas as colunas
		rows.Scan(&name)
	}
}
