package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	RootHandler(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("res.StatusCode = %d, want = %d", w.Code, http.StatusOK)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadAll() err= %v; want %v", err, nil)
	}
	got := string(body)

	want := "Main handler"
	if got != want {
		t.Fatalf("body contents = %v; want %v", got, want)
	}
}

func TestVersionHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/version", nil)

	VersionHandler(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("res.StatusCode = %d, want = %d", w.Code, http.StatusOK)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadAll() err= %v; want %v", err, nil)
	}
	got := string(body)

	want := fmt.Sprintf("Release: %s\nCommit: %s\nRepo: %s\nDate: %s", RELEASE, COMMIT, REPO, DATE)
	if got != want {
		t.Fatalf("body contents = %v; want %v", got, want)
	}
}

func TestHealthHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/version", nil)

	HealthHandler(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("res.StatusCode = %d, want = %d", w.Code, http.StatusOK)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("ReadAll() err= %v; want %v", err, nil)
	}
	got := string(body)

	want := r.URL.String()
	if got != want {
		t.Fatalf("body contents = %v; want %v", got, want)
	}
}
