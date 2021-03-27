module fib

go 1.15

require (
	github.com/Shopify/sarama v1.28.0
	github.com/dynatrace-oss/opentelemetry-metric-go v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama v0.17.0
	go.opentelemetry.io/otel v0.17.0
	go.opentelemetry.io/otel/metric v0.17.0
	go.opentelemetry.io/otel/sdk v0.17.0
	go.opentelemetry.io/otel/sdk/metric v0.17.0
	go.opentelemetry.io/otel/trace v0.17.0
)

replace github.com/dynatrace-oss/opentelemetry-metric-go => ./opentelemetry-metric-go
