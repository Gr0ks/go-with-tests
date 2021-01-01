package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Testing Hello to Ilya", func(t *testing.T) {
		got := Hello("Ilya")
		want := "Hello, Ilya"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test without arguments", func(t *testing.T) {
		got := Hello()
		want := "Hello, world"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}