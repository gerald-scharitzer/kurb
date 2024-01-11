package kurb

import (
	"context"
)

// implement k8s.io/kubernetes/pkg/scheduler/framework.Plugin
func Name() string {
	return "kurb"
}

// TODO implement k8s.io/kubernetes/pkg/scheduler/framework.PreScorePlugin
// to read the carbon efficiency from the node labels
func PreScore(ctx context.Context) *Status {
	return &Status_Success
}

// TODO implement k8s.io/kubernetes/pkg/scheduler/framework.ScorePlugin
// to score nodes by their carbon efficiency
