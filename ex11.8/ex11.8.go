package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor: // Label (플래그 변수처럼 이중 for문에서 쓸 수 있는..)
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // 범위 지정 레이블
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

// 레이블은 지원하지 않는 언어가 많기 때문에, 그리고 플래그를 써도 코드가 그렇게 안 지저분하기 때문에
// 웬만하면 안 쓰는게 낫긴 하다.

/*
결과값 :
5 * 9 = 45
*/
