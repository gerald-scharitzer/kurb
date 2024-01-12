package kurb

import (
	"testing"
)

func TestPlugin(t *testing.T) {
	want := "kurb"
	result := Name()
	if want != result {
		t.Fatal("failed")
	}
}

func TestPreScore(t *testing.T) {
	var want *Status = &Status_Success
	result := PreScore(nil, nil)
	if want != result {
		t.Fatal("failed")
	}
}
