package logger

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

var logger *zap.Logger

func InitLogger() {
	logger = zap.Must(zap.NewProduction())
}

func Info(ctx context.Context, msg string, zapField ...zap.Field) {
	_, trace := otelgrpc.Extract(ctx, &metadata.MD{})
	logger.Info(msg, zap.Any("trace_id", trace.TraceID()), zap.Any("span_id", trace.SpanID()))
}

func Error(ctx context.Context, msg string, zapField ...zap.Field) {
	_, trace := otelgrpc.Extract(ctx, &metadata.MD{})
	logger.Error(msg, zap.Any("trace_id", trace.TraceID()), zap.Any("span_id", trace.SpanID()))
}

func Warn(ctx context.Context, msg string, zapField ...zap.Field) {
	_, trace := otelgrpc.Extract(ctx, &metadata.MD{})
	logger.Warn(msg, zap.Any("trace_id", trace.TraceID()), zap.Any("span_id", trace.SpanID()))
}

func Fatal(ctx context.Context, msg string, zapField ...zap.Field) {
	_, trace := otelgrpc.Extract(ctx, &metadata.MD{})
	logger.Fatal(msg, zap.Any("trace_id", trace.TraceID()), zap.Any("span_id", trace.SpanID()))
}

func Panic(ctx context.Context, msg string, zapField ...zap.Field) {
	_, trace := otelgrpc.Extract(ctx, &metadata.MD{})
	logger.Panic(msg, zap.Any("trace_id", trace.TraceID()), zap.Any("span_id", trace.SpanID()))
}