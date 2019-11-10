package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockMessageProvider struct {
	message string
}

func (m *mockMessageProvider) Message() string {
	return m.message
}


type mockVersionProvider struct {
	version string
}

func (m *mockVersionProvider) Version() string {
	return m.version
}

func TestWebAPI_HelloSpinSumHandler(t *testing.T) {
	cases := map[string]struct{
		provider MessageProvider
		expected string
	}{
		"happy path": {
			provider: &mockMessageProvider{message: "hello world"},
			expected: "hello world",
		},
	}

	for testName, c := range cases {
		t.Run(testName, func(t *testing.T) {
			api := WebAPI{MessageProvider: c.provider}
			req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(api.HelloSpinSumHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Fatal("did not get expected status code")
			}

			if rr.Body.String() != c.expected {
				t.Fatalf("did not get expected message")
			}
		})
	}
}

func TestWebAPI_VersionHandler(t *testing.T) {
	cases := map[string]struct{
		provider VersionProvider
		expected string
	}{
		"happy path": {
			provider: &mockVersionProvider{version: "0.0.1"},
			expected: "0.0.1",
		},
	}

	for testName, c := range cases {
		t.Run(testName, func(t *testing.T) {
			api := WebAPI{VersionProvider: c.provider}
			req, _ := http.NewRequest(http.MethodGet, "/version", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(api.VersionHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Fatal("did not get expected status code")
			}

			if rr.Body.String() != c.expected {
				t.Fatalf("did not get expected version")
			}
		})
	}
}
