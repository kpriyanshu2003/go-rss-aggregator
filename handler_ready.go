package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerReadiness(ctx *gin.Context) {
	respondWithJSON(ctx.Writer, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerErr(ctx *gin.Context) {
	respondWithError(ctx.Writer, http.StatusInternalServerError, "Internal Server Error")
}

func handlerDefault(ctx *gin.Context) {
	respondWithJSON(ctx.Writer, http.StatusOK, map[string]string{"message": "API is running..."})
}
