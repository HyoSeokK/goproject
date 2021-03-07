/*
*******
 *****
  ***
   *
   피라미드 만들어보기
*/

package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		for j := 0; j < 8; j++ {
			if j >= 0+i && j < 8-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
