package main

import (
	"fmt"
	"io/ioutil"
)

/**
接受score，返回string
 */
func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	case score > 100 || score < 0:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	/*	contents, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%s\n", contents)
		}*/
	//go语言的if可以像for一样写(if后跟一个分号)
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//出了IF就不能访问 contenst, err

	fmt.Println(
		grade(0),
		grade(20),
		grade(71),
		grade(82),
		grade(90),
	)

}
