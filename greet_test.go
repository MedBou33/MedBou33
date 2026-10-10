package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Go")
	want := "Hello, Go"

	if got != want {
		t.Errorf("Greet(\"Go\") = %q, want %q", got, want)
	}
}
