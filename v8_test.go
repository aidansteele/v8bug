package v8gobug

import (
	"fmt"
	"rogchap.com/v8go"
	"testing"
)

func BenchmarkV8(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ctx, _ := v8go.NewContext(nil) // creates a new V8 context with a new Isolate aka VM
		ctx.RunScript(`
			function fib(n) {
			  if (n < 2) return n;
			  return fib(n - 2) + fib(n - 1);
			}
		`, "math.js") // executes a script on the global context
		ctx.RunScript("const result = fib(35)", "main.js") // any functions previously added to the context can be called
		val, _ := ctx.RunScript("result", "value.js") // return a value in JavaScript back to Go
		fmt.Printf("addition result: %s\n", val)
	}
}
