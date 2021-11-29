package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	// 0.100000 + 0.200000 == 0.300000 : false
	// 왜 그럴까? 아래 a+b를 출력해보자
	fmt.Println(a + b)
	// a + b는 0.30000000000000004 이 나온다.
	// 실수부를 표현할 때 부호, 지수부, 소수부로 나뉘는데 0.3을 정확히 나타내지 못하고 근사치로 갖고 오기 때문에 오차가 발생한다.
	// 근사치로밖에 못 갖고 오는 이유는 0.3을 2진수로 딱 정확히 나타낼 수가 없기 때문이다.

	fmt.Println(c) // 0.3
}
