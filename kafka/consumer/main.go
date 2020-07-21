package main

import (
    "sync"
    "github.com/Shopify/sarama"
    "fmt"
)

var wg sync.WaitGroup
var topic = "test"
func main(){
    consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
    if err != nil {
        panic(err)
    }
    partitionList, err := consumer.Partitions(topic)
    if err != nil {
        panic(err)
    }
    for partition := range partitionList{
        pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
        if err != nil {
            panic(err)
        }
        defer pc.AsyncClose()
        wg.Add(1)
        go func(sarama.PartitionConsumer){
            defer wg.Done()
            for msg := range pc.Messages(){
                fmt.Printf("topic: %s, partition: %d, offset: %d, key: %s, value: %s\n",
                    topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
            }
        }(pc)
    }
    wg.Wait()
    consumer.Close()
}
