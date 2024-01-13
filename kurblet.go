// Connect to the local Kubernetes API as in
// https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration
package kurb

import "k8s.io/client-go/rest"

// TODO connect to the local Kubernetes API
func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	println(config)
}
