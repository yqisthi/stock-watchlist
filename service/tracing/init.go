package tracing

import (
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	TracerController trace.Tracer
	TracerService 	trace.Tracer
	TracerRepository trace.Tracer
)

var once = new(sync.Once)

func Init() {

	once.Do(func() {
		TracerController = otel.Tracer("stock-watchlist-backend")
		TracerService = otel.Tracer("stock-watchlist-backend")
		TracerRepository = otel.Tracer("stock-watchlist-backend")
	})
}