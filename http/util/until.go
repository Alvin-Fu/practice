package util

import "sync"

type WaitGroupWrapper struct{
	wg sync.WaitGroup
}

func (w *WaitGroupWrapper)Wrap(fu func()){
	w.wg.Add(1)
	go func() {
		fu()
		w.wg.Done()
	}()
}