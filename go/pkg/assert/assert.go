package assert

import "slices"

import "testing"

func Nil(t *testing.T, actual any) {
	t.Helper()
	if actual != nil {
		t.Fatalf("got: %v, expected: <nil>", actual)
	}
}

func Equal(t *testing.T, expected, actual any) {
	t.Helper()
	if expected != actual {
		t.Fatalf("got: %v, want: %v", actual, expected)
	}
}

func EqualSlices[T comparable](t *testing.T, expected, actual []T) {
	t.Helper()
	if len(expected) != len(actual) {
		t.Fatalf("got len: %v, want: %v", len(actual), len(expected))
	}

	for i := range expected {
		if expected[i] != actual[i] {
			t.Fatalf("got %v at %v, want: %v", actual[i], i, expected[i])
		}
	}
}

func Contains[T comparable](t *testing.T, expected T, slice []T) {
	t.Helper()
	if !slices.Contains(slice, expected) {
		t.Fatalf("expected slice to contain: %v", expected)
	}
}
