package rueidis

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/codes"

	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/trace"
)

var (
	EnableTrace = false
	tracer      = otel.Tracer("redis")
)

const (
	CtxKeyCommand    = "rueidis.command_class"
	CtxKeySubCommand = "rueidis.sub_command_class"
	CtxKeyKeys       = "rueidis.keys"
)

func StartTrace(ctx context.Context, spanName string, kvs ...string) (context.Context, func(error)) {
	var span trace.Span
	log.Println("StartTrace==>", spanName)
	if EnableTrace {
		var attrs []attribute.KeyValue
		var key string
		for idx, str := range kvs {
			if idx%2 == 0 {
				key = str
			} else {
				attrs = append(attrs, attribute.String(key, str))
			}
		}
		if val := ctx.Value(CtxKeyCommand); val != nil {
			attrs = append(attrs, attribute.String(CtxKeyCommand, val.(string)))
		}
		if val := ctx.Value(CtxKeySubCommand); val != nil {
			attrs = append(attrs, attribute.String(CtxKeySubCommand, val.(string)))
		}
		if val := ctx.Value(CtxKeyKeys); val != nil {
			attrs = append(attrs, attribute.StringSlice(CtxKeyKeys, val.([]string)))
		}
		ctx, span = tracer.Start(ctx, spanName,
			trace.WithAttributes(attrs...),
			trace.WithSpanKind(trace.SpanKindServer))
	}
	return ctx, func(err error) {
		if span != nil {
			if err != nil {
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}
	}
}
