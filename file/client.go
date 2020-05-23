package main

import (
	"context"
	"github.com/micro/go-micro/v2"
	"io"

	"github.com/micro/go-micro/v2/client"
	proto "myGoMicro/file/proto"
	"net/http"
)

var c client.Client
var fileService proto.FileService

func UploadFile(rsp http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	// 获取文件对象
	files, ok := req.MultipartForm.File["file"]
	if !ok {
		rsp.WriteHeader(400)
		_, _ = rsp.Write([]byte("请选择文件上传"))
		return
	}
	// 将文件通过流式传输到srv
	file, err := files[0].Open()
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	// 建立链接
	// 采用固定节点获取文件
	next, _ := c.Options().Selector.Select("file.service")
	node, _ := next()
	stream, err := fileService.File(context.Background(), func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})
	if err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	for {
		buff := make([]byte, 1024*1024)
		sendLen, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				err = stream.Send(&proto.FileSlice{
					Byte: nil,
					Len:  -1,
				})
				if err != nil {
					rsp.WriteHeader(500)
					_, _ = rsp.Write([]byte(err.Error()))
					return
				}
				break
			}
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
		err = stream.Send(&proto.FileSlice{
			Byte: buff[:sendLen],
			Len:  int64(sendLen),
		})
		if err != nil {
			rsp.WriteHeader(500)
			_, _ = rsp.Write([]byte(err.Error()))
			return
		}
	}
	//等待接收，然后关闭
	fileMsg := &proto.FileSliceMsg{}
	if err := stream.RecvMsg(fileMsg); err != nil {
		rsp.WriteHeader(500)
		_, _ = rsp.Write([]byte(err.Error()))
		return
	}
	_ = stream.Close()
	println(fileMsg.FileName)
}

func main() {
	service := micro.NewService(micro.Name("file.client"))
	service.Init()
	c = service.Client()
	fileService = proto.NewFileService("file.service", c)
	http.HandleFunc("/upload", UploadFile)
	_ = http.ListenAndServe(":8085", nil)

}
