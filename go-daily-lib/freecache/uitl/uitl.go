package uitl

import "sync"

type WaitGroupWrapper struct{
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cd func()){
	w.Add(1)
	go func() {
		cd()
		w.Done()
	}()
}
