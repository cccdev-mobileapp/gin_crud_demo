package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Request URL %v", ctx.Request.URL.String())
		ctx.Next()
	}
}
