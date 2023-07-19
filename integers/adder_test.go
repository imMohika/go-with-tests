package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Add(2,2)
	want := 4
	assert(t, got, want)
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
