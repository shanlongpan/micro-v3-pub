module github.com/shanlongpan/micro-v3-pub

go 1.15

require (
	github.com/asim/go-micro/plugins/client/grpc/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/asim/go-micro/plugins/registry/etcd/v3 v3.7.0
	github.com/asim/go-micro/plugins/selector/shard/v3 v3.7.0
	github.com/asim/go-micro/plugins/server/grpc/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v3 v3.7.0
	github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v3 v3.7.0 // indirect
	github.com/asim/go-micro/plugins/wrapper/select/shard/v3 v3.7.0 // indirect
	github.com/asim/go-micro/v3 v3.7.0
	github.com/google/uuid v1.2.0
	google.golang.org/genproto v0.0.0-20210821163610-241b8fcbd6c8 // indirect
	google.golang.org/grpc/examples v0.0.0-20211015201449-4757d0249e2d // indirect
	google.golang.org/protobuf v1.27.1
)
