package prototype

import (
	"sync"
)

type Proto struct {
	Tmp int64
	mtx sync.Mutex
	Str string
	Ch chan struct{}
}

func NewProto()*Proto{
	return &Proto{
		Ch: make(chan struct{}, 0),
	}
}

func (p *Proto)CloneProto()*Proto{
	t := *p
	return &t
}
