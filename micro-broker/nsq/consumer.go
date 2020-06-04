package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/broker/nsq"
	proto "myGoMicro/micro-broker/nsq/proto"
)

type Sub struct{}

func (s *Sub) Process(ctx context.Context, evt *proto.Msg) error {
	log.Logf("收到消息：Id %d & Timestamp %d with Data: %s\n", evt.Id, evt.Current, evt.Data)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.broker.nsq.srv"),
		micro.Broker(nsq.NewBroker(
			broker.Addrs([]string{"127.0.0.1:4150"}...),
		)),
	)
	srv.Init()
	sOpts := broker.NewSubscribeOptions(
		nsq.WithMaxInFlight(5),
	)
	_ = micro.RegisterSubscriber("go.micro.broker.topic.nsq", srv.Server(), &Sub{}, server.SubscriberContext(sOpts.Context))

	if err := srv.Run(); err != nil {
		log.Fatal("error occurs: %v", err)
	}
}
