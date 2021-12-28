package main

import (
	"testing"
)

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

func TestHelloNameConst(t *testing.T) {
	got := HelloNameConst("0hlov3")
	want := "Hello, 0hlov3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloNameFallBack(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanr %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := HelloNameFallBack("0hlov3")
		want := "Hello, 0hlov3"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := HelloNameFallBack("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
