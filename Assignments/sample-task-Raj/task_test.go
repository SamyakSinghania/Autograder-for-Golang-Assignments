package sampleTask

import (
	"testing"
)

func TestTask(t *testing.T) {
	got := Task()
	want := Completed

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
