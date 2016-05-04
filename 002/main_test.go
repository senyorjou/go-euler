package main

import "testing"

import "reflect"

func Test_fib(t *testing.T) {
	got := fib(6)
	want := []int{1, 2, 3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("fib(%d) == %v, want %v", 6, got, want)
	}

}
