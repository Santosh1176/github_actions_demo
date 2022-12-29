package main

import "testing"

func TestGH_Actions(t *testing.T) {
	got := GH_Actions_Demo()
	want := "Welcome to Github Actions"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
