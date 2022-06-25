package middleware

import (
	"gin-web/global"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		global.LOG.Info(path,
			zap.String("path", path),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("cost", cost.String()),
		)
	}
}
