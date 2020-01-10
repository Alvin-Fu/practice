golang中的bufio包是用来处理I/O缓存的，主要有提供writer，reader和scanner等功能
大部分的硬件设备对处理合适的块对齐的数据是更友好的，合适的块数据可以提高程序的性能

## Writer功能
	将多次且少量的I/O操作，最终合并成一次I/O操，减少系统调用的次数
bufio.Writer是go中提供的带缓存的，不会直接进行I/O操作而是将数据先写入缓存，当缓存满的时候在统一写入
也提供了flush方法可以手动控制将缓存的数据写入到目的地
```
type Writer struct {
	err error
	buf []byte
	n   int
	wr  io.Writer
}
需要注意的是err字段当write出错的时候，会将err字段赋值，writer将不能在进行写入操作
```
当写入的数据量大于buf的大小的时候会直接进行write的操作，从而跳过内部的缓存
提供了reset方法将Writer进行重置， reset可以使得buf进行复用，但是在调用reset的时候最好先调用flush将缓存的数据落地，因为reset只是简单的将数据丢弃而不会将数据落地

## Reader功能
通过使用bufio.Reader可以用io.Reader中更大批量的读取数据。这样可以有效的减少I/O操作(系统调用)
	read中的方法 三个相似的方法ReadSlice(), ReadString(),ReadBytes()
	`后面的两个方法都是在内部调用ReadSlice()实现的`
三者之前是有区别的，后两者在继续读的时候不会影响上一次读取的结果，因为在是buf中copy出来的，ReadSlice()是直接操作buf的，而不是copy一份出来的
	
	


## Scanner功能