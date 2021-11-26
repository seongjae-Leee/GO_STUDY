package main

import "fmt"

func main() {
	// float32는 7자리까지, float64는 15자리까지
	var a float32 = 1234.567
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) // 4.2668155e+06 가 뜨면서 소수점 여섯자리가 다 짤림
	fmt.Println(d) // 1.2800446e+07 으로 오차가 더 커져서 소수 부분도 짤리고 추가적으로 정수 부분도 다르게 됨.
}
