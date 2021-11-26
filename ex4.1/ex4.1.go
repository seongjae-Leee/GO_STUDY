package main

import "fmt"

func main() {
	// 변수 이름과 타입 지정
	var a int = 10
	// a := 10   <- 이렇게 적어도 초기값만 할당해주면 var와 int를 생략할 수 있다.
	var msg string = "Hello Viable"

	a = 20
	msg = "good morning"
	fmt.Println(msg, a)
}
