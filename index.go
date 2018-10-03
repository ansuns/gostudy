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
}
