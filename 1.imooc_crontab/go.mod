module 1.imooc_crontab

go 1.12

replace (
	github.com/apache/rocketmq-client-go => ../../vendor/rocketmq-client-go
	zuji/common => ../../common
)

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.12.1 // indirect
	github.com/json-iterator/go v1.1.8 // indirect
	github.com/mongodb/mongo-go-driver v1.1.3
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/wcybxzj/go_study2 v0.0.0-20191206164705-1e1fdd774ff5
	go.etcd.io/etcd v3.3.18+incompatible
	go.mongodb.org/mongo-driver v1.1.3
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.25.1 // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)
