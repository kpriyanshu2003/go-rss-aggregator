package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/models"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/utils"
)

type PostHandler struct {
	cfg *Config
}

func NewPostHandler(cfg *Config) *PostHandler {
	return &PostHandler{cfg: cfg}
}

func (h *PostHandler) GetPostsForUser(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	dbUser := user.(database.User)
	posts, err := h.cfg.DB.GetPostsForUsers(ctx.Request.Context(), database.GetPostsForUsersParams{
		UserID: dbUser.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		log.Printf("Error getting posts for user: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't get posts")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"posts": models.DatabasePostsToPosts(posts),
		"count": len(posts),
	})
}
