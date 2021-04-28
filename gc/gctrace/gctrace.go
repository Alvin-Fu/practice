package main

func main() {
	for n := 1; n < 100000; n++ {
		allocate()
	}
}

func allocate() {
	_ = make([]byte, 1<<20)
}
