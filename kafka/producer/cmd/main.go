package main


import (
    "github.com/Shopify/sarama"
    "fmt"
)

func main(){
    config := sarama.NewConfig()
    // 等待服务器所有副本都保存成功后再响应
    config.Producer.RequiredAcks = sarama.WaitForAll
    // 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    // 是否等待成功和失败后的响应
    config.Producer.Return.Successes = true
    // 使用给定代理地址和配置创建一个同步生产者
    produer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
    if err != nil {
        panic(err)
    }
    defer produer.Close()
    // 构建发送的消息
    msg := &sarama.ProducerMessage{
        Partition: int32(10),
        Key: sarama.StringEncoder("key"),
    }
    var value, msgType string
    for {
        _, err := fmt.Scanf("%s", &value)
        if err != nil {
            break
        }
        fmt.Scanf("%s", &msgType)
        fmt.Printf("msg type: %s, value: %s\n", msgType, value)
        msg.Topic = msgType
        msg.Value = sarama.ByteEncoder(value)
        partition, offset, err := produer.SendMessage(msg)
        if err != nil {
            fmt.Printf("send message fail: %v\n", err)
        }
        fmt.Printf("partition: %d, offset: %d\n", partition, offset)
    }
}
