package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	// var c int = b
	// d = a * b
	// var e int64 = 7
	// // 위에는 타입이 달라서 그렇다 쳐도 a의 int는 int64와 같은데 그럼에도 연산이 불가능하다.
	// // 즉 int와 int64는 사이즈가 같으나 다른 타입이다.
	// f := a * e

	// 그래서 타입을 바꿔주게 되면 연산 가능
	// 소수를 정수값으로 바꾸면 소수점 뒤가 날아감
	var c int = int(b)
	d := float64(a) * b
	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
