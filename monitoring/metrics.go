package monitoring

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric/instrument"

	metricApi "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

var metricMeter metricApi.Meter

var IncomingRequestReceived func()

func CreateMetricMeters(name string) {
	exporter, err := prometheus.New()

	if err != nil {
		log.Fatal(err)
	}
	provider := metric.NewMeterProvider(metric.WithReader(exporter))

	metricMeter = provider.Meter("github.com/open-telemetry/opentelemetry-go/example/prometheus")

	go serveMetrics()

	incomingRequestCounter, _ := metricMeter.Float64Counter("incoming_requests", instrument.WithDescription("all incoming requests that this system received"))

	IncomingRequestReceived = func() {
		ctx := context.Background()

		incomingRequestCounter.Add(ctx, 1)
	}
}

func serveMetrics() {
	log.Println("serving metrics at localhost:9090/metrics")
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("error serving metrics http: %v", err)
		return
	}
}
