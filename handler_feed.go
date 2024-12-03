package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedCreate(ctx *gin.Context) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		respondWithError(ctx.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		respondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	feed, err := cfg.DB.CreateFeed(ctx.Request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		Name:      params.Name,
		Url:       params.URL,
	})
	if err != nil {
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	respondWithJSON(ctx.Writer, http.StatusOK, databaseFeedToFeed(feed))
}

func (cfg *apiConfig) handlerFeedGet(ctx *gin.Context) {
	feeds, err := cfg.DB.GetFeeds(ctx.Request.Context())
	if err != nil {
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	respondWithJSON(ctx.Writer, http.StatusOK, databaseFeedsToFeeds(feeds))
}
