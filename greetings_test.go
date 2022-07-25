package greetings

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Gara")
	expected := "Hi, Gara. Welcome!"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
