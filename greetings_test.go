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
		result, _ := Hello("Gara")
		expected := "Hi, Gara. Welcome!"

		expectEqual(t, result, expected)
	})

	t.Run("should return error if <name> is empty", func(t *testing.T) {
		_, error := Hello("")
		expected_error := ErrorEmptyName.Error()

		expectEqual(t, error.Error(), expected_error)
	})
}
