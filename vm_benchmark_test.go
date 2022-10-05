package jsonnet

import (
	"testing"
)

func Benchmark_VM(b *testing.B) {
	snippet := `local fibonacci(n) =
if n <= 1 then
	1
else
	fibonacci(n - 1) + fibonacci(n - 2);

fibonacci(25)`
	ast, err := SnippetToAST("process.jsonnet", snippet)
	if err != nil {
		b.Fatalf("err: %#v", err)
	}

	vm := MakeVM()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		_, err := vm.Evaluate(ast)
		if err != nil {
			b.Error(err)
		}
	}
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

func BenchmarkFibNative(b *testing.B) {

	for n := 0; n < b.N; n++ {
		_ = FibonacciRecursion(25)
	}

}
