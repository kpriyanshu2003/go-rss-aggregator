package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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
	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 5, time.Minute)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.NoRoute(handlerDefault)

	v1 := r.Group("/v1")
	{
		v1.GET("/health", handlerReadiness)
		v1.GET("/err", handlerErr)

		v1.GET("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))
		v1.POST("/users", apiCfg.handlerUsersCreate)

		v1.GET("/feed", apiCfg.handlerFeedGet)
		v1.POST("/feed", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))

		v1.GET("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowGet))
		v1.POST("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
		v1.DELETE("/feed-follow/:id", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

		v1.GET("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))
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
