package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/GanymedeNil/GoFrameworkBase/internal/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path
		query := context.Request.URL.RawQuery
		context.Next()

		cost := time.Since(start)
		global.LOGGER.Info(path,
			zap.Int("status", context.Writer.Status()),
			zap.String("method", context.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", context.ClientIP()),
			zap.String("user-agent", context.Request.UserAgent()),
			zap.String("errors", context.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

func Recovery(stack bool) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(context.Request, false)
				if brokenPipe {
					global.LOGGER.Error(context.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)))
					context.Error(err.(error))
					context.Abort()
					return
				}
				if stack {
					global.LOGGER.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())))
				} else {
					global.LOGGER.Error("[Recovery from panic]",
						zap.Any("error", err), zap.String("request", string(httpRequest)))
				}
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		context.Next()
	}
}
