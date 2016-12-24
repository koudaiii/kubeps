# kubeps

Get container image tag for Kubernetes Pods

As you know, `kubectl get pod -o wide` can get only pod( NAME,READY,STATUS, RESTARTS,AGE,IP,NODE).
`kubectl get po` difficult for you to get container image or tag.
`kubeps` enables you to get container image and tag in ALL pods that the specified namespace or labels.

