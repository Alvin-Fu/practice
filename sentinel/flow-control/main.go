package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fs, _ := ioutil.ReadDir(dir)
	for _, f := range fs {
		fmt.Println(f.ModTime().String())
	}
	err := t()
	fmt.Printf("%#v", err)
}

func t() error {
	return nil
}
