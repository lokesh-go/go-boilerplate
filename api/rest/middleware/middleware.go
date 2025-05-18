package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Keys to store in Gin context
const (
	CtxKeyReqMeta = "req_metadata"
)

// Request meta data
type RequestMetaData struct {
	TimeStamp     time.Time
	ClientIP      string
	UserAgent     string
	Method        string
	Path          string
	ContentLength int64
}

// Capture the request meta data
func CaptureRequestMetaData() gin.HandlerFunc {
	return func(c *gin.Context) {
		meta := RequestMetaData{
			TimeStamp:     time.Now().UTC(),
			ClientIP:      c.ClientIP(),
			UserAgent:     c.Request.UserAgent(),
			Method:        c.Request.Method,
			Path:          c.Request.URL.Path,
			ContentLength: c.Request.ContentLength,
		}

		// Store meta dat into context
		c.Set(CtxKeyReqMeta, meta)

		// Continue processing
		c.Next()
	}
}

// Custom logger
func CustomLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
