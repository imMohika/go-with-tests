package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	assert(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
