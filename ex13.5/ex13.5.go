package main

import (
	"fmt"
	"unsafe" // 안전하지 않은 함수 제공.. 메모리 공간 크기 제공 등
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {
	var a int = 8
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user), unsafe.Sizeof(a))
}

// 결과값(바이트) : 16 8
// 12가 나와야 될 것 같은데 실제로는 age와 score 사이에 4바이트 만큼의 memory padding이 있다.
