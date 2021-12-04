// 문자열은 일부만 수정은 불가능하고 전체를 새로 써서 수정해야 한다.

package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a'

	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}

/*
결과값 :
Hello World
Healo World
*/
// slice를 통해서 바이트 단위 부분 수정이 가능함을 알 수 있다.
