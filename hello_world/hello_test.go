package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying Hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		wanted := "Hello! Chris"
		assertCorrectMessage(t, got, wanted)
	})
	t.Run("say, hello world when empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello! World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("spanish translation", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola! Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("french translation", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour! Chris"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
