package main

import (
	"fmt"
	"math"
)

func max(list []int) (max int) {
	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func decomp(a int) (list []int) {
	if a == 1 {
		return []int{1}
	}

	limit := int(math.Sqrt(float64(a))) + 1

	for check := 2; check < limit; check++ {
		for a%check == 0 {
			list = append(list, check)
			a /= check
		}
	}

	if a > 1 {
		list = append(list, a)
	}

	return list
}

func main() {
	sample := 600851475143

	res := decomp(sample)
	fmt.Println("Values:", res)
	fmt.Println("Max val:", max(res))
}
