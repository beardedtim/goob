package http

import (
	"mckp/goob/monitoring"
	"mckp/goob/utils"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type ServerHeader struct {
	Key   string
	Value string
}

var serverHeaders = []ServerHeader{
	{Key: "X-Server", Value: "Goob"},
	{Key: "X-Version", Value: "0.0.0"},
}

var REQUEST_ID_KEY = "request_id"

func RequestId(ctx *gin.Context) {
	_, exists := ctx.Get(REQUEST_ID_KEY)

	if exists {
		ctx.Next()
		return
	}

	uuid := utils.UUID()

	ctx.Set(REQUEST_ID_KEY, uuid)
	ctx.Header("X-Request-Id", uuid)

	tracer, _ := monitoring.GetTracer()

	_, span := tracer.Start(ctx.Request.Context(), "with-request-id", oteltrace.WithAttributes(attribute.String("request_id", uuid)))

	defer span.End()
	ctx.Next()
}

func CustomHeaders(ctx *gin.Context) {

	for _, header := range serverHeaders {
		ctx.Header(header.Key, header.Value)
	}

	ctx.Next()
}

func TrackIncomingRequestCount(ctx *gin.Context) {
	monitoring.IncomingRequestReceived()

	ctx.Next()
}
