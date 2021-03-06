// 날짜 : 2021 03 06
// 이름 : 간단한 조건문 계산기
// 수행내역
// -bufio를 통한 시스템 입력을 받는 방법
// -go에서 시스템 입력값들을 활용하는 방법
// -bufio가 입력 받아지는 형태
// -간단한 switch 문에 대한 정보

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
	// 1) bufio의 New Reader를 통해 시스템에서 얻어진 값을 읽어오는 부분
	reader := bufio.NewReader(os.Stdin)
	// 2) bufio에서 우리는 숫자값으로 형변환을 시켜야되는데
	//    bufio에서는 바로 int형으로 가지고 오지 못하여 문자열에서
	//    제거 되지않은 \n 개행을 제거해준다
	line, _ := reader.ReadString('\n')
	// 3) trimspace를 통해 쓸데없은 앞뒤에 공백을 삭제
	line = strings.TrimSpace(line)
	// 4) 마지막으로 line에 담겨진 숫자(string)을 int로 변경해준다
	n1, _ := strconv.Atoi(line)

	fmt.Println("n2")
	reader = bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	n2, _ := strconv.Atoi(line)

	fmt.Println("어떠한결과를 원하십니까?")
	fmt.Println("1) 덧셈")
	fmt.Println("2) 뺼셈")
	fmt.Println("3) 곱셈")
	fmt.Println("4) 나눗셈")

	reader = bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)
	// bufio를 통해 입력을 하게 되면 시스템에서 아직 못읽는 값이 오기 떄문에 처리과정을 거쳐야
	// 되는 이유를 확인 할수가 있다.
	fmt.Println(reader)
	fmt.Println("os에서 입력되어진 형태 ")
	fmt.Println(line)
	fmt.Println("os에서 입력되어진 형태 ")
	var sel string = line
	// 기존과 다르게 break가 필요가 없다.
	// 만약에 스르륵 아래로 빠지게 하고 싶다면
	// fallthrough 을 붙여주면 된다 자세한 사항을 알고 싶으면
	// 해당 챕터 파일 test.go의 예제를 참조 하면된다
	switch {
	case sel == "1":
		fmt.Println(n1 + n2)
	case sel == "2":
		fmt.Println(n1 - n2)
	case sel == "3":
		fmt.Println(n1 * n2)
	case sel == "4":
		fmt.Println(n1 / n2)
	default:
		fmt.Println("잘못입력하셨습니다")
	}

}
