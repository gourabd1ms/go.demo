package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello from Go demo!"
	if got := Greet(); got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
