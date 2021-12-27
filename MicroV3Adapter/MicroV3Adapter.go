/**
* @Author:Tristan
* @Date: 2021/12/22 10:45 上午
 */

package MicroV3Adapter

import (
	"context"
	"fmt"
	_ "github.com/asim/go-micro/plugins/client/grpc/v3"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	_ "github.com/asim/go-micro/plugins/server/grpc/v3"
	"github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v3"
	"github.com/asim/go-micro/plugins/wrapper/select/shard/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/shanlongpan/micro-v3-pub/idl/grpc/microv3"
	"math/rand"
	"strconv"
	"time"
)

var clientInstance microv3.MicroV3Service

const HashKey = "hash_key"

func init() {
	// etcd 服务注册和发现以后改成环境变量配置
	reg := etcd.NewRegistry(registry.Addrs("http://127.0.0.1:2377", "http://127.0.0.1:2378", "http://127.0.0.1:2379"))
	service := micro.NewService(
		micro.Name("micro-v3-learn"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.WrapClient(hystrix.NewClientWrapper(), shard.NewClientWrapper(HashKey)),
		//micro.Client(grpc.NewClient()),
	)

	// 初始化服务
	service.Init()

	// create the proto client for MicroV3（创建 grpc client）
	clientInstance = microv3.NewMicroV3Service("micro-v3-learn", service.Client())

}

// 没有 hashKey 设置默认的
func setHashKey(ctx context.Context) context.Context {
	value, ok := metadata.Get(ctx, HashKey)
	if !ok {
		rand.Seed(time.Now().Unix())
		value = strconv.Itoa(rand.Int())
		ctx = metadata.Set(ctx, HashKey, value)
	}
	return ctx
}

func Call(ctx context.Context, req *microv3.CallRequest, opts ...client.CallOption) (*microv3.CallResponse, error) {
	ctx = setHashKey(ctx)
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
