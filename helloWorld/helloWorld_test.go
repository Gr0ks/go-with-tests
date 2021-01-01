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
		got := Hello("Ilya", "en")
		want := "Hello, Ilya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test without arguments", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func (t *testing.T) {
		got := Hello("Ilya", "sp")
		want := "Hola, Ilya"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Russian", func (t *testing.T) {
		got := Hello("Илья", "ru")
		want := "Привет, Илья"
		assertCorrectMessage(t, got, want)
	})
}