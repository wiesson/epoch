package main

import (
	"testing"
)

func TestWithCorrectTimestamp(t *testing.T) {
	ts := Timestamp{}
	ts.FromString("1497052800")

	expected := "2017-06-10 00:00:00 +0000 UTC"
	got := ts.UtcString()

	if got != expected {
		t.Error("expected", expected, "got", got)
	}
}
