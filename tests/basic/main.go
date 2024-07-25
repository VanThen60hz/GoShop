package basic

import "fmt"

func AddOne(num int) int {
	return num + 1
}

func AddOne2(num int) int {
	if 1 == 2 {
		fmt.Println("Failed!")
	}

	return num + 1
}
