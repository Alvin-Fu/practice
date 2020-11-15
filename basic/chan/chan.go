package main

func main() {
	done := make(chan string)
	data := make([]byte, 10)
	go func() {
		done <- string(data)
	}()
	data[0] = 1
	<-done
}
