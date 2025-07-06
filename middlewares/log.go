package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // 🕒 Start timer
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next() // ⏭️ Process request

		duration := time.Since(start) // ⏱️ Calculate duration
		status := c.Writer.Status()   // ✅ Status code

		log.Printf("[GIN] %s %s %d %s\n", method, path, status, duration)
	}
}
