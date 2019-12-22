golang中的archive/zip包提供zip归档文件
在使用的过程中主要是对目录进行zip格式的打包和解包，相对来说比较麻烦，特别是要保持原来的目录结构
## 使用遇到的问题
对单纯的文件进行打包和解包的时候比较简单，不需要做过多的处理
主要有目录层次的在打包的过程中需要保持原有的目录结构
### 打包
在打包的过程中需要去判断当前的file是否是目录，如果是目录需要进行递归，一层一层的打开，进行打包
```
for _, f := range files {
		if f.IsDir() {
			//下面的代码就是对目录进行处理
			dirName = currentPath(dirName, f.Name()+"/")
			fis, err := ioutil.ReadDir(path + f.Name())
			if err != nil {
				log.Fatalf("read path err: %v", err)
				return err
			}
			WriteData(writer, fis, path+f.Name()+"/", dirName)
			// 回退到上一级目录
			dirName = upperPath(dirName, f.Name()+"/")
		} else {
			...		
		}
}
```

### 解包
在解包的时候也是一样的，需要判断当前是否是目录，如果是目录就需要先创建目录，然后再将创建文件，以恢复目录中的文件
如果没有创建目录再使用的时候可能会出错
```
for _, file := range readCloser.File {
	if isContinue(dir + file.Name) {
		continue
	}
	...
}
 isContinue这段代码就是为了检查是否是目录,如果是目录就创建目录并继续
```
