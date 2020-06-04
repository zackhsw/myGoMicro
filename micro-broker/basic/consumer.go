package main

import (
	"fmt"
	"github.com/micro/go-micro/broker"
	"log"
)

var (
	topic1 = "go.micro.topic.foo"
)

func sharedSub() {
	_, err := broker.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("[sub] 获取到消息:", string(p.Message().Body), " Header:", p.Message().Header)
		return nil
	}, broker.Queue("c1"))
	if err != nil {
		fmt.Println(err)
	}
}

func sub() {
	_, err := broker.Subscribe(topic1, func(p broker.Event) error {
		fmt.Println("[sub] 获取消息：”", string(p.Message().Body), " Header:", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error:%v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	//sub()
	sharedSub()
	select {}
}
