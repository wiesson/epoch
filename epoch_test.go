package main

import (
	"testing"
)

func TestWithCorrectTimestamp(t *testing.T) {
	expected := "2017-06-10 00:00:00 +0000 UTC"

	ts := Timestamp{}
	ts.FromString("1497052800")

	if ts.UtcString() != expected {
		t.Error("expected", expected, "got", ts.UtcString())
	}
}
