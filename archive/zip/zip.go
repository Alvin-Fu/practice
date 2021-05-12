package zip

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"practice/archive/util"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func UnCompress(name string, zipName string, dirName string) error {
	if name == "" {
		return errors.Errorf("name err")
	}
	if zipName == "" {
		return errors.Errorf("zip file name err")
	}
	files, err := FileInfo(name)
	if err != nil {
		log.Fatalf("get file info err: %v", err)
		return err
	}
	path := util.ParentDir(name)
	zFile, err := os.Create(zipName)
	if err != nil {
		log.Fatalf("create zip file err: %v", err)
		return err
	}
	defer zFile.Close()
	writer := zip.NewWriter(zFile)
	defer writer.Close()
	return WriteData(writer, files, path, dirName)
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
		// 对目录和文件进行区分
		if util.IsContinue(dir + file.Name) {
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

func FileInfo(name string) ([]os.FileInfo, error) {
	if len(name) <= 0 {
		log.Fatalf("name err")
		return nil, errors.Errorf("name err")
	}
	fi, err := os.Stat(name)
	if err != nil {
		log.Fatalf("get fileInfo err: %v", err)
		return nil, err
	}
	files := make([]os.FileInfo, 0)
	if fi.IsDir() {
		f, err := ioutil.ReadDir(name)
		if err != nil {
			log.Fatalf("read dir err: %v", err)
			return nil, err
		}
		files = append(files, f...)
	} else {
		files = append(files, fi)
	}
	return files, nil
}

func currentPath(prentDir string, dir string) string {
	return prentDir + dir
}

func upperPath(currentDir string, dir string) string {
	index := strings.LastIndex(currentDir, dir)
	return string([]byte(currentDir)[:index])
}
