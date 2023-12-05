module github.com/t33m/nbs-csi-driver

go 1.21

replace github.com/ydb-platform/nbs => ./nbs

require (
	github.com/container-storage-interface/spec v1.9.0
	github.com/ydb-platform/nbs v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)

require (
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/oauth2 v0.11.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
)
