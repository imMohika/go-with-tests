package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello with name", func(t *testing.T) {
		got := Hello("Mohika", "")
		want := "Hello, Mohika"
		assert(t, got, want)
	})

	t.Run("hello with name in Spanish", func(t *testing.T) {
		got := Hello("Mohika", "Spanish")
		want := "Hola, Mohika"
		assert(t, got, want)
	})

	t.Run("hello with name in French", func(t *testing.T) {
		got := Hello("Mohika", "French")
		want := "Bonjour, Mohika"
		assert(t, got, want)
	})

	t.Run("hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
