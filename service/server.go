package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	proto "myGoMicro/service/proto"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "service proto 你好，" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter.service"),
		micro.Version("latest"),
	)
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
