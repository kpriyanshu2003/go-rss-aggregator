package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kpriyanshu2003/go-rss-aggregator/internal/utils"
)

func HandleReadiness(ctx *gin.Context) {
	utils.RespondWithJSON(ctx.Writer, http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "API is ready",
	})
}

func HandleError(ctx *gin.Context) {
	utils.RespondWithError(ctx.Writer, http.StatusInternalServerError, "Internal Server Error")
}

func HandleDefault(ctx *gin.Context) {
	utils.RespondWithJSON(ctx.Writer, http.StatusOK, map[string]string{
		"message": "RSS Aggregator API is running",
		"version": "v1",
	})
}
