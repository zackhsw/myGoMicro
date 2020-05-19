package main

import (
	"context"
	"github.com/micro/go-micro/v2"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	print("a new request cones")
	*msg = "你好，" + *name + "---其实这是一条response的回复字串"
	return nil
}

func (g *Greeter) Hello2(ctx context.Context, arg *string, msg *string) error {
	print("这是hello2的请求")
	*msg = "收到您的消息"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("service.greeter"),
	)
	service.Init()
	micro.RegisterHandler(service.Server(), new(Greeter))
	service.Run()
}
