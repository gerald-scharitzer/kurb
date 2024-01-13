package kurb

// Carbon efficiency is the amount of energy that can be obtained per mass of carbon dioxide equivalent in joule per kilogram.
type CarbonEfficiency int64

const MinimumCarbonEfficiency CarbonEfficiency = -0x8000_0000_0000_0000

// Nodes run Pods.
// Nodes have resources to process workloads.
type Node struct {
	CarbonEfficiency CarbonEfficiency
}

// Pods are units of work (an amount of workload) that can be scheduled to run on nodes.
type Pod struct {
	// The pod runs on this node. `nil` is no node and then the pod does not run.
	Node *Node
}

// Get the lowest carbon efficiency of all nodes.
// Return MinimumCarbonEfficiency for an empty set of nodes
func Schedule(nodes []Node) CarbonEfficiency {
	var ce CarbonEfficiency = MinimumCarbonEfficiency
	// max over node.CarbonEfficiency
	for _, node := range nodes {
		nce := node.CarbonEfficiency
		if nce > ce {
			ce = nce
		}
	}
	return ce
}
