package main

import (
	"backend/internal/config"
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"

	"backend/internal/database/generated"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	dsn := cfg.DatabaseURL

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
