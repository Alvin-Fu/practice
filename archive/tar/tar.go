package tar

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"practice/archive/zip"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// 测试打开tar包
func Unpack(name string) error {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("open file err: %v", err)
		return err
	}
	defer file.Close()
	read := tar.NewReader(file)
	for header, err := read.Next(); err != io.EOF; header, err = read.Next(){
		if err != nil {
			log.Fatalf("read next err: %v", err)
			return err
		}
		if zip.IsContinue(header.Name){
			continue
		}
		f, err := os.Create(header.Name)
		if err != nil {
			log.Fatalf("create file err: %v", err)
			return err
		}
		defer f.Close()
		_, err = io.Copy(f, read)
		if err != nil {
			log.Fatalf("io copy err: %v", err)
			return err
		}
	}


	return nil
}

func unpackDir(name string, fi os.FileInfo)error{
	if fi.IsDir(){
		unpackDir(fi.Name(), fi)
	}
	return nil
}
func unpackFile(){}

// 将一个文件打包
func Pack(name string, packName string) error {
	file, err := os.Create(packName)
	if err != nil {
		log.Fatalf("create file err: %v", err)
		return err
	}
	defer file.Close()
	write := tar.NewWriter(file)
	defer write.Close()
	fileInfo, err := os.Stat(name)
	if err != nil {
		log.Fatalf("get file info err: %v", err)
		return err
	}
	basePath := filepath.Dir(filepath.Clean(name))
	relativePath := filepath.Base(filepath.Clean(name))
	if fileInfo.IsDir(){
		return packDir(basePath, relativePath, write, fileInfo)
	} else {
		return packFile(basePath, relativePath, write, fileInfo)
	}
	return nil
}
func packDir(basePath string, relativePath string, tw *tar.Writer, fi os.FileInfo)error{
	if !isPathSeparator(basePath){
		basePath += string(os.PathSeparator)
	}
	if !isPathSeparator(relativePath){
		relativePath += string(os.PathSeparator)
	}
	path := basePath + relativePath
	fis, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalf("read dir err: %v", err)
		return err
	}
	for _, fi := range fis{
		if fi.IsDir(){
			packDir(basePath, relativePath + fi.Name(), tw, fi)
		} else {
			packFile(basePath, relativePath + fi.Name(), tw, fi)
		}
	}
	if len(relativePath) > 0{
		h, err :=tar.FileInfoHeader(fi, "")
		if err != nil {
			log.Fatalf("get header err: %v", err)
			return err
		}
		h.Name = relativePath
		if err := tw.WriteHeader(h); err != nil{
			log.Fatalf("write header err: %v", err)
			return err
		}
	}
	return nil
}

func packFile(basePath string, relativePath string, tw *tar.Writer, fi os.FileInfo)error{
	path := basePath + relativePath
	header, err := tar.FileInfoHeader(fi, "")
	if err != nil {
		log.Fatalf("get header err: %v", err)
		return err
	}
	header.Name = relativePath
	if err := tw.WriteHeader(header); err != nil{
		log.Fatalf("write header err: %v", err)
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file err: %v", err)
		return err
	}
	defer file.Close()
	if _, err := io.Copy(tw, file); err != nil {
		log.Fatalf("io copy err: %v", err)
		return err
	}
	return nil
}

func isPathSeparator(path string)bool{
	last := len(path) - 1
	if path[last] != os.PathSeparator{
		return false
	}
	return true
}