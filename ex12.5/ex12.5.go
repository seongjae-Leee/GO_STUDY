package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	b := [5]int{1000, 2000, 3000, 4000, 5000}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	fmt.Println()
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

// 배열의 크기가 같아야 복사가 가능하다.
/*
결과값
a[0] = 10
a[1] = 20
a[2] = 30
a[3] = 40
a[4] = 50

b[0] = 1000
b[1] = 2000
b[2] = 3000
b[3] = 4000
b[4] = 5000

b[0] = 10
b[1] = 20
b[2] = 30
b[3] = 40
b[4] = 50

*/
