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
	/* Ex1-1 Declare Global tracer - insert // on line 15 and 19
	if err := initGlobalTracer(nil); err != nil {
		panic(err)
	}
	*/
	http.HandleFunc("/fib", FibServer)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":28080", nil)
}

// FibServer handles HTTP requests for fibonacci calculation
func FibServer(w http.ResponseWriter, r *http.Request) {
	/* Ex1-2 Producing SPAN on http request - insert // on line 27 and 34
	tracer := otel.Tracer("http")
	ctx := context.Background()

	var span trace.Span
	ctx, span = tracer.Start(ctx, "http-request")
	defer span.End()
	*/

	if n, err := getIntParam(r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		///* Ex2-4 Execute Fibonacci function requires trace context - delete // on line 39 and 41 - OLD code
		result, numIterations := fibonacci.New().Calc(n)
		//*/

		/* Ex2-4 Execute Fibonacci function requires trace context - insert // on line 43 and 45 - NEW code
		result, numIterations := fibonacci.New(ctx).Calc(n)
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
