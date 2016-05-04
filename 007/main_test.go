package main

import (
	"testing"
)

func Test_primes(t *testing.T) {
	var pos, want, got int

	pos = 6
	want = 13
	got = primeList(pos)[pos-1]

	if got != want {
		t.Errorf("primes(%d) == %d, want %d", pos, got, want)
	}
}
