package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/config"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/handlers"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/middleware"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/routes"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/services"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	conn, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}
	defer conn.Close()

	if err := conn.Ping(); err != nil {
		log.Fatal("Can't ping database:", err)
	}

	db := database.New(conn)

	scraperService := services.NewScraperService(db, cfg.ScrapingConcurrency, cfg.ScrapingInterval)
	go scraperService.Start()

	handlerConfig := &handlers.Config{DB: db}
	router := setupRouter(handlerConfig)
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRouter(cfg *handlers.Config) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	authMiddleware := middleware.NewAuthMiddleware(cfg.DB)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes.SetupRoutes(router, cfg, authMiddleware)

	return router
}
