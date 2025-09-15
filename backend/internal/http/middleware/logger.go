// Package middleware provides middleware for the http server
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// RequestLogger logs the request
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		defer func() {
			path := c.FullPath()
			if path == "" {
				path = c.Request.URL.Path
			}

			log.Info().
				Str("method", c.Request.Method).
				Str("path", path).
				Str("ua", c.Request.UserAgent()).
				Int("status", c.Writer.Status()).
				Int("size", c.Writer.Size()).
				Dur("latency", time.Since(start)).
				Msg("http_request")
		}()
	}
}
