package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	MainHandler(w, r)

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
