package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	wanted := "Hello! World"
	if got != wanted {
		t.Errorf("got %q but wanted %q", got, wanted)
	}
}
