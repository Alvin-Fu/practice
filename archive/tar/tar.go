package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// 测试打开tar包
func Unpack(name string) error {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("tar open fail, err: %v", err)
		return err
	}
	defer file.Close()
	// 将打开的tar文件，填入reader结构体中
	reader := tar.NewReader(file)
	for h, err := reader.Next(); err != io.EOF; h, err = reader.Next() {
		if err != nil {
			log.Fatalf("reader next err: %v", err)
			return err
		}
		fmt.Printf("heard: %+v\n", h)
		// 打开tar文件后将内容写文件中
		f, err := os.Create(h.FileInfo().Name())
		if err != nil {
			log.Fatalf("file create fial, err: %v", err)
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, reader)
		if err != nil {
			log.Fatalf("io copy err: %v", err)
			return err
		}
	}

	return nil
}

// 将一个文件打包
func Pack(name string, packName string) error {
	f, err := os.Create(packName)
	if err != nil {
		log.Fatalf("file create fail, err: %v", err)
		return err
	}
	defer f.Close()
	write := tar.NewWriter(f)
	defer write.Close()
	fileInfo, err := os.Stat(name)
	if err != nil {
		log.Fatalf("get file info err: %v", err)
		return err
	}
	header, err := tar.FileInfoHeader(fileInfo, "hello")
	if err != nil {
		log.Fatalf("get header err: %v", err)
		return err
	}
	err = write.WriteHeader(header)
	if err != nil {
		log.Fatalf("writer header err: %v", err)
		return err
	}
	opFile, err := os.Open(name)
	if err != nil {
		log.Fatalf("file open err: %v", err)
		return err
	}
	num, err := io.Copy(write, opFile)
	if err != nil {
		log.Fatalf("io copy err: %v", err)
		return err
	}
	fmt.Println(num)
	return nil
}
