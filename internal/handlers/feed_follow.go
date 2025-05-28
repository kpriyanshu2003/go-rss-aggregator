package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/models"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/utils"
)

type FeedFollowHandler struct {
	cfg *Config
}

func NewFeedFollowHandler(cfg *Config) *FeedFollowHandler {
	return &FeedFollowHandler{cfg: cfg}
}

type CreateFeedFollowRequest struct {
	FeedID uuid.UUID `json:"feed_id" binding:"required"`
}

func (h *FeedFollowHandler) CreateFeedFollow(ctx *gin.Context) {
	var req CreateFeedFollowRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	feedFollow, err := h.cfg.DB.CreateFeedFollow(ctx.Request.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		FeedID:    req.FeedID,
	})
	if err != nil {
		log.Printf("Error creating feed follow: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (h *FeedFollowHandler) GetFeedFollows(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	feedFollows, err := h.cfg.DB.GetFeedFollows(ctx.Request.Context(), dbUser.ID)
	if err != nil {
		log.Printf("Error getting feed follows: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (h *FeedFollowHandler) DeleteFeedFollow(ctx *gin.Context) {
	feedFollowIDStr := ctx.Param("id")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	err = h.cfg.DB.DeleteFeedFollows(ctx.Request.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: dbUser.ID,
	})
	if err != nil {
		log.Printf("Error deleting feed follow: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusOK, map[string]string{
		"message": "Feed follow deleted successfully",
	})
}
