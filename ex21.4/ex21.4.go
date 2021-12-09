package main

import "fmt"

// 아래 두개처럼 함수를 따로 정의하지 말고 getOperator에서 바로 정의해서 사용해보자.
// func add(a, b int) int {
// 	return a + b
// }

// func mul(a, b int) int {
// 	return a * b
// }

//													함수타입으로, int 두개를 입력받고 int 하나를 출력하는 함수
func getOperator(op string) func(int, int) int {
	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
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
