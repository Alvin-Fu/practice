package util

import "sync"

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(fu func()) {
	w.Add(1)
	go func() {
		fu()
		w.Done()
	}()
}
