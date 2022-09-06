package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {

	t.Run("Saying Hello to people", func(t *testing.T) {

		got := hello("Chris", "")

		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)

	})

	t.Run("Saying hello to the world", func(t *testing.T) {

		got := hello("", "")

		want := "Hello, World!"

		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello freind in spanish", func(t *testing.T) {
		got := hello("Elodie", "spanish")
		want := "holla, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in french", func(t *testing.T) {

		got := hello("jimmy", "french")

		want := "Bonjour, jimmy"

		assertCorrectMessage(t, got, want)

	})

	t.Run("greet function", func(t *testing.T) {

		buffer := bytes.Buffer{}

		Greet(&buffer, "Jimmy")

		got := buffer.String()
		want := "Hello, Jimmy"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
