// Solution to problem 6: https://projecteuler.net/problem=6
//

package main

import (
	"fmt"
)

func listN(n int) (list []int) {
	for i := 1; i <= n; i++ {
		list = append(list, i)
	}

	return list
}

func squareList(listIn []int) (listOut []int) {
	for _, i := range listIn {
		listOut = append(listOut, i*i)
	}

	return listOut
}

func sumList(listIn []int) (sum int) {
	for _, i := range listIn {
		sum += i
	}

	return sum
}

func main() {
	fmt.Println("Problem #6")

	list := listN(100)
	sum := sumList(list)
	sum_sq := sumList(squareList(list))

	fmt.Println("Sum of squares", sum_sq)
	fmt.Println("Square of sum", sum*sum)
	fmt.Println("Diff", sum*sum-sum_sq)

}
