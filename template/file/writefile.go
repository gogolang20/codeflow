package main

import (
	"bufio"
	"log"
	"os"
)

// 快写文件
func writeAll(filename string) error {
	err := os.WriteFile("mike.txt", []byte("Hi mike\n"), 0666)
	if err != nil {
		return err
	}
	return nil
}

/*
	按行写文件
*/

// 直接操作IO
func writeLine(filename string) error {
	data := []string{
		"mike",
		"test",
		"123",
	}
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// 使用缓存区写入
func writeLine2(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	for i := 0; i < 2; i++ {
		// 写字符串到buffer
		bytesWritten, err := bufferedWriter.WriteString(
			"mike 真帅\n",
		)
		if err != nil {
			return err
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}

	// 写内存buffer到硬盘
	return bufferedWriter.Flush()
}

/*
	偏移量写入
*/

func writeAt(filename string) error {
	data := []byte{
		0x41, // A
		0x73, // s
		0x20, // space
		0x20, // space
		0x67, // g
	}
	f, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	replaceSplace := []byte{
		0x6F, // o
		0x6E, // n
	}
	_, err = f.WriteAt(replaceSplace, 2)
	if err != nil {
		return err
	}

	return nil
}

/*
	缓存区写入
*/

func writeBuffer(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字符串到buffer
	bytesWritten, err := bufferedWriter.WriteString(
		"mike 真帅\n",
	)
	if err != nil {
		return err
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	if err != nil {
		return err
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil {
		return err
	}

	return nil
}
