package api

import (
	"context"
	"encoding/json"
	"fmt"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	proto "github.com/microservice/micro-api/api/noproto"
	"golang.org/x/tools/internal/telemetry/log"
	"log"
	"strings"
)

type Foo struct{}

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
