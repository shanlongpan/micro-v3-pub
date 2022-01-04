/**
* @Author:Tristan
* @Date: 2021/12/22 10:45 上午
 */

package MicroV3Adapter

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/plugins/selector/shard/v3"
	"github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/google/uuid"
	"github.com/shanlongpan/micro-v3-pub/idl/grpc/microv3"
	"time"
)

var clientInstance microv3.MicroV3Service

var (
	EtcdEndPoints             = []string{""}
	HashKey                   = "hash_key"
	TraceId                   = "trace_id"
	TimeOut                   = 3000 //3s 超时
	MaxConcurrentRequestNum   = 1000 //最大并发数1000
	ErrorPercentThresholdPer  = 30   // 30%请求报错，熔断触发
	RequestVolumeThresholdNum = 10   // 窗口时间内，超过次数量开始健康监控
)

func init() {
	// etcd 服务注册和发现以后改成环境变量配置
	reg := etcd.NewRegistry(registry.Addrs(EtcdEndPoints...))
	service := micro.NewService(
		micro.Name("micro-v3-learn"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.Client(grpc.NewClient(
			client.Registry(reg),
		)),
	)
	hystrix.ConfigureDefault(hystrix.CommandConfig{
		Timeout:                TimeOut,
		MaxConcurrentRequests:  MaxConcurrentRequestNum,
		RequestVolumeThreshold: RequestVolumeThresholdNum,
		ErrorPercentThreshold:  ErrorPercentThresholdPer,
	})

	// 初始化服务
	service.Init(micro.WrapClient(hystrix.NewClientWrapper()))

	clientInstance = microv3.NewMicroV3Service("micro-v3-learn", service.Client())
}

func FromContext(ctx context.Context, key string) (string, bool) {
	u, ok := ctx.Value(key).(string)
	return u, ok
}

func setMetaTraceId(ctx context.Context) context.Context {
	traceId, ok := FromContext(ctx, TraceId)
	if !ok {
		traceId = uuid.New().String()
	}
	ctx = metadata.Set(ctx, TraceId, traceId)
	return ctx
}

func getHashKey(ctx context.Context) string {
	hashKey, ok := FromContext(ctx, HashKey)
	if !ok || len(hashKey) == 0 {
		hashKey = fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hashKey
}

func Call(ctx context.Context, req *microv3.CallRequest, opts ...client.CallOption) (*microv3.CallResponse, error) {
	ctx = setMetaTraceId(ctx)
	opts = append(opts, shard.Strategy(getHashKey(ctx)))
	rsp, err := clientInstance.Call(ctx, req, opts...)
	if err != nil {
		// 打日志
		fmt.Println("Error calling MicroV3: ", err)
	}

	return rsp, err
}

func Stream(ctx context.Context, in *microv3.StreamingRequest, opts client.CallOption) (microv3.MicroV3Service_StreamService, error) {
	//
	//rsp, err := clientInstance.Stream(ctx, in,opts)
	//if err != nil {
	//	// 打日志
	//	fmt.Println("Error calling MicroV3: ", err)
	//}
	//
	//return rsp,err
	return nil, nil
}

func PingPong() {

}
