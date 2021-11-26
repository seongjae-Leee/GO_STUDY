package main

import "fmt"

func main() {
	// 기본 선언
	var a int = 3
	// 초기값을 뺀 선언(int의 default 값인 0으로 들어감)
	var b int
	// int라고 타입을 주지는 않았지만 초기값을 주어서 int로 인식
	var c = 4
	// 기본선언에서 var과 타입을 뺀 것
	d := 5
	// string이라고 타입을 주지는 않았지만 초기값을 주어서 string으로 인식
	var e = "hi hi"
	// float라는 타입을 주진 않았지만 소수값이므로 float 타입으로 인식
	f := 3.14

	// ★★★코드가 바뀌고 다시 실행할 때는 go build를 꼭 다시 입력해줘야 함★★★
	fmt.Println(a, b, c, d, e, f)
}

// 타입이 다른 float과 int 끼리는 연산이 안 된다.
