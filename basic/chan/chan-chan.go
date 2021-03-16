package main

func main() {
	var x chan<- chan error   // nil
	var y chan (<-chan error) // nil
	println(x == y)
}
