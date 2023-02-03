package monitoring

import (
	"context"
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var sdkTracer *sdktrace.TracerProvider
var tracer oteltrace.Tracer

type TracerConfig struct {
	Name    string
	Version string
	Env     string
}

func CreateTracer(config TracerConfig) error {
	tracer = otel.Tracer(config.Name)

	exporter, err := stdout.New(stdout.WithPrettyPrint())

	if err != nil {
		return err
	}

	resource, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.Name),
			semconv.ServiceVersionKey.String(config.Version),
			attribute.String("environment", config.Env),
		),
	)

	sdkTracer = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource),
	)

	otel.SetTracerProvider(sdkTracer)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return nil
}

func GetSDKTracer() (*sdktrace.TracerProvider, error) {
	if sdkTracer == nil {
		return nil, errors.New("cannot get a tracer before it is created; please call CreateTracer before calling GetTracer")
	}

	return sdkTracer, nil
}

func GetTracer() (oteltrace.Tracer, error) {
	if tracer == nil {
		return nil, errors.New("cannot get a tracer before it is created; please call CreateTracer before calling GetTracer")
	}

	return tracer, nil
}

func ShutDownTracer() {
	defer func() {
		if err := sdkTracer.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
}

func WrapMiddleware(fn gin.HandlerFunc, name string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, span := tracer.Start(ctx.Request.Context(), name)
		defer span.End()
		fn(ctx)
	}
}
