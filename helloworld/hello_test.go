package main

import "testing"


func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		// %q	a double-quoted string safely escaped with Go syntax
		t.Errorf("got '%q' want '%q'", got, want)
	}
}