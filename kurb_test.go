package kurb

import (
	"testing"
)

func TestKurb(t *testing.T) {
	want := "kurb"
	result := Name()
	if want != result {
		t.Fatal("failed")
	}
}
