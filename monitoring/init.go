package monitoring

import (
	"fmt"
	"log"
)

func ConfigureEnv(name string) {
	log.SetPrefix(fmt.Sprintf("[Goob - %s]", name))

	CreateTracer(TracerConfig{
		Name:    name,
		Version: "v0.0.0",
		Env:     "dev",
	})

	CreateMetricMeters(name)
}
