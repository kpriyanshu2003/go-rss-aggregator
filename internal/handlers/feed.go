package handlers

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/models"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/utils"
)

type FeedHandler struct {
	cfg *Config
}

func NewFeedHandler(cfg *Config) *FeedHandler {
	return &FeedHandler{cfg: cfg}
}

type CreateFeedRequest struct {
	Name string `json:"name" binding:"required,min=1,max=200"`
	URL  string `json:"url" binding:"required,url"`
}

func (h *FeedHandler) CreateFeed(ctx *gin.Context) {
	var req CreateFeedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	if _, err := url.ParseRequestURI(req.URL); err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, "Invalid URL format")
		return
	}

	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	feed, err := h.cfg.DB.CreateFeed(ctx.Request.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    dbUser.ID,
		Name:      req.Name,
		Url:       req.URL,
	})
	if err != nil {
		log.Printf("Error creating feed: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}

func (h *FeedHandler) GetFeeds(ctx *gin.Context) {
	feeds, err := h.cfg.DB.GetFeeds(ctx.Request.Context())
	if err != nil {
		log.Printf("Error getting feeds: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}
