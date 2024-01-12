package kurb

import (
	"testing"
)

func TestNode(t *testing.T) {
	want := Node{CarbonEfficiency: 1}
	result := want
	if want != result {
		t.Fatal("failed")
	}
}

// TODO parameterize want
func TestPod0(t *testing.T) {
	var want *Node = nil
	pod := Pod{want}
	result := pod.Node
	if want != result {
		t.Fatal("failed")
	}
}

func TestPod1(t *testing.T) {
	want := &Node{CarbonEfficiency: 1}
	pod := Pod{want}
	result := pod.Node
	if want != result {
		t.Fatal("failed")
	}
}

func TestSchedule(t *testing.T) {
	want := MinimumCarbonEfficiency
	var nodes []Node
	result := Schedule(nodes)
	if want != result {
		t.Fatal("failed")
	}
}
