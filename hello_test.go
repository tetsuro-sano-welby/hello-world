package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello people", func(t *testing.T) {
        got := Hello("Chris")
        want := "Hello, Chris"

        assertCorrectNessage(t, got, want)
    })
    t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := Hello("")
        want := "Hello, World"

        assertCorrectNessage(t, got, want)
    })
}

func assertCorrectNessage(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
