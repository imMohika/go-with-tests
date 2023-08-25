package blog

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

var ErrStubFailingFS = errors.New("oh no, i always fail")

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, ErrStubFailingFS
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("posts length", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte("hi")},
			"hello-world2.md": {Data: []byte("hola")},
		}

		posts, err := NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}
	})

	t.Run("with StubFailingFS", func(t *testing.T) {
		fs := StubFailingFS{}

		_, err := NewPostsFromFS(fs)

		assertError(t, err, ErrStubFailingFS)
	})

	t.Run("post data", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
			secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := NewPostsFromFS(fs)

		got := posts[0]
		want := Post{
			Title:       "Post 1",
			Description: "Description 1",
			Tags:        []string{"tdd", "go"},
			Body: `Hello
World`,
		}

		assertPost(t, got, want)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("wanted an error but didn't got any")
	}

	if !errors.Is(want, got) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
