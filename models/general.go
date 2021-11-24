package models

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// GeneralModel ..
type GeneralModel struct {
	ParentSpan opentracing.Span
	OttoZaplog *zap.Logger
	SpanId     string
	Context    context.Context
}

// DataModule ..
type DataModule struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

