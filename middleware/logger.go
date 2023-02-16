package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		ctx.Set("example", 12345)

		ctx.Next()

		latency := time.Since(t)
		status := ctx.Writer.Status()
		log.Printf("[Middleware logger] Request took: %v with response status: %d", latency, status)
	}
}