package series

import "fmt"

func init() {
	fmt.Println("Init 1")
}

func init() {
	fmt.Println("Init 2")
}

func Square(n int) int {
	return n * n
}

// The Fibonacci Func
func GerFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
