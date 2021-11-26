package main

import "fmt"

// 패키지 전역 변수
var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}
	// 변수의 범위로 인해 아래 식에서 s는 정의되지 않는다.
	// m = s + 20
}
