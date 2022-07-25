package greetings

import (
	"testing"
)

func expectEqual(t testing.TB, result, expected interface{}) {
	t.Helper()
	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}

func TestHello(t *testing.T) {
	t.Run("should create greeting to <name>", func(t *testing.T) {
		result, err := Hello("Gara")

		expectEqual(t, result, "Hi, Gara. Welcome!")
		expectEqual(t, err, nil)
	})

	t.Run("should return error if <name> is empty", func(t *testing.T) {
		result, err := Hello("")

		expectEqual(t, err.Error(), ErrorEmptyName.Error())
		expectEqual(t, result, "")
	})
}
