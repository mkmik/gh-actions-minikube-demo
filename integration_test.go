// +build integration

package main

import (
	"net/http"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	addr := os.ExpandEnv("${SERVER_ADDR}")
	res, err := http.Get(addr)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := res.StatusCode, http.StatusOK; got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
