package database

import (
	"backend/internal/config"
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

// GetConnection returns a singleton database connection
func GetConnection(cfg config.Config) *sql.DB {
	dbOnce.Do(func() {
		dsn := cfg.DatabaseURL
		if dsn == "" {
			log.Fatal("DATABASE_URL não definida")
		}

		var err error
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Erro ao testar conexão com o banco: %v", err)
		}

		log.Println("Conexão com o banco estabelecida com sucesso")
	})

	return db
}

// CloseConnection closes the database connection
func CloseConnection() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
