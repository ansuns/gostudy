package main

import (
	"fmt"
	"regexp"
)

const text = `my email is abc@qq.com
emainl2 nila@163.com
email3 22@gmail.com
`

func main() {
	//re, err := regexp.Compile(text)
	prex := `[a-zA-Z0-9.]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+` //tab键上面这个
	prex2 := `([a-zA-Z0-9.]+)@([a-zA-Z0-9]+)\.([a-zA-Z0-9]+)` //tab键上面这个 加上括号可以提取
	re := regexp.MustCompile(prex)          //查找abc
	match := re.FindAllString(text, -1)     //在字符串中查找:-1所有

	re2 := regexp.MustCompile(prex2)          //查找abc
	match2:=re2.FindAllStringSubmatch(text,-1)
	fmt.Println(match)
	fmt.Println(match2)
}
