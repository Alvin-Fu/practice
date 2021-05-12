STW（Stop the World）,是目前所有的拥有gc的语言都会有的现象，这个主要是为了保证在垃圾回收的时候，防止内存的继续增长从而导致GC不能正常的工作
在实际得使用过程中如果gc得STW时间太长会导致服务得响应延迟变大，甚至崩溃
```
func main() {
	go func() {
		for {

		}
	}()
	time.Sleep(time.Millisecond)
	runtime.GC()
	println("ok")
}
运行这段程序在现有得go版本中ok是不会被打印的，因为当GC在进行STW得时候需要将所有得goroutine停下来但是这时for循环退不出来得，导致STW被拖慢，造成卡死得情况。
这个问题在1.14里面被解决了，因为这一类得goroutine可以被异步抢占，从而使得STW得响应在0.5毫秒以内
```

