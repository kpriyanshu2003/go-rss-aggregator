package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(ctx *gin.Context) {
	type parameters struct {
		Name string `json:"name"`
	}

	var params parameters
	if err := ctx.ShouldBindJSON(&params); err != nil {
		respondWithError(ctx.Writer, http.StatusBadRequest, "Invalid input")
		return
	}

	user, err := cfg.DB.CreateUser(ctx.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		log.Println(err)
		respondWithError(ctx.Writer, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(ctx.Writer, http.StatusOK, databaseUserToUser(user))
}

func (cfg *apiConfig) handlerUsersGet(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		respondWithError(ctx.Writer, http.StatusUnauthorized, "User not found in context")
		return
	}

	dbUser := user.(database.User)
	respondWithJSON(ctx.Writer, http.StatusOK, databaseUserToUser(dbUser))
}
