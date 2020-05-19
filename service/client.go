package main

import (
	"context"
	"fmt"
	//"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2"
	proto "myGoMicro/service/proto" //proto文件所在的位置
)

func main() {
	service := micro.NewService(micro.Name("greeter.client"))
	service.Init()
	greeter := proto.NewGreeterService("greeter.service", service.Client())
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Hello go"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Greeting)
}
