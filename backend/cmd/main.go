package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"backend/internal/database/generated"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL não definida")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	queries := generated.New(db)

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		users, err := queries.GetUserByID(context.Background(), 1)
		if err != nil {
			log.Printf("Erro ao buscar usuários: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuários"})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.Run(":8080")
}