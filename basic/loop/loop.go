package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	//对2取模，取出来就是最低位
	//for 这里没有起始条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

/**
读取文件
 */
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//13
	/**
	最低位一定是1
	13%2
	 */
	fmt.Println(
		converToBin(13),
		converToBin(19910805),
	)
	readFile("abc.txt")
}
