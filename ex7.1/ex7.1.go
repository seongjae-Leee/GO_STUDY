package main

import "fmt"

// func 함수이름 (매개변수) 반환타입{함수코드}
func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

// 즉 반복을 줄이기 위해 함수(코드블럭)를 만들어놓고 호출해서 쓴다.
