// 재귀호출(Recursive Call : 자기 자신을 다시 부르는 것)
package main

import "fmt"

func printNo(n int) {
	if n == 0 { // <- 재귀호출 탈출 조건(필수★)
		return
	}
	fmt.Println(n)          // 3, 2, 1
	printNo(n - 1)          // <- 재귀 호출!!!
	fmt.Println("After", n) // <- 재귀 호출 이후 출력 (After 1, After 2, After 3)
}

func main() {
	printNo(3) // <- 함수 호출
}

/*
3
2
1
After 1
After 2
After 3
*/
