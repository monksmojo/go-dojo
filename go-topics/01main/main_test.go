package main

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	var name string = "Foo"
	run1Name := fmt.Sprintf("Greetings to %v", name)
	run2Name := fmt.Sprint("Greetings to World !")
	t.Run(run1Name, func(t *testing.T) {
		actual := Greet(name)
		expected := fmt.Sprintf("Greetings ! %s \n Welcome To Go Topics, By Monks Mojo", name)
		assertTestResponse(t, actual, expected)
	})

	t.Run(run2Name, func(t *testing.T) {
		actual := Greet("")
		expected := fmt.Sprint("Greetings ! World \n Welcome To Go Topics, By Monks Mojo")
		assertTestResponse(t, actual, expected)
	})

}

func assertTestResponse(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("got: %q want: %q \n", actual, expected)
	}
}
