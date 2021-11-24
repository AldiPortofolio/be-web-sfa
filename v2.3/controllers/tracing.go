package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"ottodigital.id/library/ottotracing"
	"time"
	zaplog "github.com/opentracing-contrib/go-zap/log"
)

func TracingFirstControllerCtx(c *gin.Context, request interface{}, namectrl string) opentracing.Span {
	var span opentracing.Span
	if cSpan, ok := c.Get("tracing-context"); ok {
		span = ottotracing.StartSpanWithParent(cSpan.(opentracing.Span).Context(), namectrl, c.Request.Method, c.Request.URL.Path)

	} else {
		span = ottotracing.StartSpanWithHeader(&c.Request.Header, c.Request.Method, namectrl, c.Request.URL.Path)
	}
	zaplog.InfoWithSpan(span, namectrl,
		zap.Any("REQ", request),
		zap.Any("Header", c.Request.Header),
		zap.Duration("backoff", time.Second))
	return span
}

func TracingEmptyFirstControllerCtx(c *gin.Context, namectrl string) opentracing.Span {
	var span opentracing.Span
	if cSpan, ok := c.Get("tracing-context"); ok {
		span = ottotracing.StartSpanWithParent(cSpan.(opentracing.Span).Context(), namectrl, c.Request.Method, c.Request.URL.Path)

	} else {
		span = ottotracing.StartSpanWithHeader(&c.Request.Header, c.Request.Method, namectrl, c.Request.URL.Path)
	}
	zaplog.InfoWithSpan(span, namectrl,
		// zap.Any("REQ", request),
		zap.Any("Header", c.Request.Header),
		zap.Duration("backoff", time.Second))
	return span
}

