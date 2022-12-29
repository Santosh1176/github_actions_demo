package main

import "testing"

func TestGH_Actions(t *testing.T) {
	got := GH_Actions_Demo()
	want := "It's amazing learning Github Actions with #90DaysOfdevOps"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
