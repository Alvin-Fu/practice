package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//TestFileInfoHeader()
	TestReaderTar()
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
	fmt.Printf("header: %+v", header)
	return nil
}

func TestReaderTar()error{
	file, err := os.Open("tls.tar")
	if err != nil {
		log.Fatalf("tar open fail, err: %v", err)
		return err
	}
	defer file.Close()
	reader := tar.NewReader(file)
	for h, err := reader.Next(); err != io.EOF; h, err = reader.Next(){
		if err != nil {
			log.Fatalf("reader next err: %v", err)
			return err
		}
		fmt.Printf("heard: %+v",h)
	}
	return nil
}
