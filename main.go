package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
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
			handlerReadiness(ctx)
		})

		v1.GET("/err", func(ctx *gin.Context) {
			handleErr(ctx)
		})

	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r.Handler(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
