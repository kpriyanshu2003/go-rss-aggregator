package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/auth"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/database"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/utils"
)

type AuthMiddleware struct {
	db *database.Queries
}

func NewAuthMiddleware(db *database.Queries) *AuthMiddleware {
	return &AuthMiddleware{db: db}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiKey, err := auth.GetAPIKey(ctx.Request.Header)
		if err != nil {
			utils.RespondWithError(ctx.Writer, http.StatusForbidden, fmt.Sprintf("Auth Error: %v", err))
			ctx.Abort()
			return
		}

		user, err := m.db.GetUserByAPIKey(ctx.Request.Context(), apiKey)
		if err != nil {
			utils.RespondWithError(ctx.Writer, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
