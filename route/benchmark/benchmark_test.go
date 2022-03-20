package benchmark

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	res := fibonacci(20)
	fmt.Println(res)
}

func TestFibonacci2(t *testing.T) {
	res := fibonacci2(20)
	fmt.Println(res)
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
	/*						동일 시간 반복 횟수    operation 1회의 수행 시간
	BenchmarkFibonacci
	BenchmarkFibonacci-8    	   41810	       28430 ns/op
	BenchmarkFibonacci2
	BenchmarkFibonacci2-8   	10378826	       114.4 ns/op
	*/
}
