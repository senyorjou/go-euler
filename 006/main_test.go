package main

import (
	"reflect"
	"testing"
)

func Test_listN(t *testing.T) {
	var num int
	var want, got []int

	num = 10
	want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	got = listN(num)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("listN(%d) == %v, want %v", num, got, want)
	}
	if len(got) != len(want) {
		t.Errorf("len(listN(%d) == %d, want %d", num, len(got), len(want))
	}

	num = 5
	want = []int{1, 2, 3, 4, 5, 6}
	got = listN(num)

	if reflect.DeepEqual(want, got) {
		t.Errorf("listN(%d) == %v, want %v", num, got, want)
	}

	if len(got) == len(want) {
		t.Errorf("len(listN(%d) == %d, want %d", num, len(got), len(want))
	}
}

func Test_squareList(t *testing.T) {
	var list, want, got []int

	list = []int{1, 2, 3, 4}
	want = []int{1, 4, 9, 16}
	got = squareList(list)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("squareList(%v) == %v, want %v", list, got, want)
	}
}

func Test_sumList(t *testing.T) {
	var list []int
	var want, got int

	list = []int{1, 2, 3, 4}
	want = 10
	got = sumList(list)

	if want != got {
		t.Errorf("sumList(%v) == %v, want %v", list, got, want)
	}
}
