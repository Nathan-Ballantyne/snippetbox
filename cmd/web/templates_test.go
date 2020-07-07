package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	// Initialize a new time.Time object and pass it to the humanDate function.
	tm := time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC)
	got := humanDate(tm)
	want := "17 Dec 2020 as 10:00"

	// Check that the output from the humanDate function is in the format we
	// expect. If it isn't what we expect, use the t.Errorf() function to
	// indicate that the test has failed and log the expected and actual
	// values.
	if got != want {
		t.Errorf("want %q; got %q", want, got)
	}
}
