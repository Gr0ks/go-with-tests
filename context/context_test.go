package context

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

type StubStore struct {
	response string
}

func (s StubStore) Fetch() string {
	return s.response 
}

func TestServer(t *testing.T) {
	data := "hello, world"
	srv := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	srv.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}