package assert

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
