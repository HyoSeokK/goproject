package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 5
	if a > b {
		fmt.Println("이것은 참이다")
	} else if b > a {
		fmt.Println(" 거짓이다")
	} else {
		fmt.Println("이것은 같다")
	}

	if a > b || a == b {
		fmt.Println("a는 b와 같거나 크다")
	}
}
