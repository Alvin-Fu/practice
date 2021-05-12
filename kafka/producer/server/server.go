package server

import (
    "sync"
    "sync/atomic"
    "practice/kafka/util"
    "practice/kafka/producer/option"
)

type KafkaSvr struct {
    sync.Mutex
    opts atomic.Value
    exitChan chan struct{}
    waitGroup util.WaitGroupWrapper
}

func NewKafkaSvr(opts *option.Options)*KafkaSvr{
    svr := &KafkaSvr{
        exitChan: make(chan struct{}),
    }
    svr.swapOpts(opts)
    return svr
}

func (s *KafkaSvr) Main()error{
    var err error

    return nil
}

func (s *KafkaSvr) Exit(){}

func (s *KafkaSvr) swapOpts(opts *option.Options){
    s.opts.Store(opts)
}

func (s *KafkaSvr) getOpts()*option.Options{
    return s.opts.Load().(*option.Options)
}