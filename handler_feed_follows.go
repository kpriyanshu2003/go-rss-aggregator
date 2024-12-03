package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowCreate(ctx *gin.Context) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
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
	feed_follow, err := cfg.DB.CreateFeedFollow(ctx.Request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create feed follows")
		return
	}

	respondWithJSON(ctx.Writer, http.StatusOK, databaseFeedFollowToFeedFollow(feed_follow))
}

func (cfg *apiConfig) handlerFeedFollowGet(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		respondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}
	dbUser := user.(database.User)

	feed_follows, err := cfg.DB.GetFeedFollows(ctx.Request.Context(), dbUser.ID)
	if err != nil {
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	respondWithJSON(ctx.Writer, http.StatusCreated, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (cfg *apiConfig) handlerFeedFollowDelete(ctx *gin.Context) {
	feedFollowID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		respondWithError(ctx.Writer, http.StatusBadRequest, "Invalid feed follow id")
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		respondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}
	dbUser := user.(database.User)

	err = cfg.DB.DeleteFeedFollows(ctx.Request.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: dbUser.ID,
	})

	if err != nil {
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}
	respondWithJSON(ctx.Writer, http.StatusOK, struct{}{})
}
