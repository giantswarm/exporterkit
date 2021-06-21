module github.com/giantswarm/exporterkit

go 1.13

require (
	github.com/giantswarm/microendpoint v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/microkit v0.2.2
	github.com/giantswarm/micrologger v0.5.0
	github.com/prometheus/client_golang v1.9.0
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
)

replace (
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/gogo/protobuf v1.2.1 => github.com/gogo/protobuf v1.3.2
)
