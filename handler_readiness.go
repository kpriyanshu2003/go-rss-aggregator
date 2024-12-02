package main

import (
	"github.com/gin-gonic/gin"
)

func handlerReadiness(c *gin.Context) {
	respondWithJSON(c.Writer, 200, struct{}{})
}
