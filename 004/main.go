// Solution to problem 4: https://projecteuler.net/problem=4

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func max(list []int) (max int) {
	for _, i := range list {
		if i > max {
			max = i
		}
	}

	return max
}

func nums3dig() (list []int) {
	for i := 100; i < 1000; i++ {
		list = append(list, i)
	}

	return list
}

func reverse(str string) string {
	n := 0
	rune := make([]rune, len(str))
	for _, r := range str {
		rune[n] = r
		n++
	}

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	return string(rune)
}

func checkPalin(i, j int) (res bool) {
	number := strconv.Itoa(i * j)
	reversed := reverse(number)

	if strings.Compare(number, reversed) == 0 {
		return true
	}

	return false
}

func main() {
	list := nums3dig()
	var res []int

	for pos, i := range list {
		for _, j := range list[pos:] {
			if checkPalin(i, j) {
				res = append(res, i*j)
			}
		}
	}

	fmt.Println(max(res))
}
