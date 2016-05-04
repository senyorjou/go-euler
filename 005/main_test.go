package main

import (
	"testing"
)

func Test_isDivBy(t *testing.T) {
	var a, b int
	var want, got bool

	a, b = 100, 20
	want, got = true, isDivBy(a, b)

	if want != got {
		t.Errorf("isDivBy(%d, %d) == %v, want %v", a, b, got, want)
	}

	a, b = 100, 15
	want, got = false, isDivBy(a, b)

	if want != got {
		t.Errorf("isDivBy(%d, %d) == %v, want %v", a, b, got, want)
	}
}
