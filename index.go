package main

import (
	"fmt"
	"runtime"
)

func main() {

	for skip := 0; ; skip++{
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}

	firstName, secodName, nickName, realName := getName(3)
	fmt.Print(firstName, secodName, nickName, realName)
	fmt.Println()
	f := func(x, y int) (sum int) {
		return x + y
	}

	go fmt.Print(f(7, 8))
}

func getName(id int) (firstName, secodName, nickName, realName string)  {

	return "Zhang", "san", "Xiaosan", "zhang san"
}


