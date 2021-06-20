package controller

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"process_data/cache"
)

func SaveSingleData(filePath string, bufSize int) (err error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("file size=", size, "B")

	buf := make([]byte, bufSize)
	_, err = file.Read(buf)
	if err != nil {
		fmt.Println("SaveMockData:ReadBuf:Err:", err)
		return
	}
	err = cache.Cache().SET(fmt.Sprintf("key_%db", bufSize), string(buf))
	return err
}

func SaveMultiData(filePath string) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	row := 1
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
		for i := 1; i <= 50; i++ {
			key := fmt.Sprintf("%d_%d", row, i)
			err = cache.Cache().SET(key, line)
			if err != nil {
				fmt.Println("SaveMultiData:Set:", err)
			}
		}
		row++
	}

	return
}
