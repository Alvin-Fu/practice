package main

import (
    "github.com/Shopify/sarama"
    "fmt"
)

func main(){
    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true

    produer, err := sarama.NewSyncProducer([]string{"192.168.28.175:9092"}, config)
    if err != nil {
        panic(err)
    }
    defer produer.Close()
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
