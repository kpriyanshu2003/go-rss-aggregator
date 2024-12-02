package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func init() {
	godotenv.Load()

}

func main() {
	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")

	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	if db_url == "" {
		log.Fatal("Database Url no found")
	}

	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Can't connect database", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running..."})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			handlerReadiness(ctx.Writer, ctx.Request)
		})

		v1.GET("/err", func(ctx *gin.Context) {
			handleErr(ctx)
		})

		v1.POST("/users", func(ctx *gin.Context) {
			apiCfg.handlerUsersCreate(ctx.Writer, ctx.Request)
		})

	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r.Handler(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
