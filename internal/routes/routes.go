package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/handlers"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/middleware"
)

func SetupRoutes(router *gin.Engine, cfg *handlers.Config, authMiddleware *middleware.AuthMiddleware) {
	router.NoRoute(handlers.HandleDefault)

	v1 := router.Group("/v1")
	{
		v1.GET("/health", handlers.HandleReadiness)
		v1.GET("/err", handlers.HandleError)

		userHandler := handlers.NewUserHandler(cfg)
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users", authMiddleware.RequireAuth(), userHandler.GetUser)

		feedHandler := handlers.NewFeedHandler(cfg)
		v1.GET("/feed", feedHandler.GetFeeds)
		v1.POST("/feed", authMiddleware.RequireAuth(), feedHandler.CreateFeed)

		feedFollowHandler := handlers.NewFeedFollowHandler(cfg)
		v1.GET("/feed-follow", authMiddleware.RequireAuth(), feedFollowHandler.GetFeedFollows)
		v1.POST("/feed-follow", authMiddleware.RequireAuth(), feedFollowHandler.CreateFeedFollow)
		v1.DELETE("/feed-follow/:id", authMiddleware.RequireAuth(), feedFollowHandler.DeleteFeedFollow)

		postHandler := handlers.NewPostHandler(cfg)
		v1.GET("/posts", authMiddleware.RequireAuth(), postHandler.GetPostsForUser)
	}
}
