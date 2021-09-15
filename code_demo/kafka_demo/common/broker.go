package common

import (
	"github.com/Shopify/sarama"
	"log"
)
const (
	VERSION = "2.6.0"
	ADDR = "localhost:9002"
	GROUP = "dream"
)

func NewAsyncProducer()  sarama.AsyncProducer {
	cfg := sarama.NewConfig()
	version,err := sarama.ParseKafkaVersion(VERSION)
	if err != nil {
		log.Fatalf("NewAsyncProducer Parse kafka version failed:%s",err.Error())
	}
	cfg.Version = version
	//三种模式：
	//①NoResponse：0,数据发送成功与否都不关心
	//②WaitForLocal：1,当leader接收成功,返回
	//③WaitForAll：-1,所有的leader和follower的都接收成功
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

}

