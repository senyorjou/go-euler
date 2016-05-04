// Solution to problem 7: https://projecteuler.net/problem=7
//

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n%2 == 0 && n > 2 {
		return false
	}

	for i := 3; i < int(math.Sqrt(float64(n)))+1; i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primeList(max int) (list []int) {

	for i := 2; len(list) < max; i++ {
		if isPrime(i) {
			list = append(list, i)
		}
	}

	return list

}

func main() {
	fmt.Println("Problem #7")
	pos := 10001
	fmt.Println("Prime in pos", pos, ":", primeList(pos)[pos-1])

}
