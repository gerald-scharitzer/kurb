package kurb

import (
	"testing"
)

func TestKurb(t *testing.T) {
	want := "kurb"
	result := Run()
	if want != result {
		t.Fatal("failed")
	}
}
