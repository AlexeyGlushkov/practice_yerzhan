package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		correlationID := generateCorrelationID(start)

		c.Header("Correlation-ID", correlationID)

		log.Printf("[%s] Recieved a HTTP Request method=[%s] path=[%s]\n", correlationID, c.Request.Method, c.Request.URL.Path)

		c.Next()

		log.Printf("[%s] Request executed in %v\n", correlationID, time.Since(start))
	}
}

func generateCorrelationID(salt time.Time) string {
	return fmt.Sprintf("%d", salt.UnixNano())
}
