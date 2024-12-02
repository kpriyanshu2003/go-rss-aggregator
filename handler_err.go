package main

import (
	"github.com/gin-gonic/gin"
)

func handleErr(c *gin.Context) {
	respondWithError(c.Writer, 500, "Something went wrong")
}
