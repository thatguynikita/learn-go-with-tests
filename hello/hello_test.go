package main

import "testing"

// TestHello tests Hello
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Spanish", func(t *testing.T) {
		got := Hello("Marco", "Spanish")
		want := "Hola, Marco!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in French", func(t *testing.T) {
		got := Hello("Marco", "French")
		want := "Bonjour, Marco!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say in Russian", func(t *testing.T) {
		got := Hello("Marco", "Russian")
		want := "Привет, Marco!"
		assertCorrectMessage(t, got, want)
	})
}
