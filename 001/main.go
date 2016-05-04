// Solution to problem 1: https://projecteuler.net/problem=1

package main

import (
	"fmt"
)

func listAdd(list []int) (retval int) {
	retval = 0
	for _, value := range list {
		retval += value
	}

	return
}

func listSet(lists ...[]int) (list []int) {
	type IntSet map[int]bool

	set := make(IntSet)

	for _, list := range lists {
		for _, item := range list {
			set[item] = true
		}
	}

	list = make([]int, len(set))
	for item := range set {
		list = append(list, item)
	}

	return
}

func dividersByN(n, max int) (divs []int) {
	for i := n; i < max; i++ {
		if i%n == 0 {
			divs = append(divs, int(i))
		}
	}

	return
}

func main() {
	set := listSet(dividersByN(3, 1000), dividersByN(5, 1000))
	fmt.Println(listAdd(set))
}
