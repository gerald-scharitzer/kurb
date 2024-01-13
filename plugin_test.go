package kurb

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestPlugin(t *testing.T) {
	want := "kurb"
	got := Name()
	assert.Equal(t, got, want)
}

func TestPreScore(t *testing.T) {
	var want *Status = &Status_Success
	got := PreScore(nil, nil)
	assert.Equal(t, got, want)
}
