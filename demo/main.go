package main

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	usdt "opentelemtry-go-usdt"
)

func initTracer() (*sdktrace.TracerProvider, error) {
	// Create stdout exporter
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	// Create TracerProvider with USDT and stdout processors
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSpanProcessor(usdt.New()),
	)

	// Set global TracerProvider
	otel.SetTracerProvider(tp)

	return tp, nil
}

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	tracer := tp.Tracer("demo-tracer")

	ctx, span := tracer.Start(context.Background(), "parent")
	defer span.End()

	// Add some attributes to the span
	span.SetAttributes(attribute.String("foo", "bar"))

	// Simulate some work
	time.Sleep(100 * time.Millisecond)

	// Create a child span
	_, childSpan := tracer.Start(ctx, "child")
	childSpan.SetAttributes(attribute.Int("count", 42))
	time.Sleep(50 * time.Millisecond)
	childSpan.End()

	log.Println("Traces completed. Check your USDT probes and stdout for output.")
}
