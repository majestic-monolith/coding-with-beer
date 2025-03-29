package main

import (
	"backend/internal/auth"
	"backend/internal/config"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"backend/internal/database"
	"backend/internal/database/generated"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	db := database.GetConnection(cfg)
	defer database.CloseConnection()

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

	authGroup := r.Group("/auth")

	authGroup.POST("/register", func(c *gin.Context) {
		auth.Register(c, queries)
	})

	authGroup.POST("/login", func(c *gin.Context) {
		auth.Login(c, queries)
	})

	r.Run(":8080")
}
