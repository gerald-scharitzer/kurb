// Connect to the local Kubernetes API as in
// https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration
package kurb

import "k8s.io/client-go/rest"

// TODO connect to the local Kubernetes API
// If `remote` is true, then connect to a remote cluster. Else connect to the local cluster.
func run(remote bool) {
	if remote {
		// TODO remote https://github.com/kubernetes/client-go/tree/master/examples/out-of-cluster-client-configuration
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		println(config)
	}
}
