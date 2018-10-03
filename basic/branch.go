package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename  = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(0),
		grade(30),
		grade(60),
		grade(77),
		grade(90),
		grade(110),
		)
}

func grade(score int)  string{
	str := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		str = "F"
	case score < 80:
		str = "C"
	case score < 90:
		str = "B"
	case score	<= 100:
		str = "A"

	}

	return str
}

func fors()  {

}
