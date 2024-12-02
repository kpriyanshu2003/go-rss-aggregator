package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(c *gin.Context) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(c.Request.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(c.Writer, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(c.Request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: (time.Now().UTC()),
		UpdatedAt: (time.Now().UTC()),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(c.Writer, 400, fmt.Sprintf("Couldn't Create User: %v", err))
		return
	}

	respondWithJSON(c.Writer, 200, user)
}
