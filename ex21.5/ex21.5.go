package main

import "fmt"

func main() {
	i := 0

	f := func() {
		i += 10 // 함수 literal 에서는 외부의 변수를 캡쳐해서 내부 상태로 가져올 수 있다.
		// 근데 중요한게 캡쳐는 값복사가 아니라 레퍼런스 복사이다. 즉, 포인터가 복사되는 것이다.
	}
	i++
	f()
	fmt.Println(i)
}

// 결과값 : 11
