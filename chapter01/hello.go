package main

import (
	"fmt"
)

func main() {
	var num1 int = 4
	var num2 int = 5
	var data string = "이것이 고 언어인가?"
	fmt.Println(num1 + num2)
	fmt.Println(data)
	fmt.Println(addnum(5, 4))
}

func addnum(num1 int16, num2 int16) (str string, data int16) {
	data = num1 + num2
	str = "결과값입니다."
	return str, data
}
