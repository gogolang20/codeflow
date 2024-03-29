package main

import (
	"log"
	"os"
)

/*
创建文件
打开文件
读取文件
写入文件
关闭文件
打包/解包
压缩/解压缩
改变文件权限
删除文件
移动文件
重命名文件
清空文件
*/

var fileName = "mac.txt"

func main() {
	// 创建文件
	f, err := os.Create("mike.txt")
	if err != nil {
		log.Fatalf("create file failed err=%s\n", err)
	}
	// 获取文件信息
	fileInfo, err := f.Stat()
	if err != nil {
		log.Fatalf("get file info failed err=%s\n", err)
	}

	log.Printf("File Name is %s\n", fileInfo.Name())
	log.Printf("File Permissions is %s\n", fileInfo.Mode())
	log.Printf("File ModTime is %s\n", fileInfo.ModTime())

	// 改变文件权限
	err = f.Chmod(0777)
	if err != nil {
		log.Fatalf("chmod file failed err=%s\n", err)
	}

	// 改变拥有者
	err = f.Chown(os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatalf("chown file failed err=%s\n", err)
	}

	// 再次获取文件信息 验证改变是否正确
	fileInfo, err = f.Stat()
	if err != nil {
		log.Fatalf("get file info second failed err=%s\n", err)
	}
	log.Printf("File change Permissions is %s\n", fileInfo.Mode())

	// 关闭文件
	err = f.Close()
	if err != nil {
		log.Fatalf("close file failed err=%s\n", err)
	}

	// 删除文件
	err = os.Remove("mike.txt")
	if err != nil {
		log.Fatalf("remove file failed err=%s\n", err)
	}
}
