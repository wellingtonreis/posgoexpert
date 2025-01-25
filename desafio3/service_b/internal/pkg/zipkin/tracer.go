package zipkin

import (
	"log"

	"github.com/gofiber/contrib/otelfiber"
	fiber "github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	resource "go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracer(url string, serviceName string) (*sdktrace.TracerProvider, error) {
	exporter, err := zipkin.New(url)
	if err != nil {
		log.Printf("Zipkin exporter error service: %v", err)
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

func SetupOTelMiddleware(app *fiber.App, serviceName string) *sdktrace.TracerProvider {
	tracerProvider, err := InitTracer("http://zipkin:9411/api/v2/spans", serviceName)
	if err != nil {
		log.Fatalf("failed to initialize zipkin exporter: %v", err)
	}

	app.Use(otelfiber.Middleware(otelfiber.WithTracerProvider(tracerProvider)))

	return tracerProvider
}
