package main

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//UncompressFile("readme.md", "readme.zip")
	UncompressDir("../actomic/", "actomic.zip")
}
func UncompressFile(name string, zipName string) error {
	zipFile, err := os.Create(zipName)
	if err != nil {
		log.Fatalf("create file err: %v", err)
		return err
	}
	defer zipFile.Close()
	write := zip.NewWriter(zipFile)
	defer write.Close()

	fileInfo, err := os.Stat(name)
	if err != nil {
		log.Fatalf("file get info err: %v", err)
		return err
	}
	// 生成目标文件的header即描述
	fHeader, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		log.Fatalf("get header err: %v", err)
		return err
	}
	w, err := write.CreateHeader(fHeader)
	if err != nil {
		log.Fatalf("create header err: %v", err)
		return err
	}
	of, err := os.Open(name)
	io.Copy(w, of)
	return nil
}

func UncompressDir(dir string, zipName string) error {
	if dir == "" {
		return errors.Errorf("dir err")
	}
	if zipName == "" {
		return errors.Errorf("zip file name err")
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalf("read the dir err: %v", err)
		return err
	}
	zFile, err := os.Create(zipName)
	if err != nil {
		log.Fatalf("create zip file err: %v", err)
		return err
	}
	defer zFile.Close()
	writer := zip.NewWriter(zFile)
	defer writer.Close()
	return WriteData(writer, files, dir)
}

func WriteData(writer *zip.Writer, files []os.FileInfo, dir string) error {
	for _, f := range files {
		w, err := writer.Create(f.Name())
		if err != nil {
			log.Fatalf("cteate err: %v", err)
			return err
		}
		data, err := ioutil.ReadFile(dir + f.Name())
		if err != nil {
			log.Fatalf("read file err: %v", err)
			return err
		}
		_, err = w.Write(data)
		if err != nil {
			log.Fatalf("write data err: %v", err)
			return err
		}
	}
	return nil
}

func Compress() {}
