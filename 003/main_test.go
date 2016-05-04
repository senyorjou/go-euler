package main

import (
	"reflect"
	"testing"
)

func Test_decomp(t *testing.T) {
	var got, want []int

	got = decomp(10)
	want = []int{2, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("decomp(%d) == %v, want %v", 10, got, want)
	}

	got = decomp(2 * 3 * 7 * 11 * 13)
	want = []int{2, 3, 7, 11, 13}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("decomp(%d) == %v, want %v", 10, got, want)
	}
}

func Test_max(t *testing.T) {
	var got, want int
	var sample []int

	sample = []int{1, 2, 3, 10}
	got = max(sample)
	want = 10

	if got != want {
		t.Errorf("max(%v) == %d, want %d", sample, got, want)
	}

	sample = []int{300, 2000, 8888, -10}
	got = max(sample)
	want = 8888

	if got != want {
		t.Errorf("max(%v) == %d, want %d", sample, got, want)
	}
}
