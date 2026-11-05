package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func TestMyGreeterHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	MyGreeterHandler(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
	}

	got := response.Body.String()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
