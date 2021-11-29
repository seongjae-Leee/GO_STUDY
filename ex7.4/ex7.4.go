package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} else {
		return a / b, true
	}
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	// success를 또 선언할 수 있는 이유는 d가 선언이 안되어있기 때문에 묶여서 선언할 수 있는 것.
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}

// 3 true
// 0 false

/////////////////////////////////////////////////////////////////////////////////////////
//// 이렇게도 쓸 수 있음
// func Divide(a, b int) (result int, success bool) {
// 	if b == 0 {
// 		result = 0
// 		success = false
// 		return
// 	} else {
// 		result = a / b
// 		success = true
// 		return
// 	}
// }

// func main() {
// 	c, success := Divide(9, 3)
// 	fmt.Println(c, success)
// 	d, success := Divide(9, 0)
// 	fmt.Println(d, success)
// }
