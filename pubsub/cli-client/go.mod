module cli-client

go 1.16

require (
	github.com/asim/go-micro/plugins/registry/kubernetes/v4 v4.0.0-20220118152736-9e0be6c85d75
	github.com/pborman/uuid v1.2.1
	go-micro.dev/v4 v4.5.0
	pubsub-srv v0.0.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace pubsub-srv v0.0.0 => ../../pubsub/srv
