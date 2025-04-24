package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file error:", err)
		return []string{}
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file) // 类似于java中的迭代器，这里是扫描器
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}

func main() {

	list := readFile("output.txt")
	urlMap := make(map[string]int) //（url，contentLength）	go1.24这样声明 调试会直接panic，bug？
	countMap := make(map[int]int)  //（contentLength，count）

	// 遍历切片，放到urlMap中
	for _, v := range list {
		parts := strings.Split(v, " ")
		url := parts[0]
		// 去掉方括号并构建urlMap
		contentLength, _ := strconv.Atoi(strings.Trim(parts[1], "[]"))
		urlMap[url] = contentLength
	}

	//println(len(urlMap))

	// 对`content_length`进行统计，放到countMap中（content_length，count）
	for _, contentLength := range urlMap {
		//fmt.Println(k, v)
		countMap[contentLength] = 0 // 初始化countMap
	}

	// 遍历countMap， 对contentLength进行计数并放入countMap
	for contentLength, _ := range countMap {
		i := 0
		for _, contentLength_ := range urlMap {
			if contentLength_ == contentLength {
				i++
			}
		}
		countMap[contentLength] = i
	}

	// 寻找count频率较低的url，去除干扰
	for contentLength, count := range countMap {
		//fmt.Println(contentLength, count)
		if count <= 3 && contentLength >= 1048576 { // 只要大于1MB的
			for url, contentLength_ := range urlMap {
				if contentLength_ == contentLength {
					fmt.Printf("%v %.1f MB lenth频率 %v\n", url, float64(contentLength)/1024/1024, count)
				}
			}
		}
	}
}
