/*
 일자 : 21.03.07

 함수란?
	기본적인 형식은 같지만 형성하는 형태에 다름이 존재한다.

*/

package main

import "fmt"

// 함수 선언식 기본
// func 함수명 ( 입력 변수 ) 출력 변수 : 리턴값 { }
func add(x int, y int) int {
	return x + y
}

// 출력값이 2개인 경우
// func 함수명 (입력값) (출력값){   }
func switnum(x int, y int) (int, int) {
	return y, x
}

// 재귀함수
// 팩토리얼 예제
func factoryall(x int) int {
	if x == 1 {
		return 1
	}
	return x * factoryall(x-1)
}

/*
	강의자 왈
	재귀함수와 같은경우 수학적인 식들을 프로그램에서 처리 하는것에 유리하다.
*/
func main() {
	var x int = 5
	var y int = 6
	fmt.Println(add(x, y)) // 함수 호출은 기존 언어들과 동일

	fmt.Println(switnum(x, y))
	fmt.Println(factoryall(3))
}
