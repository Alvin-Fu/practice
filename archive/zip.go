package main

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//UncompressFile("readme.md", "readme.zip")
	//UncompressDir("../../rpcx/", "rpc.zip")
	Compress("rpc.zip", "rpcx/")
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

func UncompressDir(path string, zipName string) error {
	if path == "" {
		return errors.Errorf("path err")
	}
	if zipName == "" {
		return errors.Errorf("zip file name err")
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("read the path err: %v", err)
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
	return WriteData(writer, files, path, "rpcx/")
}

func WriteData(writer *zip.Writer, files []os.FileInfo, path string, dirName string) error {
	for _, f := range files {
		if f.IsDir() {
			dirName = currentPath(dirName, f.Name()+"/")
			fis, err := ioutil.ReadDir(path + f.Name())
			if err != nil {
				log.Fatalf("read path err: %v", err)
				return err
			}
			WriteData(writer, fis, path+f.Name()+"/", dirName)
			dirName = upperPath(dirName, f.Name()+"/")
		} else {
			w, err := writer.Create(dirName + f.Name())
			if err != nil {
				log.Fatalf("cteate err: %v", err)
				return err
			}
			data, err := ioutil.ReadFile(path + f.Name())
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
	}
	return nil
}

func Compress(zipName string, dir string) error {
	if dir == "" {
		return errors.Errorf("dir err")
	}
	if zipName == "" {
		return errors.Errorf("zip file name err")
	}
	err := os.MkdirAll(dir, 754)
	if err != nil {
		log.Fatalf("make dir err: %v", err)
		return err
	}
	readCloser, err := zip.OpenReader(zipName)
	if err != nil {
		log.Fatalf("open reader err: %v", err)
		return err
	}
	defer readCloser.Close()
	for _, file := range readCloser.File {
		if isContinue(dir + file.Name) {
			continue
		}
		r, err := file.Open()
		if err != nil {
			log.Fatalf("file open err: %v", err)
			return err
		}
		defer r.Close()
		f, err := os.Create(dir + file.Name)
		if err != nil {
			log.Fatalf("create file err: %v, name: %s", err, file.Name)
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, r)
		if err != nil {
			log.Fatalf("io copy err: %v", err)
			return err
		}
	}
	return nil
}

func isContinue(str string) bool {
	bytes := []byte(str)
	if bytes[len(bytes)-1] == '/' {
		err := os.MkdirAll(str, 754)
		if err != nil {
			log.Println(err)
		}
		return true
	}
	n := strings.LastIndex(str, "/")
	str = string(bytes[:n]) + "/"
	err := os.MkdirAll(str, 754)
	if err != nil {
		log.Println(err)
	}
	return false
}

func currentPath(prentDir string, dir string) string {
	return prentDir + dir
}

func upperPath(currentDir string, dir string) string {
	index := strings.LastIndex(currentDir, dir)
	return string([]byte(currentDir)[:index])
}
