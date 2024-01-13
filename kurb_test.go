package kurb

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestNode(t *testing.T) {
	want := Node{CarbonEfficiency: 1}
	got := want
	assert.Equal(t, got, want)
}

// TODO parameterize want
func TestPod(t *testing.T) {
	var want *Node = nil
	pod := Pod{want}
	got := pod.Node
	assert.Equal(t, got, want)

	want = &Node{CarbonEfficiency: 1}
	pod = Pod{want}
	got = pod.Node
	assert.Equal(t, got, want)
}

func TestSchedule(t *testing.T) {
	want := MinimumCarbonEfficiency
	var nodes []Node
	got := Schedule(nodes)
	assert.Equal(t, got, want)

	want = 3
	nodes = []Node{
		{CarbonEfficiency: 1},
		{CarbonEfficiency: 3},
		{CarbonEfficiency: 2},
	}
	got = Schedule(nodes)
	assert.Equal(t, got, want)
}
