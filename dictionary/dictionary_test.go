package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assert(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unk")
		want := ErrUnknownKey

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("new word", func(t *testing.T) {
		err := dictionary.Add("test", "this is a test")
		assertNoError(t, err)

		got, _ := dictionary.Search("test")
		want := "this is a test"

		assert(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		got := dictionary.Add("test", "this is a test")
		want := ErrExistingKey

		assertError(t, got, want)
	})
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}

	t.Run("known word", func(t *testing.T) {
		err := dictionary.Update("test", "updated")
		assertNoError(t, err)

		got, _ := dictionary.Search("test")
		want := "updated"

		assert(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got := dictionary.Update("unknown", "this is a test")
		want := ErrUnknownKey

		assertError(t, got, want)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{
		"test": "this is a test",
	}

	dictionary.Delete("test")

	_, got := dictionary.Search("test")
	want := ErrUnknownKey

	assertError(t, got, want)
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("wanted an error but didn't got any")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got %v, didn't wanted any", got)
	}
}
