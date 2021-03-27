package fibonacci

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/label"
	"go.opentelemetry.io/otel/trace"
)

// Fibonacci calculates fibonacci numbers
type Fibonacci interface {
	Calc(n int) (int, int) // Calc has no documentation
}

///* Ex2-1 Passing the trace context into Fibonacci function - delete // on line 17 and 22 - OLD code
// New creates a new Fibonacci Calculator
func New() Fibonacci {
	return &fibonacci{}
}
//*/

/* Ex2-1 Passing the trace context into Fibonacci function - insert // on line 24 and 29 - NEW code
// New creates a new Fibonacci Calculator
func New(ctx context.Context) Fibonacci {
	return &fibonacci{Context: ctx}
}
*/

///* Ex2-1 Passing the trace context into Fibonacci function - delete // on line 31 and 33 - OLD code
type fibonacci struct{}
//*/

/* Ex2-1 Passing the trace context into Fibonacci function - insert // on line 35 and 39 - NEW code
type fibonacci struct {
	Context context.Context
}
*/

// Calc calculates the n-th fibonacci number
// The first return value is the fibonacci number to be calculated
// The second return value reports the number of recursive invocation that were required in order to calculate the result
func (f *fibonacci) Calc(n int) (int, int) {

/* Ex2-2 Producing SPAN inside Calc method - insert // on line 46 and 51
	var span trace.Span
	tracer := otel.Tracer("")
	f.Context, span = tracer.Start(f.Context, fmt.Sprintf("fib(%d)", n))
	defer span.End()
*/
	var result int
	var numIterations int
	if n < 3 {
		result = 1
		numIterations = 1
	} else {
		resultA, numIterationsA := f.Calc(n - 1)
		resultB, numIterationsB := f.Calc(n - 2)
		result = resultA + resultB
		numIterations = numIterationsA + numIterationsB + 1
	}

	/* Ex2-3 Creating a SPAN attribute
	span.SetAttributes(label.Key("fib.result").Int(result))
	*/
	
	return result, numIterations
}

/*
DO NOT REMOVE ANY TEXT BELOW THIS LINE






















*/

// PreserveImport is never getting called.
// In order to keep this session as simple as possible it's purpose is to preserve imports on top of this file
func PreserveImport() {
	lbl := label.Key("")
	tracer := otel.Tracer("")
	var ctx context.Context
	var span trace.Span
	fmt.Printf("%v, %v, %v, %v", span, ctx, tracer, lbl)
}
