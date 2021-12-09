package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

//													함수타입으로, int 두개를 입력받고 int 하나를 출력하는 함수
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	} else {
		return nil
	}
}

func main() {
	var operator func(int, int) int
	operator = getOperator("+")

	var result = operator(3, 4)
	fmt.Println(result)
}

// 결과값 : 7
