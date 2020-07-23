package main

import (
    beegoConfig "github.com/astaxie/beego/config"
    "github.com/Shopify/sarama"
    "fmt"
    "github.com/judwhite/go-svc/svc"
    "path/filepath"
    "os"
    "syscall"
    "log"
    "practice/kafka/producer/option"
    "flag"
    "github.com/mreiferson/go-options"
    "practice/kafka/producer/server"
)
type config map[string] interface{}

type program struct {
    svr *server.KafkaSvr
}
func main(){
    prg := &program{}
    if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
        log.Fatal(err)
    }
}

func (p *program) Init(env svc.Environment)error{
    if env.IsWindowsService(){
        dir := filepath.Dir(os.Args[0])
        return os.Chdir(dir)
    }
    return nil
}

func (p *program) Start()error{
    opts := option.NewOptions()
    flagSet := initFlagSet(opts)
    flagSet.Parse(os.Args[1:])
    if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool){
       os.Exit(0)
    }
    var cnf config
    configFile := flagSet.Lookup("config").Value.String()
    if len(configFile) > 0 {
        parseConfigFile(configFile, (*map[string]interface{})(&cnf))
    }
    options.Resolve(opts, flagSet, cnf)
    svr := server.NewKafkaSvr(opts)
    svr.Main()
    p.svr = svr
    return nil
}

func (p *program) Stop()error{
    return nil
}

func initFlagSet(opts *option.Options)*flag.FlagSet{
    flagSet := flag.NewFlagSet("", flag.ExitOnError)
    flagSet.Bool("version", false, "print version string")
    flagSet.String("http-addr", opts.HTTPAddr, "<addr>:<port> to listen on for HTTP clients")
    flagSet.String("kafka-addr", opts.KafkaAddr, "<addr>:<port> to connect kafka service")
    return flagSet
}

func parseConfigFile(configFile string, result *map[string]interface{})error{
    config, err := beegoConfig.NewConfig("ini", configFile)
    if err != nil {
        return err
    }
    cfg := make(map[string]interface{})
    httpAddr := config.DefaultString("http.addr", "")
    if len(httpAddr) > 0 {
        cfg["http-addr"] = httpAddr
    }

    kafkaAddr := config.DefaultString("kafka.addr", "")
    if len(kafkaAddr) > 0 {
        cfg["kafka-addr"] = kafkaAddr
    }
    *result = cfg
    return nil
}

func kafka(){
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