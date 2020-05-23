package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/v2"
	"io/ioutil"
	proto "myGoMicro/file/proto"
)

type File struct{}

func (g *File) File(ctx context.Context, file proto.File_FileStream) error {
	tmp, err := ioutil.TempFile("", "micro")
	if err != nil {
		return errors.InternalServerError("file.service", err.Error())
	}
	for {
		b, err := file.Recv()
		if err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
		if b.Len == -1 {
			break
		}
		if _, err := tmp.Write(b.Byte); err != nil {
			return errors.InternalServerError("file.service", err.Error())
		}
	}
	println(tmp.Name())
	return file.SendMsg(&proto.FileSliceMsg{FileName: tmp.Name()})
}

func main() {
	service := micro.NewService(
		micro.Name("file.service"),
		micro.Version("latest"),
	)
	service.Init()
	_ = proto.RegisterFileHandler(service.Server(), new(File))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
