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

## Pod Lifecycle

When a pod is scheduled to run on a node, then the pod [remains on that node](https://kubernetes.io/docs/concepts/workloads/pods/#working-with-pods)
until the pod finishes execution, is deleted or evicted, or the node fails.

The pod is scheduled as part of its [lifecycle](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase) phase `Pending`
and when the pod has been scheduled to a node,
then the [pod condition contains](https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-conditions) `PodScheduled`.

Pods are usually created indirectly via workload controllers.
Change the pod template of the controller to trigger new pods on carbon-efficient nodes.

## Control Plane

The communication from the control plane to nodes [must be secured](https://kubernetes.io/docs/concepts/architecture/control-plane-node-communication/#control-plane-to-node) over untrusted or public networks.

# Rationale

Carbon-efficient energy sources like sun and wind are more variable than other sources.
This affects the availability of carbon-efficient energy to substantial consumers like data centers.
The amount of available energy and its carbon-efficiency both vary with time and space.
Computing workload is relatively free in both time and space.
Following the availability of carbon-efficient energy along the globe improves the carbon efficieny of cloud computing.
The carbon efficiency of energy sources varies with geographic distances.
Within one data center (zone) and across geographically close data centers (within one region) the carbon efficiency is probably the same.
Geographically large distances between data centers result in different carbon efficiencies per location (zone).

Sustainable Computing -- Without the Hot Air
https://arxiv.org/abs/2207.00081

Electricity Maps
https://www.electricitymaps.com/

## Kubernetes

schedules workloads on nodes, which map to physical computers and locations.
This is like an operating system for networked computers.

Covering large geographic areas with one Kubernetes cluster is probably not supported by cloud providers.

Scheduling workload with short response times can be done with workload controllers for continuously available services like `Deployment`, `StatfulSet`, and `DaemonSet`.
Together with the `HorizontalPodAutoscaler` they can scale their capacity with the incoming messages.
Global network traffic managers can route messages to more carbon-efficient locations.
Stateless services just reply to the messages that they receive.
Pure stateless services are pure functions, which return the same replay to the same message.
Stateful services must overcome data gravity.

Batch `Job` and `CronJob` can be scheduled for workload with long response times.

## Azure

For Azure Kubernetes Service (AKS) in [multiple regions](https://learn.microsoft.com/en-us/azure/architecture/reference-architectures/containers/aks/baseline-aks#multiple-regions)
the provider recommends to run [multiple clusters in different regions](https://learn.microsoft.com/en-us/azure/architecture/reference-architectures/containers/aks-multi-region/aks-multi-cluster).

AKS clusters can span multiple zones.

AKS clusters are deployed to [single regions](https://learn.microsoft.com/en-us/azure/aks/operator-best-practices-multi-region#plan-for-multiregion-deployment).

# Develop

1. Get with `git clone https://github.com/gerald-scharitzer/kurb.git`
2. Test with `go test`
