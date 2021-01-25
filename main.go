package main

import (
	"fmt"
	"time"
)

type memoizing interface {
	memoize()
}

//-----------------------create two struct for implementing interface(memoization)
type fibonacci struct{}

type factorial struct{}

//-----------------------create map to contain data
var fibMap map[int]int = make(map[int]int)

var facMap map[int]int = make(map[int]int)

// -------------------------------method to memoizationing fibonacci(num)
func (fi fibonacci) memoize() func(int) int {
	return func(n int) int {
		if fibMap[n] == 0 {
			fibMap[n] = fib(n)
		}
		return fibMap[n]
	}
}

// -------------------------------method to memoizationing factorial(num)
func (fa factorial) memoize() func(int) int {
	return func(n int) int {
		if facMap[n] == 0 {
			facMap[n] = fac(n)
		}
		return facMap[n]
	}
}

//-------------------------------func that counts fibonacci(num)

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//-------------------------------func that counts factorial(num)

func fac(n int) int {
	if n == 1 {
		return 1
	}
	return n * fac(n-1)
}

func main() {
	Fibonacci := fibonacci{}
	newFib := Fibonacci.memoize()

	from := time.Now()

	for i := 1; i < 20; i++ {
		fmt.Println(newFib(i))
	}

	to := time.Now()
	fmt.Println("time taken : ", to.Sub(from))

	from1 := time.Now()

	for i := 1; i < 20; i++ {
		fmt.Println(fib(i))
	}

	to1 := time.Now()

	fmt.Println("time taken : ", to1.Sub(from1))

	// test for factorial

	Factorial := factorial{}
	newFac := Factorial.memoize()

	from2 := time.Now()

	for i := 1; i < 20; i++ {
		fmt.Println(newFac(i))
	}

	to2 := time.Now()
	fmt.Println("time taken : ", to2.Sub(from2))

	from3 := time.Now()

	for i := 1; i < 20; i++ {
		fmt.Println(fac(i))
	}

	to3 := time.Now()

	fmt.Println("time taken : ", to3.Sub(from3))

}
