package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloName(t *testing.T) {
	got := HelloName("0hlov3")
	want := "Hello, 0hlov3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
