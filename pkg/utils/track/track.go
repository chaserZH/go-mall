package track

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"net/http"
)

// GetParentSpan 分布式追踪
func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext))
	if err != nil {
		return nil, err
	}
	return parentSpan, err
}

// StartSpan 用于生成一个顶级 Span（没有父 Span 的根 Span）。
func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	// 设置顶级span
	span := tracer.StartSpan(name)
	return span
}
