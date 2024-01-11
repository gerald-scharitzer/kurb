// internal Kubernetes types from https://github.com/kubernetes/kubernetes/blob/master/pkg/scheduler/framework/interface.go

package kurb

type Code int

type Status struct {
	code    Code
	reasons []string
	err     error
	plugin  string
}

const (
	Success Code = iota
)

var Status_Success Status = Status{code: Success}
