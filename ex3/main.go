package main

import (
	"context"
	"fib/fibonacci"
	"fib/kafka"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	if err := initGlobalTracer(nil); err != nil {
		panic(err)
	}
	/* Ex5-1 Instantiate Dynatrace Exporter - insert // on line 18 and 20
	initMetricsProvider()
	*/
	http.HandleFunc("/fib", FibServer)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":28080", nil)
}

// FibServer handles HTTP requests for fibonacci calculation
func FibServer(w http.ResponseWriter, r *http.Request) {
	/* Ex3-2 Producing SPAN on http request - insert // on line 28 and 35
	tracer := otel.Tracer("http")
	ctx := context.Background()

	var span trace.Span
	ctx, span = tracer.Start(ctx, "http-request")
	defer span.End()
	*/

	if n, err := getIntParam(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		///* Ex3-3 Execute Fibonacci function that has the trace WRAPPER function - delete // on line 40 and 42 - OLD code 
		result, numIterations := fibonacci.New().Calc(n)
		//*/

		/* Ex3-3 Execute Fibonacci function that has the trace WRAPPER function - insert // on line 44 and 46 - NEW code
		result, numIterations := fibonacci.Wrap(ctx, fibonacci.New()).Calc(n)
		*/
		reportMetric(n, numIterations)
		kafka.Send(result)
		w.Write([]byte(fmt.Sprintf("%d", result)))
	}
}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/
func hide(v interface{}) {
	kafka.PreserveImport()
	fibonacci.PreserveImport()
	context.Background()
	var tracer trace.Tracer
	tracer = otel.Tracer("")
	hide(tracer)
	hide(fmt.Sprintf("%d", 3))
}
