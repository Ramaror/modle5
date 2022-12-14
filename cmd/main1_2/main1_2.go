package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	go spinner(100 * time.Millisecond)
	go spinner(100 * time.Millisecond)
	n := 44
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	d := 45
	fibD := fib(d)
	fmt.Printf("\rFibonacci(%d) = %d\n", d, fibD)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
