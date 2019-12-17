package main

import (
	"archive/tar"
	"fmt"
	"log"
	"os"
)

func main() {
	TestFileInfoHeader()
}

func TestFileInfoHeader() error {
	fi, err := os.Stat("readme.md")
	if err != nil {
		log.Fatalf("file stat err: %v", err)
		return err
	}
	header, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		log.Fatalf("file info header err: %v", err)
	}
	header.Linkname = ""
	fmt.Println(header.Linkname)
	return nil
}
