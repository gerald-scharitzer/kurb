// internal Kubernetes types from https://github.com/kubernetes

package kurb

import (
	"sync"
)

// kubernetes/blob/master/pkg/scheduler/framework
// interface.go
type Code int

// interface.go
type Status struct {
	code    Code
	reasons []string
	err     error
	plugin  string
}

const (
	// interface.go
	Success Code = iota
)

var Status_Success Status = Status{code: Success}

// cycle_stage.go
type CycleState struct {
	storage             sync.Map
	recordPluginMetrics bool
	SkipFilterPlugins   Set[string] // FIXME sets.Set
	SkipScorePlugins    Set[string] // FIXME sets.Set
}

// apimachinery/blob/master/pkg/util/sets

// empty.go
type Empty struct{}

// set.go
type Set[T comparable] map[T]Empty
