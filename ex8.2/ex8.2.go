package main

import "fmt"

func main() {
	const pi1 float64 = 3.141592 // 상수
	var pi2 float64 = 3.141592   // 변수

	// pi1 = 3
	pi2 = 4

	// fmt.Printf("원주율 : %f\n", pi1) //cannot asign to pi1 (declared const)
	fmt.Printf("원주율 : %f\n", pi2)
}

// 상수는 변하지 않는 값으로써 const로 정의한 값은 의미를 부여받는다.
