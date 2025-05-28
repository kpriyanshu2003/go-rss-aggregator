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

type UserHandler struct {
	cfg *Config
}

func NewUserHandler(cfg *Config) *UserHandler {
	return &UserHandler{cfg: cfg}
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required,min=1,max=100"`
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, "Invalid input: "+err.Error())
		return
	}

	user, err := h.cfg.DB.CreateUser(ctx.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      req.Name,
	})
	if err != nil {
		log.Printf("Error creating user: %v", err)
		utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	utils.RespondWithJSON(ctx.Writer, http.StatusCreated, models.DatabaseUserToUser(user))
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		utils.RespondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	utils.RespondWithJSON(ctx.Writer, http.StatusOK, models.DatabaseUserToUser(dbUser))
}
