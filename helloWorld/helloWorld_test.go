package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Testing Hello to Ilya", func(t *testing.T) {
		got := Hello("Ilya")
		want := "Hello, Ilya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test without arguments", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}