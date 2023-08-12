package context

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("stub store got cancelled")
				return
			default:
				time.Sleep(20 * time.Millisecond)
				result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-data:
		return result, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		resBody := response.Body.String()
		if resBody != data {
			t.Errorf(`got "%s", want "%s"`, resBody, data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &StubStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		server.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
