package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

/*
	读取全文件
*/

func readAll(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	log.Printf("read %s content is %s", filename, data)
	return nil
}

func ReadAll2(filename string) error {
	file, err := os.Open("asong.txt")
	if err != nil {
		return err
	}

	content, err := io.ReadAll(file)
	log.Printf("read %s content is %s\n", filename, content)

	file.Close()
	return nil
}

/*
	逐行读取
*/

func readLine(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	bufferedReader := bufio.NewReader(file)
	for {
		// ReadLine is a low-level line-reading primitive. Most callers should use
		// ReadBytes('\n') or ReadString('\n') instead or use a Scanner.
		lineBytes, err := bufferedReader.ReadBytes('\n')
		bufferedReader.ReadLine()
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		log.Printf("readline %s every line data is %s\n", filename, line)
	}
	file.Close()
	return nil
}

/*
	按块读取文件
*/
// use bufio.NewReader
func readByte(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	// 创建 Reader
	r := bufio.NewReader(file)

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// use os
func readByte2(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// use os and io.ReadAtLeast
func readByte3(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := io.ReadAtLeast(file, buf, 0)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

/*
	分隔符读取
*/
func readScanner(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	// 可以定制Split函数做分隔函数
	// ScanWords 是scanner自带的分隔函数用来找空格分隔的文本字
	scanner.Split(bufio.ScanWords)
	for {
		success := scanner.Scan()
		if success == false {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
				break
			} else {
				return err
			}
		}
		// 得到数据，Bytes() 或者 Text()
		log.Printf("readScanner get data is %s", scanner.Text())
	}
	file.Close()
	return nil
}
