package main

import (
	"strings"
	"testing"
)

func Test_reverse(t *testing.T) {
	sample := "Hello"
	want := "olleH"
	got := reverse(sample)

	if strings.Compare(want, got) != 0 {
		t.Errorf("reverse(%s) == %s, want %s", sample, got, want)
	}

}

func Test_checkPalin(t *testing.T) {
	num1, num2 := 99, 91 // is 9009
	want := true
	got := checkPalin(num1, num2)

	if want != got {
		t.Errorf("checkPalin(%d, %d) == %v, want %v", num1, num2, got, want)
	}

}
