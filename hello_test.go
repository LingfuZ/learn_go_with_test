package main

import "testing"

func TestHello(t *testing.T) {
	// input := "Chris"

	// got := Hello(input)
	// want := "Hello, " + input

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// t.Run("saying hello to people", func(t *testing.T) {
	// 	got := Hello("Chris")
	// 	want := "Hello, Chris"
	// 	assertCorrectMessage(t, got, want)
	// 	// if got != want {
	// 	// 	t.Errorf("got %q want %q", got, want)
	// 	// }
	// })

	// t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
	// 	got := Hello("")
	// 	want := "Hello, World"

	// 	assertCorrectMessage(t, got, want)
	// 	// if got != want {
	// 	// 	t.Errorf("got %q want %q", got, want)
	// 	// }
	// })

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
