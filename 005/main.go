// Solution to problem 5: https://projecteuler.net/problem=5
//

package main

import (
	"fmt"
)

func main() {
	last := 20
	num := 10
	for {
		for j := 20; j > 0; j-- {
			last = j
			if num%j != 0 {
				break
			}
		}
		if last == 1 {
			break
		}
		num++
	}

	fmt.Println(num)
}
