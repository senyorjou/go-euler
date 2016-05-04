// Solution to problem 2: https://projecteuler.net/problem=2

package main

import "fmt"

func fib(max int) []int {
	a, b := 1, 2
	list := []int{a, b}

	for a+b < max {
		a, b = b, a+b
		list = append(list, b)
	}

	return list
}

func main() {
	fibs := fib(4 * 1000 * 1000)
	sum := 0
	for _, i := range fibs {
		if i%2 == 0 {
			sum += i
		}
	}
	fmt.Println(fibs)
	fmt.Println("Sum of fibs is:", sum)
}
