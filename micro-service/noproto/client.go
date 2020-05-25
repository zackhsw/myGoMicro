package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func main() {
	service := micro.NewService()
	service.Init()
	c := service.Client()
	//客户端请求：
	//通过指定service的名称进行服务发现
	//通过服务发现选取到服务中的某个Node
	//给Node发送Grpc请求
	//得到响应
	//     request:=c.NewRequest("service.greeter","Greeter.Hello","AllenHuang",client.WithContentType("application/json"))
	request := c.NewRequest("service.greeter", "Greeter.Hello",
		"AllenHuang", client.WithContentType("application/json"))
	var response string
	if err := c.Call(context.TODO(), request, &response); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)

}
