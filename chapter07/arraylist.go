/*
	일자 : 21.03.07
	배열
	- 자바와 유사하게 쓰인다는 것을 암
	- 문자열 len을 통한 크기값 알아내기
	- 문자 기본 타입이 ascii가 아니라 utf8로 이루어짐
	- 한글 코드 출력 3바이트 가능
	- rune를 통한 3바이트 하나의 코드로 표현 및 출력
*/
package main

import "fmt"

func main() {
	// 선언식
	var a [10]int
	// len을 통해 Length를 보낼수가 있다.
	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}
	fmt.Println(a)
	// 출력 형태
	//[0 1 4 9 16 25 36 49 64 81]

	/*
		 문자 관련
		- 문자는 모두 배열로 이루어져 있다
		- 기본적으로 코딩언어들이 ascii 2 방식을 취하고 있지만 Go와 같은 경우에는 utf-8
		로 이루어져 있기떄문에 한가지 특징을 볼수가 있는데
	*/
	var test string = "Hello world"
	// 코드값으로 호출
	for i := 0; i < len(test); i++ {
		fmt.Print("[", test[i], "],")
	}
	fmt.Println("")
	//string을 통해 숫자 코드값을 언어로 변환가능
	for i := 0; i < len(test); i++ {
		fmt.Print("[", string(test[i]), "],")
	}
	fmt.Println("")
	// 유니코드 이기때문에 한글또한 호출이 가능한데
	// 한글과 같은경우 3바이트로 이루어져 있기떄문에
	// 3바이트로 분류가 되서 코드로 호출하게 되면 깨짐 현상을 볼수가 있다.
	test = "Hello 월드"
	// 코드값으로 호출
	for i := 0; i < len(test); i++ {
		fmt.Print("[", test[i], "],")
	}
	fmt.Println("")
	//string을 통해 숫자 코드값을 언어로 변환가능
	for i := 0; i < len(test); i++ {
		fmt.Print("[", string(test[i]), "],")
	}
	fmt.Println("")

	// 이제 이러한 문자를 Go 랭기지에서는 꺠지지 않게 변형해주는 새로운 타입이 존재합니다.
	// rune <- 변수타입으로 배열을 선언해주고 받아주면 이것이 가능합니다
	var test2 []rune = []rune(test)
	// test := []rune(test) <- 간략식으로 표현도 가능
	for i := 0; i < len(test2); i++ {
		fmt.Print("[", test2[i], "],")
	}
	fmt.Println("")
	//string을 통해 숫자 코드값을 언어로 변환가능
	for i := 0; i < len(test2); i++ {
		fmt.Print("[", string(test2[i]), "],")
	}
	fmt.Println("")
	/*
		이렇게 3바이트를 하나의 코드로 표현해주고
		[72],[101],[108],[108],[111],[32],[50900],[46300],
		정상적인 출력을 가능하게 해준다
		[H],[e],[l],[l],[o],[ ],[월],[드],
	*/

}
