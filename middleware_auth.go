package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/auth"
)

func (cfg *apiConfig) middlewareAuth(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey, err := auth.GetAPIKey(ctx.Request.Header)
		if err != nil {
			respondWithError(ctx.Writer, http.StatusForbidden, fmt.Sprintf("Auth Error: %v", err))
			ctx.Abort()
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(ctx.Request.Context(), apikey)
		if err != nil {
			respondWithError(ctx.Writer, http.StatusBadRequest, fmt.Sprintf("Couldn't get user: %v", err))
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		handler(ctx)
	}
}
