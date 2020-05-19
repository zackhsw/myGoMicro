package main

import (
	"context"
	"io/ioutil"
	proto "myGoMicro/file/proto"
)

type File struct{}

func (g *File) File(ctx context.Context, file proto.File_FileStream) error{
	tmp, err := ioutil.TempFile("","micro")
	if err :=nil{

	}
}
