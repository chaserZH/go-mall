package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go-mall/consts"
	"go-mall/pkg/utils/track"
)

// Jaeger 主要作用是为每个 HTTP 请求创建或关联一个 OpenTracing Span，并将 Span 存入 Gin 的上下文 (c.Set) 供后续使用。
func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("uber-trace-id") // 获取请求头中的uber-trace-id
		var span opentracing.Span
		if traceId != "" {
			var err error
			// 如果有 traceId，关联到父 Span
			// 从请求头中获取父 Span 的 traceId 和 spanId，然后创建一个新的 Span
			span, err = track.GetParentSpan(c.FullPath(), traceId, c.Request.Header)
			if err != nil {
				return
			}
		} else {
			// 如果没有 traceId，创建一个新的 Span
			span = track.StartSpan(opentracing.GlobalTracer(), c.FullPath())
		}
		// 确保在请求结束时关闭 Span
		defer span.Finish()

		// 将 Span 存入 Gin 的上下文
		c.Set(consts.SpanCTX, opentracing.ContextWithSpan(c, span))
		// 将 Span 存入 Gin 的上下文
		c.Next()
	}
}
