package main

import (
	"context"
	"encoding/json"
	"fmt"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	"log"
	proto "myGoMicro/micro-api/api/proto"
	"strings"
)

type Foo struct{}
type Example struct{}

func (f *Foo) Bar(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Println("收到一条消息")
	fmt.Println("%v\n", req)
	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.example", "no content")
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": "got your request" + strings.Join(name.Values, " "),
	})
	rsp.Body = string(b)
	return nil
}
func (e *Example) Call(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("log------revc call message")
	fmt.Printf("%+v\n", req)
	name := req.Name
	rsp.Msg = string("OK" + name)
	return nil

}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.example"))
	service.Init()
	proto.RegisterExampleHandler(service.Server(), new(Example))
	proto.RegisterFooHandler(service.Server(), new(Foo))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
