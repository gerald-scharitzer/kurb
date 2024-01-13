# Kurb

optimizes the carbon efficiency of Kubernetes workloads.

# Goal

Run workload where and when carbon-efficient energy is available for computing.

# Method

Configure the Kubernetes scheduler to create pods on the eligible nodes with the highest carbon efficiency and subsequently the lowest greenhouse gas emissions.

Node score is a 64-bit signed binary integer, so that can store high values of carbon efficiency.

# Setup

1. Create a Kubernetes cluster (https://kubernetes.io/docs/setup/)
2. Label the work nodes with their carbon efficiency in Joule per kg CO2-equivalent (https://kubernetes.io/docs/tasks/configure-pod-container/assign-pods-nodes/#add-a-label-to-a-node)
3. Configure a custom scheduler plugin that picks nodes with the highest carbon efficiency first (https://github.com/kubernetes-sigs/scheduler-plugins).

# Limits

Configuring the scheduler such that nodes with the highest carbon efficiency are picked first (https://kubernetes.io/docs/reference/scheduling/config/) by configuring the plugin NodeResourcesFit (https://kubernetes.io/docs/concepts/scheduling-eviction/resource-bin-packing/) does not work, because the carbon efficiency is a node level attribute.
None of the default plugins extending the `score` phase supports that.

## Scheduler Plugins

must be [compiled together](https://github.com/kubernetes/kubernetes#to-start-using-k8s) with the `kube-scheduler`,
because while the versioned API is stable,
the [internal structures](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md#so-you-want-to-change-the-api) (Go programming interfaces),
like the extension points for scheduler plugins,
are not stable.

The scheduler framework is Go package [`k8s.io/kubernetes/pkg/scheduler/framework`](https://pkg.go.dev/k8s.io/kubernetes/pkg/scheduler/framework), but using the packages `k8s.io/kubernetes/...` as libraries is [not supported](https://github.com/kubernetes/kubernetes/tree/master#to-start-using-k8s).

Copy internal Kubernetes types to verify the concept.

# Rationale

Carbon-efficient energy sources like sun and wind are more variable than other sources.
This affects the availability of carbon-efficient energy to substantial consumers like data centers.
Computing workload is relatively free in both time and space.
Following the availability of carbon-efficient energy along the globe improves the carbon efficieny of cloud computing.

Sustainable Computing -- Without the Hot Air
https://arxiv.org/abs/2207.00081

# Develop

1. Get with `git clone https://github.com/gerald-scharitzer/kurb.git`
2. Test with `go test`
