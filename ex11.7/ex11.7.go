package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true // found(★플래그 변수)는 a가 걸려있는 for문까지 빠져나오게 하는 수단
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/*
결과값 :
5 * 9 = 45
*/
