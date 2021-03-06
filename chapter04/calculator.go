// 날짜 : 2021 03 06
// 이름 : 계산기 반복문을 통한 무한 입력처리
// 수행내역
// - for문 수행 내용

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("숫자를 입력하세요! ")
	fmt.Println("n1")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n1, _ := strconv.Atoi(line)

	fmt.Println("n2")
	reader = bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	i := "0"

	// 기본적인 for문 문자열 비교도 가능하다
	for i != "5" {
		fmt.Println("어떠한결과를 원하십니까?")
		fmt.Println("1) 덧셈")
		fmt.Println("2) 뺼셈")
		fmt.Println("3) 곱셈")
		fmt.Println("4) 나눗셈")
		fmt.Println("5) 나가기")

		reader = bufio.NewReader(os.Stdin)
		line, _ = reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(reader)
		// fmt.Println("os에서 입력되어진 형태 ")
		// fmt.Println(line)
		// fmt.Println("os에서 입력되어진 형태 ")
		var sel string = line
		i = line
		// 기존과 다르게 아래와 같이 스위치문 구성도 가능

		switch sel {
		case "1":
			fmt.Println(n1 + n2)
		case "2":
			fmt.Println(n1 - n2)
		case "3":
			fmt.Println(n1 * n2)
		case "4":
			fmt.Println(n1 / n2)
		case "5":
			fmt.Println("수고하셨습니다!")
		default:
			fmt.Println("잘못입력하셨습니다")
		}
	}

}

//위와같은 형태 이외에도 일반적인 형기으로 진행하는것도 가능하다

/*
1) 선언식
for i := 0; i < 10; i++{ // 선언과 동시에 for 문 사용하기 가능 <- 여기서 선언된 i는 외부에서 사용하는것은 불가능하다.
	                     // ++ 연산을 통해 처리도 가능
	if i == 0{
		fmt.Printf("시작!")
	}
	fmt.Printf("%v", i)
}
2) 기본상태
 for { <- default로 항상 true의 형태를 가지고 있다

 }
 3) 무한루프 제어
  break 및 continue 사용이 일반적이랑 같으니 애용하여 사용하자

 테스트2

*/
